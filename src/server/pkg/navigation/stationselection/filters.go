package stationselection

import (
	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/navigation/geometry"
)

// FilterVisibleStations returns the subset of DME stations that are
// within range of the aircraft (i.e., the aircraft is inside each
// station's service volume).
//
// Reference: IDEA.md §4, Conclusion condition 1:
//
//	"Simultaneously receive signals from at least two standard-compliant
//	 DME stations; since the coverage range of a single DME operation
//	 signal is a circle with the horizontal projection radius of the
//	 measurement distance, the airspace that can meet the aircraft
//	 receiving signals from two navigation stations simultaneously is
//	 the intersection of two circles minus the blind area."
//
// A station is considered visible if the horizontal distance from the
// aircraft to the station is within the station's ServiceRadiusNM.
func FilterVisibleStations(
	aircraft models.Aircraft,
	stations []models.DMEStation,
) []models.DMEStation {

	visibleStations := make([]models.DMEStation, 0)

	for _, station := range stations {
		distanceNM := geometry.HorizontalDistance(
			aircraft.Position,
			station.Position,
		)

		if distanceNM <= station.ServiceRadiusNM {
			visibleStations = append(visibleStations, station)
		}
	}

	return visibleStations
}
