package service

import (
	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/simulation/coverage"

	"github.com/D2-GCA/PRISM/simulation-service/pkg/dto"
)

// RunCoverageSimulation executes a spatial coverage analysis across
// a grid of points, evaluating DME/DME RNAV accuracy at each location.
//
// Reference: IDEA.md §2.3 — Produces coverage maps (Figures 3–4)
// showing where the RNAV-1 and RNAV-2 specifications are met.
func RunCoverageSimulation(
	req *dto.CoverageRequest,
) (*dto.CoverageResponse, error) {

	// Convert DTO stations to domain model.
	stations := make([]models.DMEStation, 0, len(req.Stations))
	for _, s := range req.Stations {
		stations = append(stations, models.DMEStation{
			ID: s.ID,
			Position: models.Coordinate{
				EastingNM:  s.X,
				NorthingNM: s.Y,
			},
			ElevationFeet:   s.ElevationFt,
			ServiceRadiusNM: s.ServiceVolumeNM,
		})
	}

	// Generate evaluation grid.
	grid := coverage.GenerateEvaluationGrid(
		req.MinX, req.MaxX,
		req.MinY, req.MaxY,
		req.GridStepNM,
	)

	// Run coverage simulation.
	report, err := coverage.RunCoverageSimulation(grid, stations)
	if err != nil {
		return nil, err
	}

	return &dto.CoverageResponse{
		Points:     report.Points,
		Statistics: report.Statistics,
	}, nil
}
