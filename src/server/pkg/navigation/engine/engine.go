// Package engine orchestrates the complete DME/DME RNAV positioning
// computation pipeline, combining geometry, WLS, covariance, accuracy,
// and compliance into a single NavigationSolution.
package engine

import (
	"github.com/D2-GCA/PRISM/pkg/navigation/accuracy"
	"github.com/D2-GCA/PRISM/pkg/navigation/compliance"
	"github.com/D2-GCA/PRISM/pkg/navigation/covariance"
	"github.com/D2-GCA/PRISM/pkg/navigation/geometry"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
	"github.com/D2-GCA/PRISM/pkg/navigation/wls"
)

// Compute2DNavigationSolution performs a complete 2D DME/DME RNAV
// positioning analysis per IDEA.md §2.
//
// The pipeline follows the paper's mathematical framework:
//  1. Build geometry matrix G from azimuth angles (§2.1, Eq. 6)
//  2. Build weight matrix W = diag(1/σᵢ²) from ranging errors (§2.2, Eq. 13)
//  3. Solve WLS: x̂ = (GᵀWG)⁻¹GᵀWb for position estimate (§2.2, Eq. 8)
//  4. Compute covariance: cov(x̂) = (GᵀWG)⁻¹ (§2.2, Eq. 14)
//  5. Extract RMS accuracy: σ = √(u₁₁+u₂₂) (§2.2, Eq. 12)
//  6. Check RNAV-1 and RNAV-2 compliance (§1)
func Compute2DNavigationSolution(
	measurements []types.Measurement,
	observations []float64,
) (*types.NavigationSolution, error) {

	// Step 1: Build the 2D geometry matrix G (Eq. 6).
	geometryMatrix := geometry.BuildGeometryMatrix2D(measurements)

	// Step 2: Build the weight matrix W = R⁻¹ (Eq. 13).
	weightMatrix, err := wls.BuildWeightMatrix(measurements)
	if err != nil {
		return nil, err
	}

	// Step 3: Solve WLS for position estimate x̂ (Eq. 8).
	positionEstimate, err := wls.Solve(
		geometryMatrix,
		weightMatrix,
		&types.ObservationVector{Values: observations},
	)
	if err != nil {
		return nil, err
	}

	// Step 4: Compute covariance cov(x̂) = (GᵀWG)⁻¹ (Eq. 14).
	covarianceMatrix, err := covariance.ComputeFromGeometryAndWeights(
		geometryMatrix, weightMatrix,
	)
	if err != nil {
		return nil, err
	}

	// Step 5: Extract RMS accuracy σ = √(u₁₁+u₂₂) (Eq. 12).
	rmsAccuracyNM, err := accuracy.HorizontalRMSFromCovariance2D(covarianceMatrix)
	if err != nil {
		return nil, err
	}

	twoSigmaNM := accuracy.ConvertToTwoSigma(rmsAccuracyNM)

	// Step 6: Check RNAV specification compliance (§1).
	rnav1Result := compliance.CheckRNAV1Compliance(twoSigmaNM)
	rnav2Result := compliance.CheckRNAV2Compliance(twoSigmaNM)

	return &types.NavigationSolution{
		PositionEstimate:   positionEstimate,
		Covariance:         covarianceMatrix,
		RMSAccuracyNM:      rmsAccuracyNM,
		TwoSigmaNM:         twoSigmaNM,
		RNAV1:              rnav1Result.IsCompliant,
		RNAV2:              rnav2Result.IsCompliant,
	}, nil
}
