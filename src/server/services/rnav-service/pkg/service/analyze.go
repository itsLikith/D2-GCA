package service

import (
	"github.com/D2-GCA/PRISM/pkg/navigation/engine"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"

	"github.com/D2-GCA/PRISM/rnav-service/pkg/dto"
)

// RunRNAV2DAnalysis performs a 2D DME/DME RNAV positioning analysis.
//
// Reference: IDEA.md §2 — Converts the API request into the internal
// measurement format and runs the full WLS pipeline:
// G matrix (Eq. 6) → W matrix (Eq. 13) → WLS solve (Eq. 8) →
// covariance (Eq. 14) → RMS accuracy (Eq. 12) → RNAV compliance.
func RunRNAV2DAnalysis(req *dto.AnalyzeRequest) (*dto.AnalyzeResponse, error) {
	measurements := make([]types.Measurement, 0, len(req.Measurements))

	for _, m := range req.Measurements {
		measurements = append(measurements, types.Measurement{
			AzimuthDeg: m.AzimuthDeg,
			SigmaNM:    m.SigmaNM,
		})
	}

	solution, err := engine.Compute2DNavigationSolution(
		measurements,
		req.Observations,
	)
	if err != nil {
		return nil, err
	}

	return &dto.AnalyzeResponse{
		X:             solution.PositionEstimate.EastingNM,
		Y:             solution.PositionEstimate.NorthingNM,
		RMSAccuracyNM: solution.RMSAccuracyNM,
		TwoSigmaNM:    solution.TwoSigmaNM,
		RNAV1:         solution.RNAV1,
		RNAV2:         solution.RNAV2,
	}, nil
}
