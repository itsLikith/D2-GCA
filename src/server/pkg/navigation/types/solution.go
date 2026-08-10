package types

// PositionEstimate holds the WLS-estimated aircraft position error
// components in the horizontal plane.
//
// Reference: IDEA.md §2.2, Eq. 8 — The optimal estimate x̂ = (GᵀWG)⁻¹GᵀWb
// yields position error components [ΔE, ΔN]ᵀ.
type PositionEstimate struct {
	// EastingNM is the estimated Easting component (ΔE) in nautical miles.
	EastingNM float64

	// NorthingNM is the estimated Northing component (ΔN) in nautical miles.
	NorthingNM float64
}
