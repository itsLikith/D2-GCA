package coverage

import (
	"github.com/D2-GCA/PRISM/pkg/simulation/types"
)

// GenerateEvaluationGrid creates a uniform 2D grid of points for
// the coverage simulation.
//
// Reference: IDEA.md §2.3 — The paper evaluates positioning accuracy
// across the airspace within 130 NM of the DME stations. This function
// generates the grid points at which AnalyzeSinglePoint will be called.
func GenerateEvaluationGrid(
	minEastingNM float64,
	maxEastingNM float64,
	minNorthingNM float64,
	maxNorthingNM float64,
	stepNM float64,
) []types.GridPoint {

	var gridPoints []types.GridPoint

	for easting := minEastingNM; easting <= maxEastingNM; easting += stepNM {
		for northing := minNorthingNM; northing <= maxNorthingNM; northing += stepNM {
			gridPoints = append(gridPoints, types.GridPoint{
				EastingNM:  easting,
				NorthingNM: northing,
			})
		}
	}

	return gridPoints
}
