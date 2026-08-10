// Package wls implements the Weighted Least Squares (WLS) algorithm
// for DME/DME RNAV positioning as described in the paper.
//
// The WLS algorithm optimally estimates the aircraft's position error
// by minimizing the weighted cost function J(x̂) = (b - Gx̂)ᵀW(b - Gx̂),
// giving higher weight to DME stations with lower ranging errors.
//
// Reference: IDEA.md §2.2
package wls

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// BuildWeightMatrix constructs the diagonal accuracy weight factor
// matrix W from the ranging errors of each DME measurement.
//
// Reference: IDEA.md §2.2, Eq. 13:
//
//	W = R⁻¹ = diag(σ₁⁻², σ₂⁻², ..., σₙ⁻²)
//
// The weight matrix links the accuracy of each DME station to the
// least squares solution. Stations with smaller σᵢ receive higher
// weight, "achieving weighting that favors the selection of higher-
// accuracy DME stations" (IDEA.md §2.2).
//
// Since the measurement errors are independent and uncorrelated
// (IDEA.md §2.2, Eq. 10: R = diag(σ₁², σ₂², ..., σₙ²)),
// the weight matrix is the inverse of the error covariance matrix R.
func BuildWeightMatrix(measurements []types.Measurement) (*types.WeightMatrix, error) {
	stationCount := len(measurements)
	diagonalData := make([]float64, stationCount*stationCount)

	for i, measurement := range measurements {
		if measurement.SigmaNM <= 0 {
			return nil, fmt.Errorf(
				"invalid sigma at measurement index %d: σ must be positive",
				i,
			)
		}

		// W(i,i) = 1/σᵢ² — inverse variance weighting per Eq. 13.
		inverseSigmaSquared := 1.0 / (measurement.SigmaNM * measurement.SigmaNM)
		diagonalData[i*stationCount+i] = inverseSigmaSquared
	}

	weightMatrix := mat.NewDense(stationCount, stationCount, diagonalData)

	return &types.WeightMatrix{
		Matrix: weightMatrix,
	}, nil
}
