package types

// Measurement represents a single DME range observation from the
// aircraft to one ground beacon station.
//
// Reference: IDEA.md §2.1 — Each DME device measures a horizontal
// distance rᵢ (Eq. 2) with a measurement error Δrᵢ (Eq. 3).
// The azimuth angle αᵢ defines the direction vector from the
// aircraft to the station (Eq. 4).
type Measurement struct {
	// StationID identifies which DME ground station produced this measurement.
	StationID string

	// AzimuthDeg is the true azimuth angle αᵢ from the aircraft to the
	// DME station, in degrees. Used to build the direction vector
	// [sinαᵢ, cosαᵢ] in the geometry matrix G (IDEA.md §2.1, Eq. 4).
	AzimuthDeg float64

	// ElevationDeg is the elevation angle θᵢ from the aircraft to the
	// DME station, in degrees. Used in the 3D geometry matrix to build
	// the direction vector [cosθᵢ·sinαᵢ, cosθᵢ·cosαᵢ, sinθᵢ]
	// (IDEA.md §3.1, Eq. 19).
	ElevationDeg float64

	// RangeNM is the horizontal distance rᵢ from the aircraft to the
	// DME station, in nautical miles. This is the horizontal projection
	// of the slant range measured by the DME interrogator (IDEA.md §1).
	RangeNM float64

	// SigmaNM is the 1σ ranging error σᵢ for this measurement, in
	// nautical miles. This determines the station's weight in the
	// weight matrix W = diag(1/σᵢ²) per IDEA.md §2.2, Eq. 13.
	SigmaNM float64
}
