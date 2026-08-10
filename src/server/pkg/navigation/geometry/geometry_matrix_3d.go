package geometry

import (
	"math"

	"gonum.org/v1/gonum/mat"

	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// BuildGeometryMatrix3D constructs the 3D direction vector matrix G
// from DME measurements that include both azimuth and elevation angles.
//
// Reference: IDEA.md §3.1, Eq. 19 — The 3D direction vector matrix is:
//
//	G = [cosθ₁·sinα₁  cosθ₁·cosα₁  sinθ₁]
//	    [cosθ₂·sinα₂  cosθ₂·cosα₂  sinθ₂]
//	    [     ⋮             ⋮         ⋮   ]
//	    [cosθₙ·sinαₙ  cosθₙ·cosαₙ  sinθₙ]
//
// where αᵢ is the true azimuth angle and θᵢ is the elevation angle
// of the i-th DME station observing the aircraft.
//
// This extends the 2D matrix (Eq. 6) by incorporating the elevation
// angle to enable 3D positioning analysis for the operation accuracy
// model described in §3.1.
func BuildGeometryMatrix3D(measurements []types.Measurement) *types.GeometryMatrix {
	rowCount := len(measurements)
	matrixData := make([]float64, 0, rowCount*3)

	for _, measurement := range measurements {
		azimuthRadians := measurement.AzimuthDeg * math.Pi / 180.0
		elevationRadians := measurement.ElevationDeg * math.Pi / 180.0

		cosElevation := math.Cos(elevationRadians)
		sinElevation := math.Sin(elevationRadians)

		// Each row: [cosθᵢ·sinαᵢ, cosθᵢ·cosαᵢ, sinθᵢ]
		matrixData = append(matrixData,
			cosElevation*math.Sin(azimuthRadians),
			cosElevation*math.Cos(azimuthRadians),
			sinElevation,
		)
	}

	geometryMatrix := mat.NewDense(rowCount, 3, matrixData)

	return &types.GeometryMatrix{
		Matrix: geometryMatrix,
	}
}
