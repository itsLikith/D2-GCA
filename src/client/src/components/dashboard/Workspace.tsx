'use client';

import React, { useState, useEffect, useMemo } from 'react';
import { Navigation, Grid3X3, Layers } from 'lucide-react';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { api, StationDTO, Measurement, Analyze3DResponse } from '@/lib/api';

import Logger from './Logger';
import ScenarioParameters from './ScenarioParameters';
import GroundBeacons from './GroundBeacons';
import RealTimeSolver from './RealTimeSolver';
import GridCoverage from './GridCoverage';
import ElevationSweeps from './ElevationSweeps';
export default function Workspace() {
  //--------------------------------------------------
  // State variables
  //--------------------------------------------------
  const [stations, setStations] = useState<StationDTO[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  // Selected station IDs for WLS analysis
  const [selectedStationIds, setSelectedStationIds] = useState<string[]>(['NGP', 'GDA']);

  // Aircraft state
  const [acX, setAcX] = useState<number>(10);
  const [acY, setAcY] = useState<number>(20);
  const [acAlt, setAcAlt] = useState<number>(30000); // feet
  const [altMode, setAltMode] = useState<string>('RVSM');

  // Solved WLS solution state
  const [wlsSolution, setWlsSolution] = useState<Analyze3DResponse | null>(null);
  const [wlsError, setWlsError] = useState<string | null>(null);

  // Logging state
  const [logs, setLogs] = useState<string[]>([]);

  useEffect(() => {
    let active = true;
    api
      .getStations()
      .then((data) => {
        if (active) {
          setStations(data);
          setLoading(false);
          setLogs((prev) => [
            ...prev,
            `[SYSTEM] Loaded ${data.length} stations from the database.`,
          ]);
        }
      })
      .catch((err) => {
        if (active) {
          console.error('Failed to load stations', err);
          // Fallback to static cached file in case backend is loading/offline
          import('@/lib/indiaStations').then(({ INDIA_STATIONS }) => {
            setStations(INDIA_STATIONS);
            setLoading(false);
            setLogs((prev) => [
              ...prev,
              `[WARNING] Failed to load database stations: ${err.message}. Loaded cached stations dataset instead.`,
            ]);
          });
        }
      });
    return () => {
      active = false;
    };
  }, []);

  //--------------------------------------------------
  // Math & API calls for WLS Analysis
  //--------------------------------------------------
  const activeStations = useMemo(() => {
    return stations.filter((s) => selectedStationIds.includes(s.id));
  }, [stations, selectedStationIds]);

  const triggerWlsAnalysis = async () => {
    if (activeStations.length < 2) {
      setWlsError('Select at least 2 DME stations for 3D analysis');
      setWlsSolution(null);
      setLogs(['[ERROR] Cannot execute WLS analysis: Select at least 2 DME stations.']);
      return;
    }

    const newLogs: string[] = [];
    newLogs.push(`[SYSTEM] Initializing 3D Navigation Accuracy Evaluation`);
    newLogs.push(
      `[SYSTEM] Aircraft Position: Easting = ${acX} NM, Northing = ${acY} NM, Altitude = ${acAlt.toLocaleString()} FT`
    );
    newLogs.push(`[SYSTEM] Altimeter Mode: ${altMode}`);
    newLogs.push(`[SYSTEM] Active DME Ground Beacons count: ${activeStations.length}`);

    // Map station positions to slant range and azimuth measurements
    const measurements: Measurement[] = activeStations.map((s) => {
      const dx = s.x - acX;
      const dy = s.y - acY;
      const r2D = Math.sqrt(dx * dx + dy * dy);

      // slant range height difference conversion (feet to NM)
      const dhFt = acAlt - s.elevationFt;
      const dhMeters = dhFt * 0.3048;
      const dhNM = dhMeters / 1852.0;

      const slantRange = Math.sqrt(r2D * r2D + dhNM * dhNM);

      // Calculate observation elevation angle
      let elDeg = 0;
      if (r2D > 1e-9) {
        elDeg = (Math.atan2(dhNM, r2D) * 180) / Math.PI;
      } else {
        elDeg = dhNM >= 0 ? 90 : -90;
      }

      // Calculate azimuth angle relative to aircraft
      let azDeg = (Math.atan2(dx, dy) * 180) / Math.PI;
      if (azDeg < 0) azDeg += 360;

      // AC-91-FS error model
      const airborneErr = Math.max(0.05, 0.00125 * slantRange);
      const totalErr = Math.sqrt(0.05 * 0.05 + airborneErr * airborneErr);

      // Log station evaluation
      const statusVolume = r2D <= s.serviceVolumeNM ? 'OK' : 'EXCEEDED';
      newLogs.push(
        `[DME][${s.id}] Distance = ${r2D.toFixed(2)} NM (Service Radius Limit: ${s.serviceVolumeNM} NM) -> ${statusVolume}`
      );
      newLogs.push(
        `[DME][${s.id}] Slant Range = ${slantRange.toFixed(2)} NM, Elevation Angle = ${elDeg.toFixed(2)}°, Azimuth = ${azDeg.toFixed(1)}°`
      );
      newLogs.push(`[DME][${s.id}] Expected 1σ Ranging Error = ${totalErr.toFixed(4)} NM`);

      return {
        azimuthDeg: azDeg,
        elevationDeg: elDeg,
        sigmaNM: totalErr,
      };
    });

    try {
      setWlsError(null);
      const response = await api.analyze3D({
        measurements,
        altitudeMode: altMode,
      });
      setWlsSolution(response);

      newLogs.push(`[WLS] Executing Weighted Least Squares algorithm...`);
      newLogs.push(
        `[WLS] Solver converged. Horizontal RMS: ${response.horizontalRmsNM.toFixed(4)} NM, Vertical RMS: ${response.verticalRmsNM.toFixed(4)} NM`
      );
      newLogs.push(`[WLS] Horizontal 2σ Position Error: ${response.twoSigmaNM.toFixed(4)} NM`);
      newLogs.push(
        `[DECISION] RNAV-1 Compliance (2σ <= 1.0 NM): ${response.rnav1 ? 'COMPLIANT' : 'NON-COMPLIANT'}`
      );
      newLogs.push(
        `[DECISION] RNAV-2 Compliance (2σ <= 2.0 NM): ${response.rnav2 ? 'COMPLIANT' : 'NON-COMPLIANT'}`
      );

      setLogs(newLogs);
    } catch (err) {
      const errMsg = err instanceof Error ? err.message : 'Failed to solve WLS';
      setWlsError(errMsg);
      setWlsSolution(null);

      newLogs.push(`[WLS] Executing Weighted Least Squares algorithm...`);
      newLogs.push(`[ERROR] WLS Solver failed to converge: ${errMsg}`);
      setLogs(newLogs);
    }
  };

  /* eslint-disable react-hooks/set-state-in-effect, react-hooks/exhaustive-deps */
  useEffect(() => {
    triggerWlsAnalysis();
  }, [acX, acY, acAlt, altMode, selectedStationIds, stations]);
  /* eslint-enable react-hooks/set-state-in-effect, react-hooks/exhaustive-deps */

  // Click Handler to relocate aircraft
  const handleMapClick = (nmX: number, nmY: number) => {
    setLogs([]); // Clear logs on map click
    setAcX(nmX);
    setAcY(nmY);
  };

  // Toggle DME selection
  const handleToggleStation = (id: string) => {
    setSelectedStationIds((prev) => {
      if (prev.includes(id)) {
        if (prev.length <= 2) return prev; // Keep at least 2 active stations
        return prev.filter((x) => x !== id);
      }
      return [...prev, id];
    });
  };

  // Add a new station
  const handleAddStation = (newStation: StationDTO) => {
    setStations((prev) => [...prev, newStation]);
    setSelectedStationIds((prev) => [...prev, newStation.id]);
  };

  // Remove station
  const handleRemoveStation = (id: string) => {
    setStations((prev) => prev.filter((s) => s.id !== id));
    setSelectedStationIds((prev) => prev.filter((x) => x !== id));
  };

  return (
    <div className="flex-1 w-full bg-slate-950 text-slate-100 flex flex-col gap-6 p-6">
      <div className="flex-1 flex flex-col md:flex-row gap-6">
        {/* LEFT COLUMN: Controls & DME Stations Panel */}
        <div className="flex-1 flex flex-col gap-6 md:max-w-md">
          <ScenarioParameters
            altMode={altMode}
            setAltMode={setAltMode}
            acAlt={acAlt}
            setAcAlt={setAcAlt}
            acX={acX}
            acY={acY}
          />
          <GroundBeacons
            stations={stations}
            selectedStationIds={selectedStationIds}
            onToggleStation={handleToggleStation}
            onAddStation={handleAddStation}
            onRemoveStation={handleRemoveStation}
            loading={loading}
          />
        </div>

        {/* RIGHT COLUMN: Tabs Panel */}
        <div className="flex-1 flex flex-col gap-6">
          <Tabs defaultValue="solver" className="w-full flex-1 flex flex-col">
            <TabsList className="bg-slate-900/50 border border-slate-850 p-1 rounded-xl self-start mb-4 backdrop-blur-md">
              <TabsTrigger
                value="solver"
                className="data-[state=active]:bg-cyan-600 data-[state=active]:text-slate-950 font-bold rounded-lg px-4 py-2 text-xs flex items-center gap-1.5 transition-all"
              >
                <Navigation className="w-3.5 h-3.5" />
                Real-time Solver
              </TabsTrigger>
              <TabsTrigger
                value="grid"
                className="data-[state=active]:bg-emerald-600 data-[state=active]:text-slate-950 font-bold rounded-lg px-4 py-2 text-xs flex items-center gap-1.5 transition-all"
              >
                <Grid3X3 className="w-3.5 h-3.5" />
                Grid Coverage
              </TabsTrigger>
              <TabsTrigger
                value="sweep"
                className="data-[state=active]:bg-indigo-600 data-[state=active]:text-slate-950 font-bold rounded-lg px-4 py-2 text-xs flex items-center gap-1.5 transition-all"
              >
                <Layers className="w-3.5 h-3.5" />
                Elevation Sweeps
              </TabsTrigger>
            </TabsList>

            {/* TAB 1: REAL-TIME WLS SOLVER VISUALIZER */}
            <TabsContent value="solver" className="flex-1 flex flex-col">
              <RealTimeSolver
                stations={stations}
                selectedStationIds={selectedStationIds}
                activeStations={activeStations}
                acX={acX}
                acY={acY}
                wlsSolution={wlsSolution}
                wlsError={wlsError}
                onMapClick={handleMapClick}
              />
            </TabsContent>

            {/* TAB 2: 2D GRID COVERAGE SIMULATOR */}
            <TabsContent value="grid" className="flex-1 flex flex-col">
              <GridCoverage stations={stations} />
            </TabsContent>

            {/* TAB 3: ELEVATION SWEEP ACCURACY CURVES */}
            <TabsContent value="sweep" className="flex-1 flex flex-col">
              <ElevationSweeps altMode={altMode} />
            </TabsContent>
          </Tabs>
        </div>
      </div>

      {/* Logger Panel */}
      <Logger
        logs={logs}
        acX={acX}
        acY={acY}
        acAlt={acAlt}
        altMode={altMode}
        wlsSolution={wlsSolution}
      />
    </div>
  );
}
