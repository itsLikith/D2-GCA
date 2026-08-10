package service

import (
	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"

	"github.com/D2-GCA/PRISM/dme-service/pkg/dto"
)

// ComputeDMEPairAccuracy analyzes the horizontal positioning accuracy
// for a specific DME/DME station pair at a given inclusion angle.
//
// Reference: IDEA.md §2.3, Eq. 17/18 — Given two DME ranges, the
// function computes the total error for each station and applies the
// closed-form horizontal accuracy formula:
//
//	2σ = 2·√(σ₁² + σ₂²) / sin(α)
func ComputeDMEPairAccuracy(
	req *dto.PairAccuracyRequest,
) (*dto.PairAccuracyResponse, error) {

	// Compute total DME error for each station per AC-91-FS (§2.3).
	sigma1NM := analytical.TotalDMEError(req.Range1NM)
	sigma2NM := analytical.TotalDMEError(req.Range2NM)

	// Apply closed-form accuracy analysis (Eq. 17/18).
	analysisResult, err := analytical.AnalyzeStationPair(
		req.InclusionAngleDeg,
		sigma1NM,
		sigma2NM,
	)
	if err != nil {
		return nil, err
	}

	return &dto.PairAccuracyResponse{
		InclusionAngleDeg: req.InclusionAngleDeg,
		RMSNM:             analysisResult.RMSNM,
		TwoSigmaNM:        analysisResult.TwoSigmaNM,
		RNAV1:             analysisResult.RNAV1,
		RNAV2:             analysisResult.RNAV2,
	}, nil
}
