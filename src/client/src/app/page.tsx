import Link from 'next/link';
import {
  Navigation,
  Plane,
  ShieldCheck,
  BarChart3,
  Layers,
  Radio,
  Cpu,
  ArrowRight,
  FileCheck,
  Globe2,
  Activity,
  Copyright,
  ExternalLink,
} from 'lucide-react';

export default function LandingPage() {
  return (
    <div className="relative min-h-screen w-full bg-[#080C14] text-slate-100 flex flex-col overflow-x-hidden font-sans">
      {/* Meta Ambient Background Glows */}
      <div className="absolute top-0 left-1/2 -translate-x-1/2 w-[1000px] h-[500px] bg-gradient-to-b from-blue-600/15 via-indigo-600/10 to-transparent rounded-full blur-[140px] pointer-events-none" />
      <div className="absolute top-[35%] right-0 w-[600px] h-[600px] bg-cyan-500/10 rounded-full blur-[160px] pointer-events-none" />
      <div className="absolute bottom-[20%] left-0 w-[600px] h-[600px] bg-blue-600/10 rounded-full blur-[160px] pointer-events-none" />

      {/* Top Navigation Header */}
      <header className="sticky top-0 z-50 w-full border-b border-slate-800/80 bg-[#080C14]/80 backdrop-blur-xl px-6 py-3.5 flex items-center justify-between transition-all">
        <div className="flex items-center gap-3">
          <div className="w-10 h-10 rounded-xl bg-blue-600/15 border border-blue-500/30 flex items-center justify-center text-blue-400 shadow-lg shadow-blue-600/20">
            <Navigation className="w-5 h-5 animate-pulse" />
          </div>
          <div>
            <span className="text-base font-bold tracking-tight text-white flex items-center gap-2">
              PRISM{' '}
              <span className="text-[10px] px-2 py-0.5 rounded-md bg-blue-600/20 text-blue-400 border border-blue-500/30 font-mono">
                v1.0.0
              </span>
            </span>
            <p className="text-[11px] text-slate-400 font-medium">DME/DME RNAV Engine</p>
          </div>
        </div>

        <nav className="hidden md:flex items-center gap-8 text-xs font-medium text-slate-300">
          <a href="#features" className="hover:text-blue-400 transition-colors">
            Features
          </a>
          <a href="#architecture" className="hover:text-blue-400 transition-colors">
            3D WLS Solver
          </a>
          <a href="#database" className="hover:text-blue-400 transition-colors">
            Indian DME Network
          </a>
          <a href="#standards" className="hover:text-blue-400 transition-colors">
            Compliance
          </a>
        </nav>

        <div className="flex items-center gap-4">
          <Link
            href="/dashboard"
            className="meta-gradient-button px-5 py-2 rounded-xl text-xs font-semibold text-white flex items-center gap-2"
          >
            Launch Workspace
            <ArrowRight className="w-3.5 h-3.5" />
          </Link>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-1 flex flex-col items-center z-10">
        {/* HERO SECTION */}
        <section className="w-full max-w-6xl px-6 pt-16 pb-12 text-center flex flex-col items-center">
          {/* Meta Pill Badge */}
          <div className="inline-flex items-center gap-2.5 px-4 py-1.5 rounded-full border border-blue-500/30 bg-blue-950/30 text-blue-300 text-xs font-medium tracking-wide mb-8 backdrop-blur-md shadow-lg shadow-blue-950/30">
            <Cpu className="w-4 h-4 text-blue-400 animate-spin duration-[10000ms]" />
            <span>D2-GCA Navigation Research • 3D WLS Positioning System</span>
          </div>

          {/* Hero Headline */}
          <h1 className="text-4xl sm:text-6xl md:text-7xl font-extrabold tracking-tight meta-gradient-text max-w-4xl mb-6 leading-[1.1] drop-shadow-md">
            Next-Gen 3D Air Navigation & Positioning Engine
          </h1>

          <p className="text-base sm:text-xl text-slate-300 max-w-3xl mb-10 leading-relaxed font-normal">
            An industry-grade 3D weighted least squares (WLS) solver and geographic coverage
            simulator complying with FAA & ICAO{' '}
            <span className="text-blue-400 font-semibold">RNAV-1</span> and{' '}
            <span className="text-cyan-400 font-semibold">RNAV-2</span> area navigation standards.
          </p>

          {/* Action CTAs */}
          <div className="flex flex-col sm:flex-row items-center gap-4 mb-16 w-full sm:w-auto">
            <Link
              href="/dashboard"
              className="meta-gradient-button px-8 py-4 rounded-xl text-sm font-bold text-white flex items-center justify-center gap-3 w-full sm:w-auto shadow-2xl shadow-blue-600/30 group"
            >
              <Plane className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
              Enter Navigation Workspace
              <ArrowRight className="w-4 h-4" />
            </Link>
            <a
              href="#features"
              className="px-8 py-4 rounded-xl text-sm font-semibold text-slate-200 bg-slate-900/60 hover:bg-slate-800/80 border border-slate-700/60 transition-all duration-200 w-full sm:w-auto flex items-center justify-center gap-2"
            >
              Explore Capabilities
            </a>
          </div>

          {/* Hero Interactive Canvas Showcase Container */}
          <div className="w-full max-w-4xl meta-glass-card rounded-2xl p-4 sm:p-6 border border-slate-700/50 shadow-2xl shadow-blue-950/40 relative overflow-hidden group">
            {/* Top Bar of Window */}
            <div className="flex items-center justify-between pb-4 mb-4 border-b border-slate-800 text-xs text-slate-400">
              <div className="flex items-center gap-2">
                <span className="w-3 h-3 rounded-full bg-red-500/80 inline-block" />
                <span className="w-3 h-3 rounded-full bg-yellow-500/80 inline-block" />
                <span className="w-3 h-3 rounded-full bg-green-500/80 inline-block" />
                <span className="font-mono text-[11px] text-slate-400 ml-2">
                  PRISM Airspace Grid Simulator — Interactive Live State
                </span>
              </div>
              <div className="flex items-center gap-2 text-blue-400 font-mono text-[11px]">
                <Activity className="w-3.5 h-3.5 animate-pulse" />
                Live Solver Mode: RVSM Altimeter
              </div>
            </div>

            {/* Visual Radar Mockup */}
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-left">
              {/* Radar Graphic */}
              <div className="md:col-span-2 relative h-[260px] bg-slate-950/90 rounded-xl border border-slate-800/80 p-4 flex items-center justify-center overflow-hidden">
                {/* Radar Grid Crosshairs */}
                <div className="absolute inset-0 border border-slate-800/60 rounded-xl" />
                <div className="absolute w-full h-[1px] bg-slate-800/80 top-1/2 left-0" />
                <div className="absolute h-full w-[1px] bg-slate-800/80 left-1/2 top-0" />
                <div className="absolute w-48 h-48 rounded-full border border-blue-500/20 border-dashed" />
                <div className="absolute w-80 h-80 rounded-full border border-blue-500/15 border-dashed" />

                {/* Simulated Beacons */}
                <div className="absolute top-[25%] left-[30%] flex items-center gap-1.5">
                  <span className="w-2.5 h-2.5 rounded-full bg-emerald-400 animate-ping absolute" />
                  <span className="w-2.5 h-2.5 rounded-full bg-emerald-400 relative" />
                  <span className="text-[10px] font-mono font-bold text-emerald-400 bg-emerald-950/60 px-1.5 py-0.5 rounded border border-emerald-500/30">
                    NGP (Nagpur)
                  </span>
                </div>
                <div className="absolute bottom-[30%] right-[25%] flex items-center gap-1.5">
                  <span className="w-2.5 h-2.5 rounded-full bg-emerald-400 relative" />
                  <span className="text-[10px] font-mono font-bold text-emerald-400 bg-emerald-950/60 px-1.5 py-0.5 rounded border border-emerald-500/30">
                    GDA (Gondia)
                  </span>
                </div>
                <div className="absolute top-[40%] right-[35%] flex items-center gap-1.5">
                  <span className="w-2.5 h-2.5 rounded-full bg-emerald-400 relative" />
                  <span className="text-[10px] font-mono font-bold text-emerald-400 bg-emerald-950/60 px-1.5 py-0.5 rounded border border-emerald-500/30">
                    BPL (Bhopal)
                  </span>
                </div>

                {/* Simulated Aircraft */}
                <div className="absolute top-[48%] left-[52%] flex flex-col items-center">
                  <div className="w-12 h-12 rounded-full border border-amber-500/40 bg-amber-500/10 flex items-center justify-center animate-pulse">
                    <Plane className="w-5 h-5 text-amber-400 transform rotate-45" />
                  </div>
                  <span className="text-[10px] font-mono font-bold text-amber-400 bg-amber-950/80 px-2 py-0.5 rounded border border-amber-500/30 mt-1">
                    AC-TRUE (10 NM, 20 NM)
                  </span>
                </div>
              </div>

              {/* Sidebar Quick Solver Metrics */}
              <div className="flex flex-col gap-3 justify-between bg-slate-900/60 rounded-xl p-4 border border-slate-800">
                <div>
                  <span className="text-[10px] font-bold text-slate-400 uppercase tracking-wider">
                    Computed 2σ Horizontal Error
                  </span>
                  <div className="text-3xl font-mono font-extrabold text-blue-400 mt-1">
                    0.0482 <span className="text-xs text-slate-400 font-sans">NM</span>
                  </div>
                </div>

                <div className="space-y-2 pt-2 border-t border-slate-800">
                  <div className="flex justify-between items-center text-xs">
                    <span className="text-slate-400 font-medium">Terminal (RNAV-1)</span>
                    <span className="px-2 py-0.5 rounded-full bg-emerald-950/40 text-emerald-400 border border-emerald-500/30 text-[10px] font-bold">
                      COMPLIANT
                    </span>
                  </div>
                  <div className="flex justify-between items-center text-xs">
                    <span className="text-slate-400 font-medium">En-Route (RNAV-2)</span>
                    <span className="px-2 py-0.5 rounded-full bg-emerald-950/40 text-emerald-400 border border-emerald-500/30 text-[10px] font-bold">
                      COMPLIANT
                    </span>
                  </div>
                </div>

                <div className="pt-2 border-t border-slate-800 text-[11px] text-slate-400 font-mono">
                  Horizontal RMS: <span className="text-slate-200 font-bold">0.0241 NM</span>
                  <br />
                  Vertical RMS: <span className="text-slate-200 font-bold">0.0108 NM</span>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* METRICS & PROOF BANNER */}
        <section className="w-full border-y border-slate-800/80 bg-slate-900/30 backdrop-blur-md py-8">
          <div className="max-w-6xl mx-auto px-6 grid grid-cols-2 md:grid-cols-4 gap-6 text-center">
            <div className="p-4">
              <div className="text-3xl sm:text-4xl font-extrabold font-mono text-white mb-1">
                88+
              </div>
              <p className="text-xs text-slate-400 font-medium uppercase tracking-wider">
                Indian DME Beacons Loaded
              </p>
            </div>
            <div className="p-4 border-l border-slate-800/80">
              <div className="text-3xl sm:text-4xl font-extrabold font-mono text-blue-400 mb-1">
                &lt; 0.05 NM
              </div>
              <p className="text-xs text-slate-400 font-medium uppercase tracking-wider">
                Mean 2σ WLS Precision
              </p>
            </div>
            <div className="p-4 border-l border-slate-800/80">
              <div className="text-3xl sm:text-4xl font-extrabold font-mono text-emerald-400 mb-1">
                100%
              </div>
              <p className="text-xs text-slate-400 font-medium uppercase tracking-wider">
                Statutory 3D WLS Solver
              </p>
            </div>
            <div className="p-4 border-l border-slate-800/80">
              <div className="text-3xl sm:text-4xl font-extrabold font-mono text-cyan-400 mb-1">
                RNAV 1/2
              </div>
              <p className="text-xs text-slate-400 font-medium uppercase tracking-wider">
                ICAO / FAA Standard
              </p>
            </div>
          </div>
        </section>

        {/* FEATURE CARDS GRID */}
        <section id="features" className="w-full max-w-6xl px-6 py-20">
          <div className="text-center mb-16">
            <h2 className="text-3xl sm:text-4xl font-bold tracking-tight text-white mb-4">
              Comprehensive Navigation Engineering Suite
            </h2>
            <p className="text-base text-slate-400 max-w-2xl mx-auto">
              Built on mathematical formulations from cutting-edge navigation literature, enabling
              real-time WLS calculations and geographic coverage simulation.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {/* Card 1 */}
            <div className="meta-glass-card meta-glass-card-hover p-6 rounded-2xl text-left flex flex-col justify-between">
              <div>
                <div className="w-12 h-12 rounded-xl bg-blue-600/15 text-blue-400 border border-blue-500/30 flex items-center justify-center mb-5 shadow-md shadow-blue-950/40">
                  <Plane className="w-6 h-6" />
                </div>
                <h3 className="text-lg font-bold text-white mb-2">3D WLS Positioning Engine</h3>
                <p className="text-sm text-slate-400 leading-relaxed mb-4">
                  Calculates precise 3D aircraft position errors resolving barometric altimeter
                  sigmas (RVSM/CVSM) and DME slant-range intersection geometry.
                </p>
              </div>
              <div className="text-xs text-blue-400 font-mono font-semibold flex items-center gap-1 pt-2">
                Eq. 19 & 21 Solver <ArrowRight className="w-3.5 h-3.5" />
              </div>
            </div>

            {/* Card 2 */}
            <div className="meta-glass-card meta-glass-card-hover p-6 rounded-2xl text-left flex flex-col justify-between">
              <div>
                <div className="w-12 h-12 rounded-xl bg-cyan-600/15 text-cyan-400 border border-cyan-500/30 flex items-center justify-center mb-5 shadow-md shadow-cyan-950/40">
                  <BarChart3 className="w-6 h-6" />
                </div>
                <h3 className="text-lg font-bold text-white mb-2">Geographic Grid Coverage</h3>
                <p className="text-sm text-slate-400 leading-relaxed mb-4">
                  Simulates coverage grids across customizable boundaries to evaluate RNAV-1 and
                  RNAV-2 compliance percentages and select optimal beacon pairs.
                </p>
              </div>
              <div className="text-xs text-cyan-400 font-mono font-semibold flex items-center gap-1 pt-2">
                2D/3D Grid Sweep <ArrowRight className="w-3.5 h-3.5" />
              </div>
            </div>

            {/* Card 3 */}
            <div className="meta-glass-card meta-glass-card-hover p-6 rounded-2xl text-left flex flex-col justify-between">
              <div>
                <div className="w-12 h-12 rounded-xl bg-indigo-600/15 text-indigo-400 border border-indigo-500/30 flex items-center justify-center mb-5 shadow-md shadow-indigo-950/40">
                  <Layers className="w-6 h-6" />
                </div>
                <h3 className="text-lg font-bold text-white mb-2">Elevation Sweep Curves</h3>
                <p className="text-sm text-slate-400 leading-relaxed mb-4">
                  Plots analytical horizontal 2σ position accuracy vs. observation elevation angles
                  to replicate published research error curves (Figures 5, 6, 7).
                </p>
              </div>
              <div className="text-xs text-indigo-400 font-mono font-semibold flex items-center gap-1 pt-2">
                Recharts Analytics <ArrowRight className="w-3.5 h-3.5" />
              </div>
            </div>

            {/* Card 4 */}
            <div className="meta-glass-card meta-glass-card-hover p-6 rounded-2xl text-left flex flex-col justify-between">
              <div>
                <div className="w-12 h-12 rounded-xl bg-emerald-600/15 text-emerald-400 border border-emerald-500/30 flex items-center justify-center mb-5 shadow-md shadow-emerald-950/40">
                  <Radio className="w-6 h-6" />
                </div>
                <h3 className="text-lg font-bold text-white mb-2">Indian DME Network Database</h3>
                <p className="text-sm text-slate-400 leading-relaxed mb-4">
                  Pre-loaded dataset of 88 AAI DME ground stations (Nagpur, Delhi, Mumbai,
                  Bengaluru, Chennai, Kolkata, Leh) projected into projected Cartesian coordinates.
                </p>
              </div>
              <div className="text-xs text-emerald-400 font-mono font-semibold flex items-center gap-1 pt-2">
                88 Stations Pre-loaded <ArrowRight className="w-3.5 h-3.5" />
              </div>
            </div>

            {/* Card 5 */}
            <div className="meta-glass-card meta-glass-card-hover p-6 rounded-2xl text-left flex flex-col justify-between">
              <div>
                <div className="w-12 h-12 rounded-xl bg-purple-600/15 text-purple-400 border border-purple-500/30 flex items-center justify-center mb-5 shadow-md shadow-purple-950/40">
                  <Globe2 className="w-6 h-6" />
                </div>
                <h3 className="text-lg font-bold text-white mb-2">Interactive SVG Radar Grid</h3>
                <p className="text-sm text-slate-400 leading-relaxed mb-4">
                  Click anywhere on the Cartesian airspace canvas to move aircraft coordinates,
                  update active DME range circles, and visualize covariance uncertainty ellipses.
                </p>
              </div>
              <div className="text-xs text-purple-400 font-mono font-semibold flex items-center gap-1 pt-2">
                Click-to-Relocate Aircraft <ArrowRight className="w-3.5 h-3.5" />
              </div>
            </div>

            {/* Card 6 */}
            <div className="meta-glass-card meta-glass-card-hover p-6 rounded-2xl text-left flex flex-col justify-between">
              <div>
                <div className="w-12 h-12 rounded-xl bg-amber-600/15 text-amber-400 border border-amber-500/30 flex items-center justify-center mb-5 shadow-md shadow-amber-950/40">
                  <FileCheck className="w-6 h-6" />
                </div>
                <h3 className="text-lg font-bold text-white mb-2">Audit-Grade Export & PDF</h3>
                <p className="text-sm text-slate-400 leading-relaxed mb-4">
                  Generates statutory compliance PDF documentation formatted with official reference
                  headers, timestamps, and step-by-step solver evaluation traces.
                </p>
              </div>
              <div className="text-xs text-amber-400 font-mono font-semibold flex items-center gap-1 pt-2">
                Official Compliance Export <ArrowRight className="w-3.5 h-3.5" />
              </div>
            </div>
          </div>
        </section>

        {/* STANDARDS BANNER */}
        <section id="standards" className="w-full max-w-6xl px-6 pb-20">
          <div className="meta-glass-card rounded-2xl p-8 sm:p-10 border border-slate-800 flex flex-col md:flex-row items-center justify-between gap-8 text-left relative overflow-hidden">
            <div className="max-w-xl">
              <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-blue-950/50 border border-blue-500/30 text-blue-400 text-xs font-semibold uppercase tracking-wider mb-4">
                <ShieldCheck className="w-4 h-4" /> Statutory Aviation Compliance
              </div>
              <h3 className="text-2xl sm:text-3xl font-bold text-white mb-3">
                Complies with ICAO & FAA RNAV Standards
              </h3>
              <p className="text-sm text-slate-300 leading-relaxed">
                PRISM implements the exact accuracy evaluation model defined in AC-91-FS for DME
                ranging errors, alongside barometric altimeter tolerances for RVSM and CVSM
                operational airspace.
              </p>
            </div>
            <Link
              href="/dashboard"
              className="meta-gradient-button px-7 py-3.5 rounded-xl text-sm font-bold text-white flex items-center gap-2 shrink-0"
            >
              Launch Workspace Now
              <ArrowRight className="w-4 h-4" />
            </Link>
          </div>
        </section>
      </main>

      {/* Footer */}
      <footer className="w-full border-t border-slate-800/80 bg-[#080C14] py-8 text-xs text-slate-400">
        <div className="max-w-6xl mx-auto px-6 flex flex-col sm:flex-row items-center justify-between gap-4">
          <div className="flex items-center gap-2">
            <Copyright className="w-4 h-4 text-slate-500" />
            <span>
              2026 D2-GCA Organization. Licensed under{' '}
              <a
                href="https://apache.org"
                target="_blank"
                rel="noopener noreferrer"
                className="text-blue-400 hover:underline"
              >
                Apache License, Version 2.0
              </a>
            </span>
          </div>
          <div className="flex items-center gap-6">
            <Link href="/dashboard" className="hover:text-white transition-colors">
              Workspace
            </Link>
            <a
              href="https://github.com/D2-GCA/PRISM"
              target="_blank"
              rel="noopener noreferrer"
              className="hover:text-white transition-colors flex items-center gap-1"
            >
              GitHub <ExternalLink className="w-3 h-3" />
            </a>
          </div>
        </div>
      </footer>
    </div>
  );
}
