// Package accuracy provides functions to compute the RMS horizontal
// and vertical positioning accuracy from covariance matrices produced
// by the WLS algorithm.
//
// Reference: IDEA.md §2.2, Eq. 12 (2D) and §3.1, Eq. 21 (3D).
package accuracy

import (
	"fmt"
	"math"

	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// HorizontalRMSFromCovariance2D computes the 1σ RMS horizontal error
// from a 2×2 covariance matrix.
//
// Reference: IDEA.md §2.2, Eq. 12:
//
//	σ_DME/DME = √(u₁₁ + u₂₂)
//
// where u₁₁ is the Easting error variance and u₂₂ is the Northing
// error variance, extracted from the covariance matrix
// cov(x̂) = (GᵀWG)⁻¹ (Eq. 14).
func HorizontalRMSFromCovariance2D(
	covarianceMatrix *types.CovarianceMatrix,
) (float64, error) {

	if covarianceMatrix == nil || covarianceMatrix.Matrix == nil {
		return 0, fmt.Errorf("covariance matrix is nil")
	}

	rows, cols := covarianceMatrix.Matrix.Dims()
	if rows != 2 || cols != 2 {
		return 0, fmt.Errorf(
			"expected 2×2 covariance matrix, got %d×%d",
			rows, cols,
		)
	}

	// u₁₁ = variance of Easting error (ΔE).
	varianceEasting := covarianceMatrix.Matrix.At(0, 0)
	// u₂₂ = variance of Northing error (ΔN).
	varianceNorthing := covarianceMatrix.Matrix.At(1, 1)

	return math.Sqrt(varianceEasting + varianceNorthing), nil
}

// ConvertToTwoSigma converts a 1σ RMS accuracy to a 2σ (95%)
// confidence value.
//
// The RNAV specifications require 2σ accuracy thresholds (IDEA.md §1):
//   - RNAV-2: 2σ ≤ 2.0 NM
//   - RNAV-1: 2σ ≤ 1.0 NM
func ConvertToTwoSigma(rmsSigma float64) float64 {
	return 2.0 * rmsSigma
}

// HorizontalAndVerticalRMSFromCovariance3D computes the horizontal and
// vertical RMS accuracy from a 3×3 covariance matrix.
//
// Reference: IDEA.md §3.1, Eq. 21:
//
//	σ_AZ = √(u₁₁ + u₂₂)  — horizontal RMS accuracy
//	σ_EL = √(u₃₃)         — vertical RMS accuracy
//
// where the covariance matrix is the 3×3 (GᵀWG)⁻¹ from the 3D
// WLS solution (Eq. 20).
func HorizontalAndVerticalRMSFromCovariance3D(
	covarianceMatrix *types.CovarianceMatrix,
) (*types.AccuracyResult3D, error) {

	if covarianceMatrix == nil || covarianceMatrix.Matrix == nil {
		return nil, fmt.Errorf("covariance matrix is nil")
	}

	rows, cols := covarianceMatrix.Matrix.Dims()
	if rows != 3 || cols != 3 {
		return nil, fmt.Errorf(
			"expected 3×3 covariance matrix, got %d×%d",
			rows, cols,
		)
	}

	// u₁₁ = variance of Easting error.
	varianceEasting := covarianceMatrix.Matrix.At(0, 0)
	// u₂₂ = variance of Northing error.
	varianceNorthing := covarianceMatrix.Matrix.At(1, 1)
	// u₃₃ = variance of vertical error.
	varianceVertical := covarianceMatrix.Matrix.At(2, 2)

	return &types.AccuracyResult3D{
		HorizontalRMSNM: math.Sqrt(varianceEasting + varianceNorthing),
		VerticalRMSNM:   math.Sqrt(varianceVertical),
	}, nil
}
