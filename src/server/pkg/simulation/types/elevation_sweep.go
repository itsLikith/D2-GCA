package types

import "github.com/D2-GCA/PRISM/pkg/navigation/types"

// ElevationSweepConfig defines the parameters for an elevation angle
// sweep simulation.
//
// Reference: IDEA.md §3.2–3.4 — The simulation evaluates the
// closed-form Eq. 23 across a range of elevation angles and
// inclusion angles for a given altitude mode.
type ElevationSweepConfig struct {
	// DMESigma1NM is the 1σ ranging error of DME station 1 (NM).
	DMESigma1NM float64

	// DMESigma2NM is the 1σ ranging error of DME station 2 (NM).
	DMESigma2NM float64

	// AltitudeSensorSigmaNM is the 1σ error of the altitude sensor (NM).
	// RVSM: ~0.0108 NM (20m), Fixed: 0, CVSM: ~0.0270 NM (50m).
	AltitudeSensorSigmaNM float64

	// AltitudeMode identifies which altitude source is being simulated.
	AltitudeMode types.AltitudeMode

	// InclusionAnglesDeg lists the inclusion angles α to evaluate (degrees).
	InclusionAnglesDeg []float64

	// ElevationMinDeg is the starting elevation angle θ (degrees).
	ElevationMinDeg float64

	// ElevationMaxDeg is the ending elevation angle θ (degrees).
	ElevationMaxDeg float64

	// ElevationStepDeg is the increment between elevation angles (degrees).
	ElevationStepDeg float64
}

// ElevationSweepResult holds the complete output of an elevation sweep.
type ElevationSweepResult struct {
	// AltitudeMode label (e.g., "RVSM", "FIXED", "CVSM").
	AltitudeMode string

	// Points contains the accuracy at each (elevation, inclusion angle) pair.
	Points []ElevationSweepPoint
}

// ElevationSweepPoint holds the accuracy at a single (θ, α) combination.
type ElevationSweepPoint struct {
	// ElevationDeg is the observation elevation angle θ (degrees).
	ElevationDeg float64 `json:"elevationDeg"`

	// InclusionAngleDeg is the inclusion angle α (degrees).
	InclusionAngleDeg float64 `json:"inclusionAngleDeg"`

	// HorizontalRMS1SigmaNM is the 1σ horizontal accuracy (NM).
	HorizontalRMS1SigmaNM float64 `json:"horizontalRms1SigmaNM"`

	// HorizontalRMS2SigmaNM is the 2σ (95%) horizontal accuracy (NM).
	HorizontalRMS2SigmaNM float64 `json:"horizontalRms2SigmaNM"`
}
