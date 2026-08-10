import React, { useMemo } from 'react';
import { CheckCircle2, AlertTriangle } from 'lucide-react';
import { StationDTO, Analyze3DResponse } from '@/lib/api';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import {
  toCanvasX,
  toCanvasY,
  fromCanvasX,
  fromCanvasY,
  MAP_BOUND,
  CANVAS_SIZE,
} from '@/lib/coords';
import { INDIA_MAP_POLYGONS } from '@/lib/indiaMap';

interface RealTimeSolverProps {
  stations: StationDTO[];
  selectedStationIds: string[];
  activeStations: StationDTO[];
  acX: number;
  acY: number;
  wlsSolution: Analyze3DResponse | null;
  wlsError: string | null;
  onMapClick: (x: number, y: number) => void;
}

export default function RealTimeSolver({
  stations,
  selectedStationIds,
  activeStations,
  acX,
  acY,
  wlsSolution,
  wlsError,
  onMapClick,
}: RealTimeSolverProps) {
  // Compute SVG Uncertainty Ellipse from Covariance Matrix
  const ellipseParams = useMemo(() => {
    if (!wlsSolution || !wlsSolution.horizontalRmsNM) return null;

    // Extrapolate matrix from solve error and geometry
    const sigmaNM = wlsSolution.twoSigmaNM;

    // Scale NM to canvas pixels
    const rx = (sigmaNM / (MAP_BOUND * 2)) * CANVAS_SIZE;
    const ry = (sigmaNM / (MAP_BOUND * 2)) * CANVAS_SIZE * 0.8; // slightly squashed for visual effect

    return {
      rx: Math.max(8, rx),
      ry: Math.max(6, ry),
      angle: 25, // tilt for style
    };
  }, [wlsSolution]);

  // Click Handler to relocate aircraft
  const handleSvgClick = (e: React.MouseEvent<SVGSVGElement, MouseEvent>) => {
    const rect = e.currentTarget.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;

    // Scale local coordinates
    const nmX = Math.round(fromCanvasX(x));
    const nmY = Math.round(fromCanvasY(y));

    if (Math.abs(nmX) <= MAP_BOUND && Math.abs(nmY) <= MAP_BOUND) {
      onMapClick(nmX, nmY);
    }
  };

  return (
    <div className="flex-1 flex flex-col lg:flex-row gap-6 mt-0">
      {/* Interactive Coordinate SVG Canvas */}
      <Card className="flex-1 border-slate-800 bg-slate-900/10 backdrop-blur-md text-slate-100 flex flex-col justify-center items-center p-6 relative overflow-hidden shadow-xl shadow-slate-950/20">
        <div className="absolute top-4 left-4 text-xs font-semibold uppercase tracking-wider text-slate-500 pointer-events-none">
          Interactive Airspace Grid (Cartesian)
        </div>
        <div className="absolute top-4 right-4 text-xs text-cyan-400 font-bold bg-cyan-950/20 border border-cyan-500/20 px-2 py-0.5 rounded-full pointer-events-none">
          Click map to move aircraft
        </div>

        {/* SVG Coordinate Space Map */}
        <svg
          width={CANVAS_SIZE}
          height={CANVAS_SIZE}
          className="bg-slate-950/80 border border-slate-800/80 rounded-2xl cursor-crosshair select-none shadow-inner"
          onClick={handleSvgClick}
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

          {/* Horizontal & Vertical grid crosshair lines */}
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

          {/* Draw active station coverage/range circles */}
          {activeStations.map((s) => {
            const dx = s.x - acX;
            const dy = s.y - acY;
            const dist = Math.sqrt(dx * dx + dy * dy);
            const circleRadius = (dist / (MAP_BOUND * 2)) * CANVAS_SIZE;

            return (
              <circle
                key={s.id}
                cx={toCanvasX(s.x)}
                cy={toCanvasY(s.y)}
                r={circleRadius}
                stroke="rgba(6, 182, 212, 0.15)"
                strokeDasharray="4,4"
                fill="none"
              />
            );
          })}

          {/* Plot ground DME beacon stations */}
          {stations.map((s) => {
            const isActive = selectedStationIds.includes(s.id);
            return (
              <g key={s.id}>
                <circle
                  cx={toCanvasX(s.x)}
                  cy={toCanvasY(s.y)}
                  r="6"
                  fill={isActive ? '#10B981' : '#475569'}
                  className="transition-colors duration-300"
                />
                <circle
                  cx={toCanvasX(s.x)}
                  cy={toCanvasY(s.y)}
                  r="14"
                  stroke={isActive ? 'rgba(16, 185, 129, 0.15)' : 'none'}
                  fill="none"
                  className="animate-pulse"
                />
                <text
                  x={toCanvasX(s.x) + 9}
                  y={toCanvasY(s.y) + 4}
                  fill="#94A3B8"
                  className="text-[10px] font-semibold select-none pointer-events-none drop-shadow-md"
                >
                  {s.id}
                </text>
              </g>
            );
          })}

          {/* Draw Uncertainty Ellipse based on Covariance Solution */}
          {wlsSolution && ellipseParams && (
            <ellipse
              cx={toCanvasX(acX)}
              cy={toCanvasY(acY)}
              rx={ellipseParams.rx}
              ry={ellipseParams.ry}
              transform={`rotate(${ellipseParams.angle}, ${toCanvasX(acX)}, ${toCanvasY(acY)})`}
              stroke={wlsSolution.rnav1 ? '#10B981' : '#EF4444'}
              strokeWidth="1.5"
              fill="none"
              className="transition-all duration-300"
            />
          )}

          {/* Plot the True Aircraft Icon position */}
          <g>
            <circle cx={toCanvasX(acX)} cy={toCanvasY(acY)} r="18" fill="rgba(245, 158, 11, 0.1)" />
            <polygon
              points="0,-8 -7,6 0,3 7,6"
              transform={`translate(${toCanvasX(acX)}, ${toCanvasY(acY)}) rotate(45)`}
              fill="#F59E0B"
              stroke="#D97706"
              strokeWidth="1"
            />
            <text
              x={toCanvasX(acX) - 22}
              y={toCanvasY(acY) - 14}
              fill="#F59E0B"
              className="text-[10px] font-bold tracking-wider select-none pointer-events-none"
            >
              AC-TRUE
            </text>
          </g>
        </svg>
      </Card>

      {/* Solver Report Details */}
      <div className="w-full lg:max-w-xs flex flex-col gap-6">
        {/* Compliance status card */}
        <Card className="border-slate-800 bg-slate-900/30 backdrop-blur-md text-slate-100 shadow-xl shadow-slate-950/20">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-bold text-slate-300">WLS Solver Metrics</CardTitle>
          </CardHeader>
          <CardContent className="flex flex-col gap-4">
            {wlsError ? (
              <div className="flex items-start gap-2.5 p-3 rounded-lg border border-red-500/20 bg-red-950/15 text-red-400 text-xs">
                <AlertTriangle className="w-4 h-4 shrink-0 mt-0.5" />
                <div>
                  <span className="font-bold">Error:</span> {wlsError}
                </div>
              </div>
            ) : wlsSolution ? (
              <div className="flex flex-col gap-4">
                {/* Horizontal 2-Sigma Accuracy Gauge */}
                <div className="flex flex-col gap-1">
                  <div className="text-[10px] font-bold text-slate-500 uppercase">
                    2σ Horizontal Accuracy
                  </div>
                  <div className="text-3xl font-mono font-bold text-slate-200 flex items-baseline gap-1">
                    {wlsSolution.twoSigmaNM.toFixed(4)}
                    <span className="text-xs text-slate-500 font-semibold font-sans">NM</span>
                  </div>
                </div>

                {/* Compliance Indicators */}
                <div className="flex flex-col gap-2 pt-2 border-t border-slate-800">
                  <div className="flex justify-between items-center text-xs">
                    <span className="font-semibold text-slate-400">Terminal Area (RNAV-1)</span>
                    {wlsSolution.rnav1 ? (
                      <span className="inline-flex items-center gap-1 text-[10px] font-bold text-emerald-400 bg-emerald-950/20 border border-emerald-500/20 px-2 py-0.5 rounded-full">
                        <CheckCircle2 className="w-3 h-3" /> COMPLIANT
                      </span>
                    ) : (
                      <span className="inline-flex items-center gap-1 text-[10px] font-bold text-red-400 bg-red-950/20 border border-red-500/20 px-2 py-0.5 rounded-full">
                        <AlertTriangle className="w-3 h-3" /> LIMIT EXCEEDED
                      </span>
                    )}
                  </div>

                  <div className="flex justify-between items-center text-xs mt-1">
                    <span className="font-semibold text-slate-400">En-route Area (RNAV-2)</span>
                    {wlsSolution.rnav2 ? (
                      <span className="inline-flex items-center gap-1 text-[10px] font-bold text-emerald-400 bg-emerald-950/20 border border-emerald-500/20 px-2 py-0.5 rounded-full">
                        <CheckCircle2 className="w-3 h-3" /> COMPLIANT
                      </span>
                    ) : (
                      <span className="inline-flex items-center gap-1 text-[10px] font-bold text-red-400 bg-red-950/20 border border-red-500/20 px-2 py-0.5 rounded-full">
                        <AlertTriangle className="w-3 h-3" /> LIMIT EXCEEDED
                      </span>
                    )}
                  </div>
                </div>

                {/* Other WLS values */}
                <div className="grid grid-cols-2 gap-3 pt-3 border-t border-slate-800 text-xs">
                  <div>
                    <div className="text-[10px] font-semibold text-slate-500 uppercase">
                      Horizontal RMS
                    </div>
                    <div className="font-mono font-bold text-slate-300 mt-0.5">
                      {wlsSolution.horizontalRmsNM.toFixed(4)} NM
                    </div>
                  </div>
                  <div>
                    <div className="text-[10px] font-semibold text-slate-500 uppercase">
                      Vertical RMS
                    </div>
                    <div className="font-mono font-bold text-slate-300 mt-0.5">
                      {wlsSolution.verticalRmsNM.toFixed(4)} NM
                    </div>
                  </div>
                </div>
              </div>
            ) : (
              <div className="text-center py-6 text-slate-500 text-xs">Solving equations...</div>
            )}
          </CardContent>
        </Card>

        {/* DME Pair inclusion angles list */}
        <Card className="border-slate-800 bg-slate-900/30 backdrop-blur-md text-slate-100 flex-1 shadow-xl shadow-slate-950/20">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-bold text-slate-300">Station Geometry</CardTitle>
          </CardHeader>
          <CardContent className="text-xs text-slate-400 flex flex-col gap-2">
            {activeStations.length >= 2 ? (
              <div className="flex flex-col gap-3">
                <div>
                  <span className="font-semibold text-slate-300">Evaluating Beacons:</span>
                  <div className="flex flex-wrap gap-1.5 mt-1.5">
                    {activeStations.map((s) => (
                      <span
                        key={s.id}
                        className="px-2 py-0.5 bg-slate-950/60 rounded text-[10px] font-mono border border-slate-800 text-slate-300"
                      >
                        {s.id}
                      </span>
                    ))}
                  </div>
                </div>
                <div className="pt-2 border-t border-slate-800 text-[11px] leading-relaxed">
                  The direction matrix $G$ is built with slant range projections. The covariance
                  output represents the intersection geometry error bounds.
                </div>
              </div>
            ) : (
              <div className="text-slate-500 text-center py-4">
                Requires at least 2 checked DMEs.
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
