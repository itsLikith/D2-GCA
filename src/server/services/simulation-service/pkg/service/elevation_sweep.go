package service

import (
	"fmt"

	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
	"github.com/D2-GCA/PRISM/pkg/simulation/elevation"
	simtypes "github.com/D2-GCA/PRISM/pkg/simulation/types"

	"github.com/D2-GCA/PRISM/simulation-service/pkg/dto"
)

// RunElevationAngleSweep executes an elevation angle sweep simulation,
// computing horizontal accuracy across elevation and inclusion angles.
//
// Reference: IDEA.md §3.2–3.4 — Produces data for Figures 5–7,
// showing how horizontal accuracy varies with observation elevation
// angle for each altitude mode (RVSM, Fixed, CVSM).
func RunElevationAngleSweep(
	req *dto.ElevationSweepRequest,
) (*dto.ElevationSweepResponse, error) {

	altitudeMode, altitudeSigmaNM, err := resolveAltitudeMode(req.AltitudeMode)
	if err != nil {
		return nil, err
	}

	sweepConfig := simtypes.ElevationSweepConfig{
		DMESigma1NM:           req.Sigma1NM,
		DMESigma2NM:           req.Sigma2NM,
		AltitudeSensorSigmaNM: altitudeSigmaNM,
		AltitudeMode:          altitudeMode,
		InclusionAnglesDeg:    req.InclusionAnglesDeg,
		ElevationMinDeg:       req.ElevationMinDeg,
		ElevationMaxDeg:       req.ElevationMaxDeg,
		ElevationStepDeg:      req.ElevationStepDeg,
	}

	sweepResult, err := elevation.RunElevationAngleSweep(sweepConfig)
	if err != nil {
		return nil, err
	}

	return &dto.ElevationSweepResponse{
		AltitudeMode: sweepResult.AltitudeMode,
		Points:       sweepResult.Points,
	}, nil
}

// resolveAltitudeMode maps an altitude mode string to the corresponding
// AltitudeMode enum and altitude sensor sigma value.
//
// Reference: IDEA.md §3.1 — The three altitude sources have different
// error characteristics that affect the 3D positioning accuracy.
func resolveAltitudeMode(
	modeString string,
) (types.AltitudeMode, float64, error) {

	switch modeString {
	case "RVSM":
		// §3.2: RVSM barometric altimeter, σ₃ = 20m.
		return types.AltitudeModeRVSM, analytical.RVSMSigmaNM, nil

	case "FIXED":
		// §3.3: Fixed altitude, σ₃ = 0.
		return types.AltitudeModeFixed, analytical.FixedAltitudeSigmaNM, nil

	case "CVSM":
		// §3.4: CVSM barometric altimeter, σ₃ = 50m.
		return types.AltitudeModeCVSM, analytical.CVSMSigmaNM, nil

	default:
		return 0, 0, fmt.Errorf(
			"unknown altitude mode: %s (must be RVSM, FIXED, or CVSM)",
			modeString,
		)
	}
}
