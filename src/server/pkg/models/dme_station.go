package models

// DMEStation represents a Distance Measuring Equipment ground beacon station.
//
// Reference: IDEA.md §1 — "The DME navigation system consists of an airborne
// ranging interrogator and a ground ranging beacon station. The DME system
// measures the slant range between the aircraft and the ground DME beacon
// station through interrogation-response."
//
// Each station has a known position (Eᵢ, Nᵢ), an elevation, and a service
// volume that defines the maximum range at which the aircraft can receive
// its signals (IDEA.md §4, condition 1).
type DMEStation struct {
	// ID is a unique identifier for this DME station.
	ID string

	// Name is the human-readable name of this DME station.
	Name string

	// Position is the station's horizontal position in Easting/Northing
	// coordinates (nautical miles). Corresponds to (Eᵢ, Nᵢ) in the paper.
	Position Coordinate

	// ElevationFeet is the station's elevation above sea level in feet.
	// Used together with the aircraft's altitude to compute the elevation
	// angle θᵢ for the 3D geometry matrix (IDEA.md §3.1, Eq. 19).
	ElevationFeet float64

	// ServiceRadiusNM is the maximum horizontal range (in nautical miles)
	// at which the station can provide distance measurements.
	// The paper states the maximum DME ranging distance is 130 NM (§2.3).
	ServiceRadiusNM float64
}
