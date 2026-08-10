import React, { useState, useEffect, useMemo } from 'react';
import { RefreshCw } from 'lucide-react';
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from 'recharts';
import { api, SweepPoint } from '@/lib/api';
import { Card, CardDescription, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';

interface ElevationSweepsProps {
  altMode: string;
}

export default function ElevationSweeps({ altMode }: ElevationSweepsProps) {
  // Elevation Sweep state
  const [sweepSigma1, setSweepSigma1] = useState<number>(0.0986);
  const [sweepSigma2, setSweepSigma2] = useState<number>(0.0986);
  const [sweepPoints, setSweepPoints] = useState<SweepPoint[]>([]);
  const [sweepLoading, setSweepLoading] = useState<boolean>(false);

  /* eslint-disable react-hooks/set-state-in-effect */
  const [mounted, setMounted] = useState<boolean>(false);
  useEffect(() => {
    setMounted(true);
  }, []);
  /* eslint-enable react-hooks/set-state-in-effect */

  // Run elevation sweep simulation
  const runSweepSimulation = async () => {
    setSweepLoading(true);
    try {
      const response = await api.runElevationSweep({
        sigma1NM: sweepSigma1,
        sigma2NM: sweepSigma2,
        altitudeMode: altMode,
        inclusionAnglesDeg: [30, 60, 90, 120, 150],
        elevationMinDeg: 0,
        elevationMaxDeg: 30,
        elevationStepDeg: 2,
      });
      setSweepPoints(response.points || []);
    } catch (err) {
      console.error(err);
    } finally {
      setSweepLoading(false);
    }
  };

  /* eslint-disable react-hooks/set-state-in-effect, react-hooks/exhaustive-deps */
  useEffect(() => {
    runSweepSimulation();
  }, [sweepSigma1, sweepSigma2, altMode]);
  /* eslint-enable react-hooks/set-state-in-effect, react-hooks/exhaustive-deps */

  // Translate sweep points to line chart schema
  const chartData = useMemo(() => {
    const dataMap: Record<number, Record<string, number>> = {};
    sweepPoints.forEach((p) => {
      if (!dataMap[p.elevationDeg]) {
        dataMap[p.elevationDeg] = { elevation: p.elevationDeg };
      }
      dataMap[p.elevationDeg][`angle_${p.inclusionAngleDeg}`] = Number(
        p.horizontalRms2SigmaNM.toFixed(4)
      );
    });
    return Object.values(dataMap).sort((a, b) => a.elevation - b.elevation);
  }, [sweepPoints]);

  return (
    <Card className="flex-1 border-slate-800 bg-slate-900/10 backdrop-blur-md text-slate-100 flex flex-col p-6 shadow-xl shadow-slate-950/20">
      <div className="pb-4 border-b border-slate-850 flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <div>
          <CardTitle className="text-base font-bold text-slate-200">
            Elevation Sweep Curves
          </CardTitle>
          <CardDescription className="text-xs text-slate-400">
            Replicates paper Figures 5, 6, and 7 showing 2σ horizontal error vs observation
            elevation angle
          </CardDescription>
        </div>

        {/* Ranging Error Inputs */}
        <div className="flex items-center gap-4 text-xs font-mono">
          <div className="flex items-center gap-2">
            <span className="text-slate-400 text-[10px] uppercase font-bold">DME 1 σ</span>
            <Input
              type="number"
              step="0.01"
              value={sweepSigma1}
              onChange={(e) => setSweepSigma1(Number(e.target.value))}
              className="bg-slate-950 border-slate-800 text-slate-200 w-20 h-7 text-xs font-bold text-center"
            />
          </div>
          <div className="flex items-center gap-2">
            <span className="text-slate-400 text-[10px] uppercase font-bold">DME 2 σ</span>
            <Input
              type="number"
              step="0.01"
              value={sweepSigma2}
              onChange={(e) => setSweepSigma2(Number(e.target.value))}
              className="bg-slate-950 border-slate-800 text-slate-200 w-20 h-7 text-xs font-bold text-center"
            />
          </div>
        </div>
      </div>

      {/* Recharts Elevation Sweep Chart */}
      <div className="flex-1 min-h-[350px] w-full mt-6">
        {sweepLoading ? (
          <div className="w-full h-full flex flex-col items-center justify-center gap-3">
            <RefreshCw className="w-6 h-6 text-indigo-400 animate-spin" />
            <span className="text-xs text-slate-500">Recalculating curves...</span>
          </div>
        ) : mounted ? (
          <ResponsiveContainer width="100%" height="100%">
            <LineChart data={chartData} margin={{ top: 10, right: 30, left: 0, bottom: 0 }}>
              <CartesianGrid strokeDasharray="3 3" stroke="#1E293B" />
              <XAxis
                dataKey="elevation"
                stroke="#475569"
                fontSize={11}
                tickFormatter={(v) => `${v}°`}
                label={{
                  value: 'Elevation Angle (θ)',
                  position: 'insideBottom',
                  offset: -5,
                  fill: '#64748B',
                  fontSize: 11,
                }}
              />
              <YAxis
                stroke="#475569"
                fontSize={11}
                label={{
                  value: '2σ Horizontal Accuracy (NM)',
                  angle: -90,
                  position: 'insideLeft',
                  offset: 12,
                  fill: '#64748B',
                  fontSize: 11,
                }}
              />
              <Tooltip
                contentStyle={{
                  backgroundColor: '#0F172A',
                  border: '1px solid #334155',
                  borderRadius: '8px',
                }}
                labelStyle={{ color: '#E2E8F0', fontSize: '11px', fontWeight: 'bold' }}
                itemStyle={{ fontSize: '11px' }}
              />
              <Legend
                verticalAlign="top"
                height={36}
                iconType="circle"
                wrapperStyle={{ fontSize: '11px' }}
              />
              <Line
                type="monotone"
                dataKey="angle_30"
                name="α = 30°"
                stroke="#EF4444"
                strokeWidth={1.5}
                dot={false}
                activeDot={{ r: 4 }}
              />
              <Line
                type="monotone"
                dataKey="angle_60"
                name="α = 60°"
                stroke="#F59E0B"
                strokeWidth={1.5}
                dot={false}
                activeDot={{ r: 4 }}
              />
              <Line
                type="monotone"
                dataKey="angle_90"
                name="α = 90°"
                stroke="#10B981"
                strokeWidth={1.5}
                dot={false}
                activeDot={{ r: 4 }}
              />
              <Line
                type="monotone"
                dataKey="angle_120"
                name="α = 120°"
                stroke="#06B6D4"
                strokeWidth={1.5}
                dot={false}
                activeDot={{ r: 4 }}
              />
              <Line
                type="monotone"
                dataKey="angle_150"
                name="α = 150°"
                stroke="#8B5CF6"
                strokeWidth={1.5}
                dot={false}
                activeDot={{ r: 4 }}
              />
            </LineChart>
          </ResponsiveContainer>
        ) : (
          <div className="w-full h-full flex flex-col items-center justify-center gap-3">
            <RefreshCw className="w-6 h-6 text-indigo-400 animate-spin" />
            <span className="text-xs text-slate-500">Loading chart...</span>
          </div>
        )}
      </div>
    </Card>
  );
}
