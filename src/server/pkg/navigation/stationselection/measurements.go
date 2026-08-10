package stationselection

import (
	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/navigation/geometry"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// GenerateMeasurementsForPair creates DME measurements for both stations
// in a selected pair, computing the azimuth, elevation, range, and sigma
// for each.
//
// Reference: IDEA.md §2.1 — Each measurement provides the horizontal
// distance rᵢ (Eq. 2), azimuth angle αᵢ (Eq. 4), and elevation
// angle θᵢ needed to build the geometry matrix G (Eq. 6, 19).
func GenerateMeasurementsForPair(
	aircraft models.Aircraft,
	pair *types.StationPairResult,
	sigmaNM float64,
) []types.Measurement {

	return []types.Measurement{
		geometry.GenerateMeasurement(aircraft, pair.PrimaryStation, sigmaNM),
		geometry.GenerateMeasurement(aircraft, pair.SecondaryStation, sigmaNM),
	}
}
