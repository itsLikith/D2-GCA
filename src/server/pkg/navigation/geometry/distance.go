package geometry

import (
	"math"

	"github.com/D2-GCA/PRISM/pkg/models"
)

// HorizontalDistance computes the horizontal distance between two points
// in the Easting/Northing coordinate plane.
//
// Reference: IDEA.md §2.1, Eq. 2:
//
//	rᵢ = √((E - Eᵢ)² + (N - Nᵢ)²)
//
// This is the horizontal projection of the slant range measured by the
// DME interrogator. The paper states: "to achieve positioning of the
// aircraft at the point to be measured, at least two simultaneous
// distance measurements are required."
func HorizontalDistance(pointA models.Coordinate, pointB models.Coordinate) float64 {
	deltaEasting := pointB.EastingNM - pointA.EastingNM
	deltaNorthing := pointB.NorthingNM - pointA.NorthingNM

	return math.Sqrt(deltaEasting*deltaEasting + deltaNorthing*deltaNorthing)
}
