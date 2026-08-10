package stationselection

import "math"

// ComputeGeometryScore evaluates how close an inclusion angle is to
// the ideal 90° geometry. Lower scores indicate better geometry.
//
// Reference: IDEA.md §2.3, Eq. 17 — Since σ = √(σ₁²+σ₂²)/sin(α),
// the positioning error is minimized when sin(α) is maximized,
// which occurs at α = 90°. The score is the absolute deviation
// from this ideal:
//
//	score = |90° - α|
//
// A score of 0 means the inclusion angle is exactly 90° (optimal).
// A score of 60 means the angle is 30° or 150° (boundary of valid range).
func ComputeGeometryScore(inclusionAngleDeg float64) float64 {
	return math.Abs(90.0 - inclusionAngleDeg)
}
