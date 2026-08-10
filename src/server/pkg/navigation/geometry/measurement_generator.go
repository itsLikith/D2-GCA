package geometry

import (
	"math"

	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// MetersPerNauticalMile is the conversion factor from meters to NM.
const MetersPerNauticalMile = 1852.0

// FeetToMeters is the conversion factor from feet to meters.
const FeetToMeters = 0.3048

// GenerateMeasurement creates a Measurement from an aircraft position
// and a DME ground station, computing the azimuth angle, elevation angle,
// and horizontal distance.
//
// Reference: IDEA.md §1 — "The beacon elevation must be obtained from the
// flight data storage component, while the aircraft's altitude can be
// obtained from three sources." The slant range must be projected to
// horizontal distance by solving the slant range triangle.
//
// The elevation angle θᵢ is computed as:
//
//	θᵢ = atan2(Δh, rᵢ)
//
// where Δh is the altitude difference between the aircraft and the station,
// and rᵢ is the horizontal distance (Eq. 2). This angle is needed for the
// 3D geometry matrix (§3.1, Eq. 19).
func GenerateMeasurement(
	aircraft models.Aircraft,
	station models.DMEStation,
	sigmaNM float64,
) types.Measurement {

	// Compute horizontal distance rᵢ per Eq. 2.
	horizontalDistanceNM := HorizontalDistance(aircraft.Position, station.Position)

	// Compute altitude difference for elevation angle calculation.
	// The paper requires converting slant range to horizontal distance;
	// here we compute the elevation angle from altitude difference and
	// horizontal range for the 3D geometry matrix (§3.1).
	altitudeDiffFeet := aircraft.AltitudeFeet - station.ElevationFeet
	altitudeDiffMeters := altitudeDiffFeet * FeetToMeters
	altitudeDiffNM := altitudeDiffMeters / MetersPerNauticalMile

	// Compute elevation angle θᵢ in degrees.
	var elevationDeg float64
	if horizontalDistanceNM > 1e-9 {
		elevationRadians := math.Atan2(altitudeDiffNM, horizontalDistanceNM)
		elevationDeg = elevationRadians * 180.0 / math.Pi
	} else if altitudeDiffNM > 0 {
		elevationDeg = 90.0
	} else if altitudeDiffNM < 0 {
		elevationDeg = -90.0
	}

	return types.Measurement{
		StationID: station.ID,

		// Azimuth angle αᵢ from the aircraft to the station (Eq. 4).
		AzimuthDeg: TrueAzimuthDeg(aircraft.Position, station.Position),

		// Elevation angle θᵢ for 3D geometry (Eq. 19).
		ElevationDeg: elevationDeg,

		// Horizontal distance rᵢ (Eq. 2).
		RangeNM: horizontalDistanceNM,

		// Ranging error σᵢ for weight matrix (Eq. 13).
		SigmaNM: sigmaNM,
	}
}
