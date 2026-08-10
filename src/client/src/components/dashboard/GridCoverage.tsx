import React, { useState } from 'react';
import { RefreshCw } from 'lucide-react';
import { api, StationDTO, PointResult, Statistics } from '@/lib/api';
import { toCanvasX, toCanvasY, CANVAS_SIZE } from '@/lib/coords';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { INDIA_MAP_POLYGONS } from '@/lib/indiaMap';

interface GridCoverageProps {
  stations: StationDTO[];
}

export default function GridCoverage({ stations }: GridCoverageProps) {
  // Grid Simulation state
  const [simMinX, setSimMinX] = useState<number>(-400);
  const [simMaxX, setSimMaxX] = useState<number>(400);
  const [simMinY, setSimMinY] = useState<number>(-400);
  const [simMaxY, setSimMaxY] = useState<number>(400);
  const [simStep, setSimStep] = useState<number>(20);
  const [simResults, setSimResults] = useState<PointResult[]>([]);
  const [simStats, setSimStats] = useState<Statistics | null>(null);
  const [simLoading, setSimLoading] = useState<boolean>(false);
  const [simError, setSimError] = useState<string | null>(null);

  // Grid Simulation Runner
  const runGridSimulation = async () => {
    if (stations.length < 2) {
      setSimError('Create at least 2 ground stations to simulate');
      return;
    }
    setSimLoading(true);
    setSimError(null);
    try {
      const response = await api.runCoverage({
        stations,
        minX: simMinX,
        maxX: simMaxX,
        minY: simMinY,
        maxY: simMaxY,
        gridStepNM: simStep,
      });
      setSimResults(response.points || []);
      setSimStats(response.statistics || null);
    } catch (err) {
      const errMsg = err instanceof Error ? err.message : 'Simulation failed';
      setSimError(errMsg);
      setSimResults([]);
      setSimStats(null);
    } finally {
      setSimLoading(false);
    }
  };

  return (
    <div className="flex-1 flex flex-col lg:flex-row gap-6 mt-0">
      {/* Simulation Canvas Grid */}
      <Card className="flex-1 border-slate-800 bg-slate-900/10 backdrop-blur-md text-slate-100 flex flex-col justify-center items-center p-6 relative overflow-hidden shadow-xl shadow-slate-950/20">
        <div className="absolute top-4 left-4 text-xs font-semibold uppercase tracking-wider text-slate-500 pointer-events-none">
          Simulated Grid Coverage Plot
        </div>

        {simLoading ? (
          <div className="flex flex-col items-center justify-center gap-3 w-[500px] h-[500px] bg-slate-950/80 border border-slate-800/80 rounded-2xl">
            <RefreshCw className="w-8 h-8 text-cyan-400 animate-spin" />
            <span className="text-sm font-semibold text-slate-400">
              Executing Grid Sweep Calculations...
            </span>
          </div>
        ) : (
          <svg
            width={CANVAS_SIZE}
            height={CANVAS_SIZE}
            className="bg-slate-950/80 border border-slate-800/80 rounded-2xl shadow-inner"
          >
            {/* India Map Outline in background */}
            {INDIA_MAP_POLYGONS.map((polygon, pIdx) => {
              const pointsString = polygon
                .map((pt) => `${toCanvasX(pt.x)},${toCanvasY(pt.y)}`)
                .join(' ');
              return (
                <polyline
                  key={pIdx}
                  points={pointsString}
                  fill="rgba(148, 163, 184, 0.02)"
                  stroke="rgba(6, 182, 212, 0.15)"
                  strokeWidth="1.2"
                />
              );
            })}

            {/* Grid axis lines */}
            <line
              x1={0}
              y1={CANVAS_SIZE / 2}
              x2={CANVAS_SIZE}
              y2={CANVAS_SIZE / 2}
              stroke="#1E293B"
              strokeWidth="1.5"
            />
            <line
              x1={CANVAS_SIZE / 2}
              y1={0}
              x2={CANVAS_SIZE / 2}
              y2={CANVAS_SIZE}
              stroke="#1E293B"
              strokeWidth="1.5"
            />

            {/* Render coverage simulation dots */}
            {simResults.map((p, idx) => {
              let color = '#EF4444'; // Non-compliant
              if (p.rnav1)
                color = '#10B981'; // RNAV-1 Compliant
              else if (p.rnav2) color = '#06B6D4'; // RNAV-2 Compliant

              return (
                <circle
                  key={idx}
                  cx={toCanvasX(p.eastingNM)}
                  cy={toCanvasY(p.northingNM)}
                  r="3"
                  fill={color}
                  opacity="0.8"
                />
              );
            })}

            {/* Ground DME stations markers overlay */}
            {stations.map((s) => (
              <g key={s.id}>
                <polygon
                  points={`${toCanvasX(s.x)},${toCanvasY(s.y) - 6} ${toCanvasX(s.x) - 5},${toCanvasY(s.y) + 4} ${toCanvasX(s.x) + 5},${toCanvasY(s.y) + 4}`}
                  fill="#F59E0B"
                />
                <text
                  x={toCanvasX(s.x) + 7}
                  y={toCanvasY(s.y) + 3}
                  fill="#F59E0B"
                  className="text-[9px] font-bold select-none pointer-events-none"
                >
                  {s.id}
                </text>
              </g>
            ))}
          </svg>
        )}
      </Card>

      {/* Simulation controls & report */}
      <div className="w-full lg:max-w-xs flex flex-col gap-6">
        <Card className="border-slate-800 bg-slate-900/30 backdrop-blur-md text-slate-100 shadow-xl shadow-slate-950/20">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-bold text-slate-300">Sweep Parameters</CardTitle>
          </CardHeader>
          <CardContent className="flex flex-col gap-4">
            {/* Grid Boundary Slider Inputs */}
            <div className="flex flex-col gap-3.5">
              <div className="grid grid-cols-2 gap-3">
                <div className="flex flex-col gap-1">
                  <label className="text-[10px] font-bold text-slate-500 uppercase">
                    Min X (NM)
                  </label>
                  <Input
                    type="number"
                    value={simMinX}
                    onChange={(e) => setSimMinX(Number(e.target.value))}
                    className="bg-slate-950 border-slate-800 text-slate-200 text-xs font-mono h-8"
                  />
                </div>
                <div className="flex flex-col gap-1">
                  <label className="text-[10px] font-bold text-slate-500 uppercase">
                    Max X (NM)
                  </label>
                  <Input
                    type="number"
                    value={simMaxX}
                    onChange={(e) => setSimMaxX(Number(e.target.value))}
                    className="bg-slate-950 border-slate-800 text-slate-200 text-xs font-mono h-8"
                  />
                </div>
              </div>

              <div className="grid grid-cols-2 gap-3">
                <div className="flex flex-col gap-1">
                  <label className="text-[10px] font-bold text-slate-500 uppercase">
                    Min Y (NM)
                  </label>
                  <Input
                    type="number"
                    value={simMinY}
                    onChange={(e) => setSimMinY(Number(e.target.value))}
                    className="bg-slate-950 border-slate-800 text-slate-200 text-xs font-mono h-8"
                  />
                </div>
                <div className="flex flex-col gap-1">
                  <label className="text-[10px] font-bold text-slate-500 uppercase">
                    Max Y (NM)
                  </label>
                  <Input
                    type="number"
                    value={simMaxY}
                    onChange={(e) => setSimMaxY(Number(e.target.value))}
                    className="bg-slate-950 border-slate-800 text-slate-200 text-xs font-mono h-8"
                  />
                </div>
              </div>

              <div className="flex flex-col gap-1">
                <label className="text-[10px] font-bold text-slate-500 uppercase">
                  Resolution step (NM)
                </label>
                <Select value={String(simStep)} onValueChange={(val) => setSimStep(Number(val))}>
                  <SelectTrigger className="bg-slate-950 border-slate-800 text-slate-200 h-8 text-xs">
                    <SelectValue placeholder="Select step" />
                  </SelectTrigger>
                  <SelectContent className="bg-slate-900 border-slate-800 text-slate-200">
                    <SelectItem value="5">High density (5 NM)</SelectItem>
                    <SelectItem value="10">Standard density (10 NM)</SelectItem>
                    <SelectItem value="20">Low density (20 NM)</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>

            {simError && <div className="text-red-400 text-xs mt-1">{simError}</div>}

            <Button
              onClick={runGridSimulation}
              disabled={simLoading}
              className="bg-emerald-500 hover:bg-emerald-400 text-slate-950 font-bold rounded-xl w-full mt-2"
            >
              Run Simulation Grid
            </Button>
          </CardContent>
        </Card>

        {/* Simulation Statistics */}
        <Card className="border-slate-800 bg-slate-900/30 backdrop-blur-md text-slate-100 flex-1 shadow-xl shadow-slate-950/20">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-bold text-slate-300">Coverage Statistics</CardTitle>
          </CardHeader>
          <CardContent className="text-xs text-slate-400 flex flex-col gap-4">
            {simStats ? (
              <div className="flex flex-col gap-3">
                <div className="flex justify-between items-center">
                  <span className="font-semibold text-slate-400">RNAV-1 Coverage:</span>
                  <span className="text-emerald-400 font-bold font-mono text-sm">
                    {simStats.rnav1CoveragePercent.toFixed(1)}%
                  </span>
                </div>
                <div className="flex justify-between items-center pt-2 border-t border-slate-800/40">
                  <span className="font-semibold text-slate-400">RNAV-2 Coverage:</span>
                  <span className="text-cyan-400 font-bold font-mono text-sm">
                    {simStats.rnav2CoveragePercent.toFixed(1)}%
                  </span>
                </div>
                <div className="flex justify-between items-center pt-2 border-t border-slate-800/40">
                  <span className="font-semibold text-slate-400">Avg. 2σ Accuracy:</span>
                  <span className="text-slate-200 font-bold font-mono">
                    {simStats.averageTwoSigmaNM.toFixed(4)} NM
                  </span>
                </div>
                <div className="flex justify-between items-center pt-2 border-t border-slate-800/40">
                  <span className="font-semibold text-slate-400">Analyzed Grid Points:</span>
                  <span className="text-slate-200 font-mono">
                    {simStats.analyzedPoints} / {simStats.totalPoints}
                  </span>
                </div>
                <div className="flex gap-2 items-center text-[10px] mt-2 pt-2 border-t border-slate-800 text-slate-500 font-semibold leading-relaxed">
                  <div className="flex gap-1 items-center">
                    <span className="w-2.5 h-2.5 rounded-full bg-emerald-500 inline-block" /> RNAV-1
                  </div>
                  <div className="flex gap-1 items-center">
                    <span className="w-2.5 h-2.5 rounded-full bg-cyan-500 inline-block" /> RNAV-2
                  </div>
                  <div className="flex gap-1 items-center">
                    <span className="w-2.5 h-2.5 rounded-full bg-red-500 inline-block" /> FAIL
                  </div>
                </div>
              </div>
            ) : (
              <div className="text-slate-500 text-center py-6">
                Click the run button to generate coverage metrics.
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
