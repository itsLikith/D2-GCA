package coverage_test

import (
	"testing"

	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/simulation/coverage"
)

func TestSimulateCoverage_ValidGrid(t *testing.T) {

	stations := []models.DMEStation{
		{
			ID:              "A",
			ServiceRadiusNM: 130,
			Position:        models.Coordinate{EastingNM: -50, NorthingNM: 0},
		},
		{
			ID:              "B",
			ServiceRadiusNM: 130,
			Position:        models.Coordinate{EastingNM: 50, NorthingNM: 0},
		},
	}

	grid := coverage.GenerateEvaluationGrid(-100, 100, -100, 100, 50)

	report, err := coverage.RunCoverageSimulation(grid, stations)

	if err != nil {
		t.Fatal(err)
	}

	if report == nil {
		t.Fatal("nil report")
	}

	if len(report.Points) == 0 {
		t.Fatal("expected analyzed points")
	}

	// Should have statistics
	if report.Statistics.TotalPoints == 0 {
		t.Error("total points should be > 0")
	}
}

func TestGenerateGrid_Dimensions(t *testing.T) {

	grid := coverage.GenerateEvaluationGrid(0, 10, 0, 10, 5)

	// (0,5,10) × (0,5,10) = 3×3 = 9 points
	if len(grid) != 9 {
		t.Errorf("expected 9 points, got %d", len(grid))
	}
}
