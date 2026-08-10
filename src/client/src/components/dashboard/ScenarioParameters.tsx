import React from 'react';
import { Settings, SlidersHorizontal } from 'lucide-react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Slider } from '@/components/ui/slider';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

interface ScenarioParametersProps {
  altMode: string;
  setAltMode: (mode: string) => void;
  acAlt: number;
  setAcAlt: (alt: number) => void;
  acX: number;
  acY: number;
}

export default function ScenarioParameters({
  altMode,
  setAltMode,
  acAlt,
  setAcAlt,
  acX,
  acY,
}: ScenarioParametersProps) {
  return (
    <Card className="meta-glass-card border-slate-800 text-slate-100 shadow-xl shadow-blue-950/20 rounded-2xl">
      <CardHeader className="pb-3 border-b border-slate-800/80">
        <CardTitle className="text-base font-bold flex items-center gap-2 text-blue-400">
          <div className="w-7 h-7 rounded-lg bg-blue-600/15 border border-blue-500/30 flex items-center justify-center text-blue-400">
            <Settings className="w-4 h-4" />
          </div>
          Scenario Parameters
        </CardTitle>
        <CardDescription className="text-slate-400 text-xs">
          Configure altitude sensor parameters and target flight levels
        </CardDescription>
      </CardHeader>
      <CardContent className="flex flex-col gap-5 pt-4">
        {/* Altitude Mode Select */}
        <div className="flex flex-col gap-2">
          <label className="text-[11px] font-bold uppercase tracking-wider text-slate-400 flex items-center gap-1.5">
            <SlidersHorizontal className="w-3.5 h-3.5 text-blue-400" />
            Altitude Sensor Mode
          </label>
          <Select value={altMode} onValueChange={setAltMode}>
            <SelectTrigger className="bg-slate-950/80 border-slate-800 text-slate-200 focus:ring-blue-600 h-9 rounded-xl text-xs">
              <SelectValue placeholder="Select mode" />
            </SelectTrigger>
            <SelectContent className="bg-slate-900 border-slate-800 text-slate-200 rounded-xl">
              <SelectItem value="RVSM">RVSM Barometric Altimeter (1σ = 20m)</SelectItem>
              <SelectItem value="CVSM">CVSM Barometric Altimeter (1σ = 50m)</SelectItem>
              <SelectItem value="FIXED">FIXED Precise Altitude (1σ = 0m)</SelectItem>
            </SelectContent>
          </Select>
        </div>

        {/* Aircraft Altitude Slider */}
        <div className="flex flex-col gap-2.5">
          <div className="flex justify-between items-center text-xs">
            <span className="font-semibold uppercase tracking-wider text-[11px] text-slate-400">
              Aircraft Altitude
            </span>
            <span className="text-blue-400 font-mono font-bold bg-blue-950/50 px-2 py-0.5 rounded border border-blue-500/30">
              {acAlt.toLocaleString()} FT
            </span>
          </div>
          <Slider
            min={2000}
            max={45000}
            step={500}
            value={[acAlt]}
            onValueChange={(val) => setAcAlt(val[0])}
            className="py-1 cursor-pointer"
          />
        </div>

        {/* Aircraft Position Display */}
        <div className="grid grid-cols-2 gap-3 pt-1">
          <div className="p-3 rounded-xl bg-slate-950/60 border border-slate-800/80 text-center">
            <div className="text-[10px] uppercase font-bold text-slate-500 tracking-wider">
              True East (X)
            </div>
            <div className="text-lg font-mono font-bold text-slate-100 mt-0.5">
              {acX} <span className="text-xs text-slate-500 font-sans">NM</span>
            </div>
          </div>
          <div className="p-3 rounded-xl bg-slate-950/60 border border-slate-800/80 text-center">
            <div className="text-[10px] uppercase font-bold text-slate-500 tracking-wider">
              True North (Y)
            </div>
            <div className="text-lg font-mono font-bold text-slate-100 mt-0.5">
              {acY} <span className="text-xs text-slate-500 font-sans">NM</span>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
