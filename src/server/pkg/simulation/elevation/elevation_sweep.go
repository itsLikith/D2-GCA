// Package elevation implements the elevation angle sweep simulation
// for analyzing how horizontal positioning accuracy varies with
// the DME observation elevation angle.
//
// Reference: IDEA.md §3 — This simulation produces the data for
// Figures 5–7, which show the 2σ horizontal accuracy vs. elevation
// angle for each altitude mode (RVSM, Fixed, CVSM).
package elevation

import (
	"github.com/D2-GCA/PRISM/pkg/navigation/accuracy"
	"github.com/D2-GCA/PRISM/pkg/simulation/types"
)

// RunElevationAngleSweep executes an elevation angle sweep simulation,
// computing the 2σ horizontal positioning accuracy at each combination
// of elevation angle and inclusion angle.
//
// Reference: IDEA.md §3.2–3.4 — For each (θ, α) pair, the closed-form
// horizontal accuracy (Eq. 23) is evaluated:
//
//	σ_AZ = √(σ₁²/cos²θ + σ₂²/cos²θ + σ₃²(2tan²θ - 2tan²θ·cosα)) / sinα
//
// The simulation assumes symmetric elevation (θ₁ = θ₂ = θ) as in the
// paper's analysis. The results show that:
//   - Lower elevation angles improve horizontal accuracy (§3.5)
//   - All three altitude modes meet RNAV-2 and RNAV-1 within the valid
//     inclusion angle range of 30°–150°
func RunElevationAngleSweep(config types.ElevationSweepConfig) (*types.ElevationSweepResult, error) {
	var sweepPoints []types.ElevationSweepPoint

	for elevationDeg := config.ElevationMinDeg; elevationDeg <= config.ElevationMaxDeg; elevationDeg += config.ElevationStepDeg {
		for _, inclusionAngleDeg := range config.InclusionAnglesDeg {

			// Evaluate closed-form Eq. 23 with symmetric elevation θ₁ = θ₂ = θ.
			rmsNM, err := accuracy.HorizontalAccuracy3DClosedForm(
				config.DMESigma1NM,
				config.DMESigma2NM,
				config.AltitudeSensorSigmaNM,
				elevationDeg,
				elevationDeg, // Symmetric: θ₁ = θ₂ = θ
				inclusionAngleDeg,
			)
			if err != nil {
				// Skip invalid geometry points (e.g., θ=90° or α=0°/180°).
				continue
			}

			twoSigmaNM := 2.0 * rmsNM

			sweepPoints = append(sweepPoints, types.ElevationSweepPoint{
				ElevationDeg:          elevationDeg,
				InclusionAngleDeg:     inclusionAngleDeg,
				HorizontalRMS1SigmaNM: rmsNM,
				HorizontalRMS2SigmaNM: twoSigmaNM,
			})
		}
	}

	return &types.ElevationSweepResult{
		AltitudeMode: config.AltitudeMode.String(),
		Points:       sweepPoints,
	}, nil
}
