import React from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { FileDown, Terminal } from 'lucide-react';
import { Analyze3DResponse } from '@/lib/api';

interface LoggerProps {
  logs: string[];
  acX: number;
  acY: number;
  acAlt: number;
  altMode: string;
  wlsSolution: Analyze3DResponse | null;
}

export default function Logger({ logs, acX, acY, acAlt, altMode, wlsSolution }: LoggerProps) {
  const handlePrintPDF = () => {
    const printWindow = window.open('', '_blank');
    if (!printWindow) return;

    const dateStr = new Date().toLocaleString();
    const logText = logs.join('\n');
    const year = new Date().getFullYear();
    const randId = Math.floor(1000 + Math.random() * 9000);
    const reportRefFilename = `AAI_CNS_PRISM_${year}_${randId}`;
    const reportRefDisplay = `AAI/CNS/PRISM/${year}/${randId}`;
    const htmlContent = `
      <html>
        <head>
          <title>${reportRefFilename}</title>
          <style>
            @page {
              size: A4;
              margin: 15mm 20mm 20mm 20mm;
            }
            body {
              font-family: 'Georgia', 'Times New Roman', serif;
              color: #1e293b;
              line-height: 1.5;
              font-size: 11px;
              margin: 0;
              padding: 0;
            }
            .watermark {
              position: fixed;
              top: 50%;
              left: 50%;
              transform: translate(-50%, -50%) rotate(-45deg);
              font-size: 120px;
              font-weight: bold;
              color: rgba(12, 36, 69, 0.050);
              z-index: -1000;
              pointer-events: none;
              white-space: nowrap;
              user-select: none;
              font-family: 'Helvetica Neue', Arial, sans-serif;
            }
            .classification {
              text-align: center;
              font-family: 'Arial', sans-serif;
              font-size: 9px;
              font-weight: bold;
              letter-spacing: 2px;
              color: #9d783d;
              margin-bottom: 10px;
              text-transform: uppercase;
            }
            .header-container {
              display: flex;
              align-items: center;
              justify-content: flex-start;
              border-bottom: 2px solid #0c2445;
              padding-bottom: 10px;
              margin-bottom: 15px;
            }
            .logo-img {
              height: 150px;
              width: auto;
              margin-right: 24px;
            }
            .header-text {
              text-align: left;
            }
            .org-hindi {
              font-family: 'Nirmala UI', 'Mangal', sans-serif;
              font-size: 22px;
              font-weight: bold;
              color: #0c2445;
              line-height: 1.2;
            }
            .org-english {
              font-family: 'Arial', sans-serif;
              font-size: 16px;
              font-weight: 800;
              letter-spacing: 0.5px;
              color: #0c2445;
              line-height: 1.2;
              margin-top: 2px;
            }
            .dept-title {
              font-family: 'Arial', sans-serif;
              font-size: 10px;
              font-weight: bold;
              color: #9d783d;
              text-transform: uppercase;
              letter-spacing: 1px;
              margin-top: 6px;
            }
            .report-info {
              display: flex;
              justify-content: space-between;
              font-size: 9px;
              color: #64748b;
              margin-bottom: 15px;
              font-family: 'Arial', sans-serif;
              border-bottom: 1px dashed #cbd5e1;
              padding-bottom: 5px;
            }
            .report-title-container {
              text-align: center;
              margin-bottom: 20px;
            }
            .report-title {
              display: inline-block;
              font-family: 'Arial', sans-serif;
              font-size: 13px;
              font-weight: bold;
              color: #ffffff;
              background-color: #0c2445;
              padding: 6px 16px;
              text-transform: uppercase;
              letter-spacing: 1px;
              border-radius: 2px;
              border-bottom: 2px solid #9d783d;
            }
            .meta-table {
              width: 100%;
              border-collapse: collapse;
              margin-bottom: 20px;
              font-family: 'Arial', sans-serif;
              font-size: 10px;
            }
            .meta-table td {
              padding: 6px 10px;
              border: 1px solid #cbd5e1;
            }
            .meta-table td.label {
              font-weight: bold;
              background-color: #f8fafc;
              width: 25%;
              color: #0c2445;
            }
            .meta-table td.val {
              color: #334155;
            }
            h2 {
              font-family: 'Arial', sans-serif;
              font-size: 11px;
              color: #0c2445;
              border-bottom: 1px solid #9d783d;
              padding-bottom: 3px;
              margin-top: 20px;
              margin-bottom: 10px;
              text-transform: uppercase;
              letter-spacing: 0.5px;
              font-weight: bold;
            }
            .log-box {
              background-color: #f8fafc;
              border: 1px solid #e2e8f0;
              border-radius: 4px;
              padding: 12px;
              font-family: 'Courier New', Courier, monospace;
              white-space: pre-wrap;
              font-size: 10px;
              color: #334155;
              line-height: 1.4;
              margin-bottom: 25px;
            }
            .sign-block {
              margin-top: 40px;
              display: flex;
              justify-content: space-between;
              font-family: 'Arial', sans-serif;
              font-size: 9px;
            }
            .sign-col {
              text-align: center;
              width: 30%;
            }
            .sign-line {
              border-top: 1px solid #475569;
              margin-top: 35px;
              padding-top: 4px;
              color: #475569;
              font-weight: bold;
            }
            .footer {
              position: fixed;
              bottom: 0;
              left: 0;
              right: 0;
              text-align: center;
              font-family: 'Arial', sans-serif;
              font-size: 8px;
              color: #94a3b8;
              border-top: 1px solid #e2e8f0;
              padding-top: 5px;
            }
            @media print {
              .sign-block {
                page-break-inside: avoid;
              }
              .meta-table {
                page-break-inside: avoid;
              }
              .log-box {
                page-break-inside: auto;
              }
            }
          </style>
        </head>
        <body>
          <div class="watermark">PRISM</div>
          <div class="classification">RESTRICTED - FOR OFFICIAL USE ONLY</div>
          
          <div class="header-container">
            <img class="logo-img" src="${window.location.origin}/AAI.webp" alt="AAI Logo" />
            <div class="header-text">
              <div class="org-hindi">भारतीय विमानपत्तन प्राधिकरण</div>
              <div class="org-english">AIRPORTS AUTHORITY OF INDIA</div>
              <div class="dept-title">CNS / Aviation Services Division — Corporate Headquarters</div>
            </div>
          </div>

          <div class="report-info">
            <div><strong>REPORT REF:</strong> ${reportRefDisplay}</div>
            <div><strong>DATE:</strong> ${dateStr}</div>
          </div>

          <div class="report-title-container">
            <div class="report-title">DME/DME RNAV ACCURACY ANALYSIS CERTIFICATE</div>
          </div>

          <table class="meta-table">
            <tr>
              <td class="label">Project Evaluation</td>
              <td class="val">PRISM Space Navigation Evaluator</td>
              <td class="label">Altitude Mode</td>
              <td class="val">${altMode} Sensor Model</td>
            </tr>
            <tr>
              <td class="label">Aircraft Position</td>
              <td class="val">Easting: ${acX} NM, Northing: ${acY} NM</td>
              <td class="label">Aircraft Altitude</td>
              <td class="val">${acAlt.toLocaleString()} FT</td>
            </tr>
            <tr>
              <td class="label">Terminal Area (RNAV-1)</td>
              <td class="val" style="font-weight: bold; color: ${wlsSolution?.rnav1 ? '#10b981' : '#ef4444'}">
                ${wlsSolution?.rnav1 ? 'COMPLIANT' : 'NON-COMPLIANT'}
              </td>
              <td class="label">En-route Area (RNAV-2)</td>
              <td class="val" style="font-weight: bold; color: ${wlsSolution?.rnav2 ? '#10b981' : '#ef4444'}">
                ${wlsSolution?.rnav2 ? 'COMPLIANT' : 'NON-COMPLIANT'}
              </td>
            </tr>
            <tr>
              <td class="label">2σ Position Error</td>
              <td class="val" style="font-weight: bold; font-family: monospace; color: #0c2445;">
                ${wlsSolution ? wlsSolution.twoSigmaNM.toFixed(4) : 'N/A'} NM
              </td>
              <td class="label">Evaluation Authority</td>
              <td class="val">CNS Dept, AAI</td>
            </tr>
          </table>

          <h2>Operational Telemetry & Analysis Logs</h2>
          <div class="log-box">${logText}</div>

          <div class="sign-block">
            <div class="sign-col">
              <div class="sign-line">CNS SYSTEM ADMINISTRATOR</div>
              <div>PRISM Computation Engine</div>
            </div>
            <div class="sign-col">
              <div class="sign-line">VALIDATED BY</div>
              <div>CNS Shift In-Charge</div>
            </div>
            <div class="sign-col">
              <div class="sign-line">AUTHORIZED SIGNATORY</div>
              <div>Joint General Manager, CNS</div>
            </div>
          </div>

          <div class="footer">
            RESTRICTED — AAI CNS DIVISION — GENERATED AUTONOMOUSLY BY PRISM SYSTEMS
          </div>

          <script>
            window.onload = function() {
              window.print();
              setTimeout(function() { window.close(); }, 500);
            };
          </script>
        </body>
      </html>
    `;

    printWindow.document.write(htmlContent);
    printWindow.document.close();
  };

  return (
    <Card className="border-slate-800 bg-slate-900/30 backdrop-blur-md text-slate-100 shadow-xl shadow-slate-950/20 relative w-full">
      <CardHeader className="pb-3 flex flex-row items-center justify-between">
        <div>
          <CardTitle className="text-sm font-bold text-slate-200 flex items-center gap-2">
            <Terminal className="w-4 h-4 text-cyan-400" />
            Airspace Operations & Decisions Logger
          </CardTitle>
          <CardDescription className="text-xs text-slate-400">
            Real-time CNS telemetry logs and compliance details for the present point
          </CardDescription>
        </div>
        <Button
          size="sm"
          onClick={handlePrintPDF}
          disabled={logs.length === 0}
          className="bg-cyan-600 hover:bg-cyan-500 text-slate-950 font-bold gap-1.5 rounded-lg border border-cyan-500/20"
        >
          <FileDown className="w-4 h-4" /> Download PDF Report
        </Button>
      </CardHeader>
      <CardContent>
        <div className="h-44 overflow-y-auto rounded-lg bg-slate-950/60 border border-slate-850 p-4 font-mono text-xs text-emerald-400 flex flex-col gap-1 leading-relaxed">
          {logs.map((log, index) => (
            <div key={index}>{log}</div>
          ))}
          {logs.length === 0 && (
            <div className="text-slate-600 italic">
              No logs generated. Click on the map to evaluate a point.
            </div>
          )}
        </div>
      </CardContent>
    </Card>
  );
}
