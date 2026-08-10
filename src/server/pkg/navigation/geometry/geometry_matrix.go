package geometry

import (
	"math"

	"gonum.org/v1/gonum/mat"

	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// BuildGeometryMatrix2D constructs the 2D direction vector matrix G
// from a set of DME measurements.
//
// Reference: IDEA.md §2.1, Eq. 6 — The geometry matrix for n DME
// stations observing the aircraft is:
//
//	G = [sinα₁ cosα₁]
//	    [sinα₂ cosα₂]
//	    [  ⋮      ⋮  ]
//	    [sinαₙ cosαₙ]
//
// where αᵢ is the true azimuth angle of the i-th DME station
// relative to the aircraft. Each row represents the direction vector
// from the aircraft to one DME station, decomposed into Easting
// (sinα) and Northing (cosα) components per Eq. 4.
func BuildGeometryMatrix2D(measurements []types.Measurement) *types.GeometryMatrix {
	rowCount := len(measurements)
	matrixData := make([]float64, 0, rowCount*2)

	for _, measurement := range measurements {
		azimuthRadians := measurement.AzimuthDeg * math.Pi / 180.0

		// Each row: [sinαᵢ, cosαᵢ] — the direction vector components.
		matrixData = append(matrixData,
			math.Sin(azimuthRadians),
			math.Cos(azimuthRadians),
		)
	}

	geometryMatrix := mat.NewDense(rowCount, 2, matrixData)

	return &types.GeometryMatrix{
		Matrix: geometryMatrix,
	}
}
