package models

// Aircraft represents the airborne platform whose position is being
// determined by the DME/DME RNAV system.
//
// Reference: IDEA.md §1 — "RNAV equipment operates by automatically
// determining the aircraft's position." The aircraft continuously
// measures distances to two or more DME ground stations to compute
// its horizontal position (E_A, N_A) per Eq. 1.
type Aircraft struct {
	// Position is the aircraft's ground-projected horizontal position
	// in Easting/Northing coordinates (nautical miles).
	Position Coordinate

	// AltitudeFeet is the aircraft's altitude above sea level in feet.
	// Used to compute the elevation angle to each DME station and to
	// convert slant range to horizontal distance (IDEA.md §1, last paragraph).
	AltitudeFeet float64
}
