package geometry

import (
	"math"

	"github.com/D2-GCA/PRISM/pkg/models"
)

// TrueAzimuthDeg computes the true azimuth angle (in degrees) from
// one coordinate to another, measured clockwise from North.
//
// Reference: IDEA.md §2.1, Eq. 4 — The direction vector of the
// observation from the aircraft to the i-th DME beacon station
// is defined as:
//
//	sinαᵢ = (E - Eᵢ) / r̂ᵢ
//	cosαᵢ = (N - Nᵢ) / r̂ᵢ
//
// where αᵢ is the true azimuth angle. This function computes αᵢ
// using atan2(ΔE, ΔN), which yields the angle from North clockwise.
func TrueAzimuthDeg(from models.Coordinate, to models.Coordinate) float64 {
	deltaEasting := to.EastingNM - from.EastingNM
	deltaNorthing := to.NorthingNM - from.NorthingNM

	// atan2(ΔE, ΔN) gives the angle from North, clockwise.
	angleRadians := math.Atan2(deltaEasting, deltaNorthing)

	angleDegrees := angleRadians * 180.0 / math.Pi
	if angleDegrees < 0 {
		angleDegrees += 360
	}

	return angleDegrees
}
