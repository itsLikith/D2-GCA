package wls

import (
	"gonum.org/v1/gonum/mat"

	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// BuildNormalMatrix computes the normal equation matrix N = GᵀWG.
//
// Reference: IDEA.md §2.2 — The normal matrix is the core of the WLS
// solution. Its inverse (GᵀWG)⁻¹ directly yields the error covariance
// matrix (Eq. 14), and it appears in the optimal estimate formula
// x̂ = (GᵀWG)⁻¹GᵀWb (Eq. 8).
//
// The normal matrix is symmetric positive-definite when the geometry
// is non-degenerate (i.e., the inclusion angle α ≠ 0° or 180°).
func BuildNormalMatrix(
	geometryMatrix *types.GeometryMatrix,
	weightMatrix *types.WeightMatrix,
) *mat.Dense {

	// Gᵀ — transpose of the geometry matrix.
	var gTranspose mat.Dense
	gTranspose.CloneFrom(geometryMatrix.Matrix.T())

	// GᵀW — geometry transpose multiplied by weight matrix.
	var gTransposeTimesW mat.Dense
	gTransposeTimesW.Mul(&gTranspose, weightMatrix.Matrix)

	// GᵀWG — the normal matrix (n×n where n is the dimension of x).
	var normalMatrix mat.Dense
	normalMatrix.Mul(&gTransposeTimesW, geometryMatrix.Matrix)

	return &normalMatrix
}
