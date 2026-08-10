package service

import (
	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"
	"github.com/D2-GCA/PRISM/pkg/navigation/engine"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"

	"github.com/D2-GCA/PRISM/rnav-service/pkg/dto"
)

// RunRNAV3DAnalysis performs a 3D DME/DME RNAV positioning analysis
// with altitude sensor integration.
//
// Reference: IDEA.md §3.1 — Extends the 2D analysis to three dimensions
// by incorporating altitude information from one of three sources:
//   - RVSM barometric altimeter (§3.2, σ₃ = 20m)
//   - Fixed altitude flight    (§3.3, σ₃ = 0)
//   - CVSM barometric altimeter (§3.4, σ₃ = 50m)
//
// The 3D G matrix (Eq. 19/22) and extended covariance (Eq. 20) are used
// to compute both horizontal σ_AZ and vertical σ_EL accuracy (Eq. 21).
func RunRNAV3DAnalysis(req *dto.Analyze3DRequest) (*dto.Analyze3DResponse, error) {
	measurements := make([]types.Measurement, 0, len(req.Measurements))

	for _, m := range req.Measurements {
		measurements = append(measurements, types.Measurement{
			AzimuthDeg:   m.AzimuthDeg,
			ElevationDeg: m.ElevationDeg,
			SigmaNM:      m.SigmaNM,
		})
	}

	// Resolve altitude sensor error based on the operating mode (§3.1).
	var altitudeSigmaNM float64
	switch req.AltitudeMode {
	case "RVSM":
		altitudeSigmaNM = analytical.RVSMSigmaNM
	case "CVSM":
		altitudeSigmaNM = analytical.CVSMSigmaNM
	case "FIXED":
		// For FIXED mode, σ₃ is theoretically 0 (§3.3).
		// Use a negligible epsilon for numerical stability in W = diag(1/σ²).
		altitudeSigmaNM = 1e-6
	default:
		altitudeSigmaNM = analytical.RVSMSigmaNM
	}

	solution, err := engine.Compute3DNavigationSolution(
		measurements,
		altitudeSigmaNM,
		true, // Include altimeter pseudo-observation (Eq. 22).
	)
	if err != nil {
		return nil, err
	}

	return &dto.Analyze3DResponse{
		HorizontalRMSNM: solution.RMSAccuracyNM,
		VerticalRMSNM:   solution.VerticalAccuracyNM,
		TwoSigmaNM:      solution.TwoSigmaNM,
		AltitudeMode:    req.AltitudeMode,
		RNAV1:           solution.RNAV1,
		RNAV2:           solution.RNAV2,
	}, nil
}
