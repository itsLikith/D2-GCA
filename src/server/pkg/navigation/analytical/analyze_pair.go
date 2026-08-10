package analytical

import "github.com/D2-GCA/PRISM/pkg/navigation/types"

// AnalyzeStationPair performs the complete §2.3 accuracy analysis
// for a pair of DME stations at a given inclusion angle.
//
// Reference: IDEA.md §2.3 — This function:
//  1. Computes the 1σ RMS error using Eq. 17: σ = √(σ₁²+σ₂²)/sin(α)
//  2. Converts to 2σ for compliance: 2σ = 2·σ (Eq. 18)
//  3. Checks RNAV-1 and RNAV-2 flight technical error compliance
//
// Parameters:
//   - inclusionAngleDeg: the true azimuth angle α between the stations (degrees)
//   - sigma1NM, sigma2NM: the total ranging errors of the two DME stations (NM)
func AnalyzeStationPair(
	inclusionAngleDeg float64,
	sigma1NM float64,
	sigma2NM float64,
) (*types.PairAnalysisResult, error) {

	// Step 1: Compute 1σ RMS horizontal error per Eq. 17.
	rmsErrorNM, err := DMEDMEHorizontalRMSError(
		sigma1NM,
		sigma2NM,
		inclusionAngleDeg,
	)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert to 2σ (95% confidence) per Eq. 18.
	twoSigmaNM := 2 * rmsErrorNM

	return &types.PairAnalysisResult{
		InclusionAngleDeg: inclusionAngleDeg,
		RMSNM:             rmsErrorNM,
		TwoSigmaNM:        twoSigmaNM,

		// Step 3: Check flight technical error compliance.
		RNAV1: IsRNAV1FlightTechCompliant(twoSigmaNM),
		RNAV2: IsRNAV2FlightTechCompliant(twoSigmaNM),
	}, nil
}
