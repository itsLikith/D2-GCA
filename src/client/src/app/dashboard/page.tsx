import type { Metadata } from 'next';
import Link from 'next/link';
import { Navigation, ShieldCheck, Cpu, ArrowLeft, Radio } from 'lucide-react';
import Workspace from '@/components/dashboard/Workspace';

export const metadata: Metadata = {
  title: 'PRISM Navigation Workspace',
  description:
    'DME/DME RNAV Accuracy Analysis, WLS 3D Solver, and Geographic Grid Coverage Simulator.',
};

export default function DashboardPage() {
  return (
    <div className="flex flex-col min-h-screen w-full bg-[#080C14] text-slate-100 font-sans">
      {/* Top Navbar Header (Meta Glass style) */}
      <header className="sticky top-0 z-30 w-full border-b border-slate-800/80 bg-[#080C14]/85 backdrop-blur-xl px-6 py-3.5 flex items-center justify-between transition-all">
        {/* Brand Logo & Breadcrumb */}
        <div className="flex items-center gap-4">
          <Link
            href="/"
            className="w-9 h-9 rounded-xl bg-blue-600/15 text-blue-400 flex items-center justify-center border border-blue-500/30 shadow-md shadow-blue-950/30 hover:bg-blue-600/25 transition-all group"
            title="Return to Home"
          >
            <Navigation className="w-4 h-4 animate-pulse group-hover:scale-110 transition-transform" />
          </Link>
          <div className="flex items-center gap-2">
            <Link
              href="/"
              className="text-xs font-semibold text-slate-400 hover:text-slate-200 transition-colors flex items-center gap-1"
            >
              <ArrowLeft className="w-3 h-3" /> Home
            </Link>
            <span className="text-slate-600 text-xs">/</span>
            <div>
              <h1 className="text-sm font-bold tracking-tight text-white flex items-center gap-2">
                PRISM WORKSPACE
                <span className="text-[10px] px-2 py-0.5 rounded-md bg-blue-600/20 text-blue-400 border border-blue-500/30 font-mono">
                  3D WLS
                </span>
              </h1>
            </div>
          </div>
        </div>

        {/* Diagnostic Services Badges */}
        <div className="hidden sm:flex items-center gap-3 text-xs">
          <div className="flex items-center gap-2 px-3 py-1 rounded-xl bg-slate-900/80 border border-slate-800 text-slate-300">
            <Cpu className="w-3.5 h-3.5 text-blue-400 animate-spin duration-[8000ms]" />
            WLS Engine: <span className="text-emerald-400 font-bold font-mono">STATUTORY</span>
          </div>
          <div className="flex items-center gap-2 px-3 py-1 rounded-xl bg-slate-900/80 border border-slate-800 text-slate-300">
            <Radio className="w-3.5 h-3.5 text-cyan-400" />
            Network: <span className="text-cyan-400 font-bold font-mono">88 DME BEACONS</span>
          </div>
          <div className="flex items-center gap-2 px-3 py-1 rounded-xl bg-slate-900/80 border border-slate-800 text-slate-300">
            <ShieldCheck className="w-3.5 h-3.5 text-indigo-400" />
            Specs: <span className="text-indigo-400 font-bold font-mono">RNAV 1/2</span>
          </div>
        </div>
      </header>

      {/* Main Workspace Dashboard Content */}
      <main className="flex-1 flex flex-col min-h-0">
        <Workspace />
      </main>
    </div>
  );
}
