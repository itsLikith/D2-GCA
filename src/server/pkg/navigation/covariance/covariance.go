// Package covariance computes the error covariance matrix for the
// WLS position estimate.
//
// Reference: IDEA.md §2.2, Eq. 14:
//
//	cov(x̂) = (GᵀWG)⁻¹
//
// The covariance matrix diagonal elements represent the variance of
// the position error in each dimension. The RMS horizontal error is
// extracted as σ = √(u₁₁ + u₂₂) per Eq. 12.
package covariance

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	"github.com/D2-GCA/PRISM/pkg/navigation/types"
	"github.com/D2-GCA/PRISM/pkg/navigation/wls"
)

// ComputeFromGeometryAndWeights computes the error covariance matrix
// cov(x̂) = (GᵀWG)⁻¹ from the geometry and weight matrices.
//
// Reference: IDEA.md §2.2, Eq. 14 — When W = R⁻¹ (Eq. 13), the
// covariance simplifies from the general form (Eq. 11) to the
// compact (GᵀWG)⁻¹ form. This is the key result that enables
// efficient accuracy computation.
//
// The covariance matrix is:
//   - 2×2 for horizontal-only positioning (§2)
//   - 3×3 for 3D positioning with altitude (§3.1, Eq. 20)
func ComputeFromGeometryAndWeights(
	geometryMatrix *types.GeometryMatrix,
	weightMatrix *types.WeightMatrix,
) (*types.CovarianceMatrix, error) {

	// Build the normal matrix N = GᵀWG.
	normalMatrix := wls.BuildNormalMatrix(geometryMatrix, weightMatrix)

	// Invert N to get cov(x̂) = N⁻¹ = (GᵀWG)⁻¹.
	var covarianceMatrix mat.Dense
	if err := covarianceMatrix.Inverse(normalMatrix); err != nil {
		return nil, fmt.Errorf(
			"covariance matrix inversion failed (singular geometry): %w",
			err,
		)
	}

	return &types.CovarianceMatrix{
		Matrix: &covarianceMatrix,
	}, nil
}
