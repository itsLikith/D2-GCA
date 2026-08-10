package wls

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// Solve computes the WLS optimal estimate of the aircraft's position
// error using the weighted least squares algorithm.
//
// Reference: IDEA.md §2.2, Eq. 8:
//
//	x̂ = (GᵀWG)⁻¹ GᵀWb
//
// where:
//   - G is the direction vector matrix (Eq. 6 for 2D, Eq. 19 for 3D)
//   - W is the accuracy weight factor matrix W = R⁻¹ (Eq. 13)
//   - b is the observation error vector [Δr₁, Δr₂, ..., Δrₙ]ᵀ (Eq. 5)
//   - x̂ = [ΔE, ΔN]ᵀ is the estimated horizontal position error
//
// The algorithm proceeds in four steps:
//  1. Build the normal matrix N = GᵀWG
//  2. Invert N to get N⁻¹ = (GᵀWG)⁻¹
//  3. Compute the right-hand side GᵀWb
//  4. Multiply to obtain x̂ = N⁻¹ · GᵀWb
func Solve(
	geometryMatrix *types.GeometryMatrix,
	weightMatrix *types.WeightMatrix,
	observations *types.ObservationVector,
) (*types.PositionEstimate, error) {

	rowCount, _ := geometryMatrix.Matrix.Dims()
	if len(observations.Values) != rowCount {
		return nil, fmt.Errorf(
			"observation count mismatch: expected %d got %d",
			rowCount,
			len(observations.Values),
		)
	}

	// ──────────────────────────────────────────────────
	// Step 1: Build Normal Matrix  N = GᵀWG
	// ──────────────────────────────────────────────────

	normalMatrix := BuildNormalMatrix(geometryMatrix, weightMatrix)

	// ──────────────────────────────────────────────────
	// Step 2: Invert Normal Matrix  N⁻¹ = (GᵀWG)⁻¹
	// ──────────────────────────────────────────────────

	var normalMatrixInverse mat.Dense
	if err := normalMatrixInverse.Inverse(normalMatrix); err != nil {
		return nil, fmt.Errorf(
			"normal matrix inversion failed (singular geometry): %w",
			err,
		)
	}

	// ──────────────────────────────────────────────────
	// Step 3: Compute Right-Hand Side  GᵀWb
	// ──────────────────────────────────────────────────

	observationVector := mat.NewVecDense(
		len(observations.Values),
		observations.Values,
	)

	// Gᵀ
	var gTranspose mat.Dense
	gTranspose.CloneFrom(geometryMatrix.Matrix.T())

	// GᵀW
	var gTransposeTimesW mat.Dense
	gTransposeTimesW.Mul(&gTranspose, weightMatrix.Matrix)

	// GᵀWb
	var rightHandSide mat.VecDense
	rightHandSide.MulVec(&gTransposeTimesW, observationVector)

	// ──────────────────────────────────────────────────
	// Step 4: Solve  x̂ = (GᵀWG)⁻¹ · GᵀWb
	// ──────────────────────────────────────────────────

	var positionEstimate mat.VecDense
	positionEstimate.MulVec(&normalMatrixInverse, &rightHandSide)

	return &types.PositionEstimate{
		EastingNM:  positionEstimate.AtVec(0),
		NorthingNM: positionEstimate.AtVec(1),
	}, nil
}
