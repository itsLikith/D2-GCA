package coverage

import (
	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/simulation/types"
)

// RunCoverageSimulation evaluates DME/DME RNAV positioning accuracy
// across all points in the grid, producing a coverage report.
//
// Reference: IDEA.md §2.3 — This simulation produces the data for
// the coverage maps in Figures 3–4, showing where the DME/DME RNAV
// operation meets the RNAV-2 and RNAV-1 specifications.
//
// Points where analysis fails (e.g., no visible stations, no valid
// geometry) are silently skipped, as they represent airspace outside
// the DME coverage area.
func RunCoverageSimulation(
	grid []types.GridPoint,
	stations []models.DMEStation,
) (*types.CoverageReport, error) {

	pointResults := make([]types.CoveragePointResult, 0, len(grid))

	for _, gridPoint := range grid {
		aircraft := models.Aircraft{
			Position: models.Coordinate{
				EastingNM:  gridPoint.EastingNM,
				NorthingNM: gridPoint.NorthingNM,
			},
		}

		result, err := AnalyzeSinglePoint(aircraft, stations)
		if err != nil {
			// Skip points outside DME coverage or with invalid geometry.
			continue
		}

		pointResults = append(pointResults, *result)
	}

	return &types.CoverageReport{
		Points:     pointResults,
		Statistics: ComputeCoverageStatistics(pointResults),
	}, nil
}
