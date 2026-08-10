package stationselection

import (
	"fmt"

	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// SelectBestStationPair finds the optimal pair of DME stations for RNAV
// positioning from the set of available stations.
//
// Reference: IDEA.md §4 — "If multiple DME signals are received in the
// effective airspace that meets the three conditions, two DME stations
// that make the RNAV accuracy high can be selected according to Eq. 20
// for RNAV navigation."
//
// The selection algorithm:
//  1. Filter stations to those within range (§4, condition 1)
//  2. Enumerate all pairs and check inclusion angle constraint
//     30° ≤ α ≤ 150° (§4, condition 2)
//  3. Select the pair whose inclusion angle is closest to 90°,
//     which maximizes sin(α) and minimizes positioning error (Eq. 17)
// func SelectBestStationPair(
//
// Parameters:
//   - aircraft: the aircraft's current state (Aircraft)
//   - stations: list of candidate ground DME stations (DMEStation)
func SelectBestStationPair(
	aircraft models.Aircraft,
	stations []models.DMEStation,
) (*types.StationPairResult, error) {

	// ──────────────────────────────────────────────────
	// Condition 1: Filter to visible stations
	// ──────────────────────────────────────────────────

	visibleStations := FilterVisibleStations(aircraft, stations)

	if len(visibleStations) < 2 {
		return nil, fmt.Errorf(
			"insufficient visible stations: need ≥2, have %d",
			len(visibleStations),
		)
	}

	// ──────────────────────────────────────────────────
	// Find the best valid pair
	// ──────────────────────────────────────────────────

	var bestPair *types.StationPairResult

	for i := 0; i < len(visibleStations); i++ {
		for j := i + 1; j < len(visibleStations); j++ {

			inclusionAngleDeg := ComputeInclusionAngle(
				aircraft,
				visibleStations[i],
				visibleStations[j],
			)

			// Condition 2: Inclusion angle must be in [30°, 150°].
			if !analytical.IsValidInclusionAngle(inclusionAngleDeg) {
				continue
			}

			geometryScore := ComputeGeometryScore(inclusionAngleDeg)

			if bestPair == nil || geometryScore < bestPair.GeometryScore {
				bestPair = &types.StationPairResult{
					PrimaryStation:    visibleStations[i],
					SecondaryStation:  visibleStations[j],
					InclusionAngleDeg: inclusionAngleDeg,
					GeometryScore:     geometryScore,
				}
			}
		}
	}

	if bestPair == nil {
		return nil, fmt.Errorf("no valid station pair found meeting inclusion angle constraint")
	}

	return bestPair, nil
}
