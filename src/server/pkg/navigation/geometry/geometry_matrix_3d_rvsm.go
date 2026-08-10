package geometry

import (
	"math"

	"gonum.org/v1/gonum/mat"

	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// BuildGeometryMatrix3DWithAltimeter constructs the 3D direction vector
// matrix G with an additional altimeter pseudo-observation row [0, 0, 1].
//
// Reference: IDEA.md §3.2, Eq. 22 — When using DME/DME + barometric
// altimeter (RVSM or CVSM), the G matrix becomes:
//
//	G = [cosθ₁·sinα₁  cosθ₁·cosα₁  sinθ₁]
//	    [cosθ₂·sinα₂  cosθ₂·cosα₂  sinθ₂]
//	    [     0              0         1  ]
//
// The third row [0, 0, 1] represents the altimeter as a pure vertical
// observation. Its corresponding weight in the W matrix is 1/σ₃²,
// where σ₃ is the altimeter error (20m for RVSM, 50m for CVSM).
//
// For fixed altitude flight (§3.3), the altitude is considered
// accurately known (σ₃ = 0), which makes the vertical accuracy
// σ_EL = 0 and removes the altitude contribution from horizontal error.
func BuildGeometryMatrix3DWithAltimeter(
	measurements []types.Measurement,
) *types.GeometryMatrix {

	// n DME rows + 1 altimeter pseudo-observation row.
	rowCount := len(measurements) + 1
	matrixData := make([]float64, 0, rowCount*3)

	for _, measurement := range measurements {
		azimuthRadians := measurement.AzimuthDeg * math.Pi / 180.0
		elevationRadians := measurement.ElevationDeg * math.Pi / 180.0

		cosElevation := math.Cos(elevationRadians)
		sinElevation := math.Sin(elevationRadians)

		// DME row: [cosθᵢ·sinαᵢ, cosθᵢ·cosαᵢ, sinθᵢ]
		matrixData = append(matrixData,
			cosElevation*math.Sin(azimuthRadians),
			cosElevation*math.Cos(azimuthRadians),
			sinElevation,
		)
	}

	// Altimeter pseudo-observation row: [0, 0, 1]
	// This constrains the vertical dimension using altimeter data.
	matrixData = append(matrixData, 0, 0, 1)

	geometryMatrix := mat.NewDense(rowCount, 3, matrixData)

	return &types.GeometryMatrix{
		Matrix: geometryMatrix,
	}
}
