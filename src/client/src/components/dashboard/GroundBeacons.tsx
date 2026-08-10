import React, { useState } from 'react';
import { Radio, Plus, Trash2, Search } from 'lucide-react';
import { StationDTO } from '@/lib/api';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';

interface GroundBeaconsProps {
  stations: StationDTO[];
  selectedStationIds: string[];
  onToggleStation: (id: string) => void;
  onAddStation: (station: StationDTO) => void;
  onRemoveStation: (id: string) => void;
  loading?: boolean;
}

export default function GroundBeacons({
  stations,
  selectedStationIds,
  onToggleStation,
  onAddStation,
  onRemoveStation,
  loading = false,
}: GroundBeaconsProps) {
  // Local form state for adding new DME station
  const [newId, setNewId] = useState<string>('');
  const [newX, setNewX] = useState<string>('');
  const [newY, setNewY] = useState<string>('');
  const [newEl, setNewEl] = useState<string>('1000');
  const [newSv, setNewSv] = useState<string>('130');
  const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false);
  const [searchQuery, setSearchQuery] = useState<string>('');

  const filteredStations = stations.filter((s) =>
    s.id.toLowerCase().includes(searchQuery.toLowerCase().trim())
  );

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newId.trim() || !newX.trim() || !newY.trim()) return;

    const xVal = parseFloat(newX);
    const yVal = parseFloat(newY);
    const elVal = parseFloat(newEl);
    const svVal = parseFloat(newSv);

    if (isNaN(xVal) || isNaN(yVal) || isNaN(elVal) || isNaN(svVal)) return;

    const newStation: StationDTO = {
      id: newId.trim(),
      x: xVal,
      y: yVal,
      elevationFt: elVal,
      serviceVolumeNM: svVal,
    };

    onAddStation(newStation);

    // Reset Form
    setNewId('');
    setNewX('');
    setNewY('');
    setNewEl('1000');
    setNewSv('130');
    setIsDialogOpen(false);
  };

  return (
    <Card className="border-slate-800 bg-slate-900/30 backdrop-blur-md text-slate-100 h-[450px] flex flex-col shadow-xl shadow-slate-950/20">
      <CardHeader className="pb-2 flex flex-row justify-between items-center">
        <div>
          <CardTitle className="text-base font-bold text-slate-200 flex items-center gap-2">
            <Radio className="w-4 h-4 text-emerald-400" />
            Ground DME Beacons
          </CardTitle>
          <CardDescription className="text-xs text-slate-400">
            Manage candidate stations and positioning selects
          </CardDescription>
        </div>

        {/* Add DME Dialog Trigger */}
        <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
          <DialogTrigger asChild>
            <Button
              size="sm"
              className="bg-cyan-600 hover:bg-cyan-500 text-slate-950 font-bold gap-1 rounded-lg"
            >
              <Plus className="w-4 h-4" /> Add
            </Button>
          </DialogTrigger>
          <DialogContent className="bg-slate-900 border-slate-800 text-slate-100 max-w-sm rounded-2xl">
            <DialogHeader>
              <DialogTitle className="text-slate-100">Add DME Station</DialogTitle>
              <DialogDescription className="text-slate-400 text-xs">
                Define coordinates, elevation, and service bounds.
              </DialogDescription>
            </DialogHeader>
            <form onSubmit={handleSubmit} className="flex flex-col gap-4 mt-2">
              <div className="flex flex-col gap-1.5">
                <label className="text-xs font-semibold text-slate-400">Station ID</label>
                <Input
                  placeholder="e.g. DME-01"
                  value={newId}
                  onChange={(e) => setNewId(e.target.value)}
                  className="bg-slate-950 border-slate-800 text-slate-200"
                />
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div className="flex flex-col gap-1.5">
                  <label className="text-xs font-semibold text-slate-400">X Coord (NM)</label>
                  <Input
                    placeholder="e.g. -45"
                    value={newX}
                    onChange={(e) => setNewX(e.target.value)}
                    className="bg-slate-950 border-slate-800 text-slate-200"
                  />
                </div>
                <div className="flex flex-col gap-1.5">
                  <label className="text-xs font-semibold text-slate-400">Y Coord (NM)</label>
                  <Input
                    placeholder="e.g. 60"
                    value={newY}
                    onChange={(e) => setNewY(e.target.value)}
                    className="bg-slate-950 border-slate-800 text-slate-200"
                  />
                </div>
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div className="flex flex-col gap-1.5">
                  <label className="text-xs font-semibold text-slate-400">Elevation (FT)</label>
                  <Input
                    placeholder="e.g. 1200"
                    value={newEl}
                    onChange={(e) => setNewEl(e.target.value)}
                    className="bg-slate-950 border-slate-800 text-slate-200"
                  />
                </div>
                <div className="flex flex-col gap-1.5">
                  <label className="text-xs font-semibold text-slate-400">Range Limit (NM)</label>
                  <Input
                    placeholder="e.g. 130"
                    value={newSv}
                    onChange={(e) => setNewSv(e.target.value)}
                    className="bg-slate-950 border-slate-800 text-slate-200"
                  />
                </div>
              </div>
              <Button
                type="submit"
                className="bg-emerald-500 hover:bg-emerald-400 text-slate-950 font-bold rounded-xl mt-2 w-full"
              >
                Confirm and Add
              </Button>
            </form>
          </DialogContent>
        </Dialog>
      </CardHeader>

      <div className="px-4 pb-1 relative flex items-center">
        <Search className="w-3.5 h-3.5 text-slate-500 absolute left-7 pointer-events-none" />
        <Input
          placeholder="Filter beacons by ID..."
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="bg-slate-950/60 border-slate-800 text-slate-200 h-8 text-xs placeholder:text-slate-500 rounded-lg pl-8 w-full"
        />
      </div>

      <CardContent className="flex-1 overflow-auto p-0 border-t border-slate-800/60 mt-3">
        <Table>
          <TableHeader className="bg-slate-950/30 sticky top-0 backdrop-blur-md">
            <TableRow className="border-slate-800/80 hover:bg-transparent">
              <TableHead className="w-12 text-center text-slate-400">Use</TableHead>
              <TableHead className="text-slate-400 text-xs uppercase font-bold">
                Beacon ID
              </TableHead>
              <TableHead className="text-slate-400 text-xs uppercase font-bold text-center">
                Coord (X,Y)
              </TableHead>
              <TableHead className="text-slate-400 text-xs uppercase font-bold text-center">
                El (FT)
              </TableHead>
              <TableHead className="w-12 text-slate-400 text-center"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {loading ? (
              <TableRow>
                <TableCell colSpan={5} className="text-center py-8 text-slate-500 text-xs">
                  Loading stations from database...
                </TableCell>
              </TableRow>
            ) : filteredStations.length === 0 ? (
              <TableRow>
                <TableCell colSpan={5} className="text-center py-8 text-slate-500 text-xs">
                  No stations found.
                </TableCell>
              </TableRow>
            ) : (
              filteredStations.map((s) => (
                <TableRow key={s.id} className="border-slate-800/40 hover:bg-slate-900/10">
                  <TableCell className="text-center p-2">
                    <input
                      type="checkbox"
                      checked={selectedStationIds.includes(s.id)}
                      onChange={() => onToggleStation(s.id)}
                      className="w-4 h-4 rounded border-slate-800 text-cyan-600 bg-slate-950 focus:ring-0 focus:ring-offset-0 cursor-pointer"
                    />
                  </TableCell>
                  <TableCell className="font-semibold text-slate-200 text-sm p-2">
                    <div className="flex flex-col">
                      <span>{s.id}</span>
                      {s.name && (
                        <span className="text-[10px] text-slate-500 font-normal">{s.name}</span>
                      )}
                    </div>
                  </TableCell>
                  <TableCell className="text-center font-mono text-xs text-slate-400 p-2">
                    {s.x}, {s.y}
                  </TableCell>
                  <TableCell className="text-center font-mono text-xs text-slate-400 p-2">
                    {s.elevationFt}
                  </TableCell>
                  <TableCell className="p-2 text-center">
                    <button
                      onClick={() => onRemoveStation(s.id)}
                      className="text-slate-500 hover:text-red-400 transition-colors p-1"
                      title="Delete Station"
                    >
                      <Trash2 className="w-4 h-4" />
                    </button>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  );
}
