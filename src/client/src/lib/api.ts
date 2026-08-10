export interface Measurement {
  azimuthDeg: number;
  elevationDeg: number;
  sigmaNM: number;
}

export interface AnalyzeRequest {
  measurements: Measurement[];
  observations: number[];
}

export interface AnalyzeResponse {
  x: number;
  y: number;
  rmsAccuracyNM: number;
  twoSigmaNM: number;
  rnav1: boolean;
  rnav2: boolean;
}

export interface Analyze3DRequest {
  measurements: Measurement[];
  altitudeMode: string; // RVSM, FIXED, CVSM
}

export interface Analyze3DResponse {
  horizontalRmsNM: number;
  verticalRmsNM: number;
  twoSigmaNM: number;
  altitudeMode: string;
  rnav1: boolean;
  rnav2: boolean;
}

export interface StationDTO {
  id: string;
  name?: string;
  x: number;
  y: number;
  elevationFt: number;
  serviceVolumeNM: number;
}

export interface CoverageRequest {
  stations: StationDTO[];
  minX: number;
  maxX: number;
  minY: number;
  maxY: number;
  gridStepNM: number;
}

export interface PointResult {
  eastingNM: number;
  northingNM: number;
  inclusionAngleDeg: number;
  rmsAccuracyNM: number;
  twoSigmaNM: number;
  rnav1: boolean;
  rnav2: boolean;
}

export interface Statistics {
  totalPoints: number;
  analyzedPoints: number;
  rnav1Points: number;
  rnav2Points: number;
  rnav1CoveragePercent: number;
  rnav2CoveragePercent: number;
  averageRmsNM: number;
  averageTwoSigmaNM: number;
}

export interface CoverageResponse {
  points: PointResult[];
  statistics: Statistics;
}

export interface ElevationSweepRequest {
  sigma1NM: number;
  sigma2NM: number;
  altitudeMode: string;
  inclusionAnglesDeg: number[];
  elevationMinDeg: number;
  elevationMaxDeg: number;
  elevationStepDeg: number;
}

export interface SweepPoint {
  elevationDeg: number;
  inclusionAngleDeg: number;
  horizontalRms2SigmaNM: number;
}

export interface ElevationSweepResponse {
  altitudeMode: string;
  points: SweepPoint[];
}

const BASE_URL = '/api/v1';

async function post<T, R>(path: string, body: T): Promise<R> {
  const response = await fetch(`${BASE_URL}${path}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  });

  if (!response.ok) {
    let errorMsg = `API request failed with status ${response.status}`;
    try {
      const errJson = await response.json();
      if (errJson && errJson.message) {
        errorMsg = errJson.message;
      }
    } catch {
      // ignore
    }
    throw new Error(errorMsg);
  }

  return response.json() as Promise<R>;
}

export const api = {
  getStations: () =>
    fetch(`${BASE_URL}/dme/stations`).then((res) => {
      if (!res.ok) throw new Error('Failed to fetch stations');
      return res.json() as Promise<StationDTO[]>;
    }),
  analyze2D: (req: AnalyzeRequest) => post<AnalyzeRequest, AnalyzeResponse>('/rnav/analyze', req),
  analyze3D: (req: Analyze3DRequest) =>
    post<Analyze3DRequest, Analyze3DResponse>('/rnav/analyze3d', req),
  runCoverage: (req: CoverageRequest) =>
    post<CoverageRequest, CoverageResponse>('/simulation/coverage', req),
  runElevationSweep: (req: ElevationSweepRequest) =>
    post<ElevationSweepRequest, ElevationSweepResponse>('/simulation/elevation', req),
};
