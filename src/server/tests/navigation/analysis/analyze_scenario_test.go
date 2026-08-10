package analysis_test

import (
	"testing"

	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/navigation/analysis"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

func TestAnalyzeScenario_ValidPair(t *testing.T) {

	scenario := types.Scenario{
		Aircraft: models.Aircraft{
			Position: models.Coordinate{
				EastingNM:  0,
				NorthingNM: 0,
			},
		},

		SigmaNM: 1,

		Stations: []models.DMEStation{
			{
				ID:              "A",
				ServiceRadiusNM: 500,
				Position:        models.Coordinate{EastingNM: 100, NorthingNM: 0},
			},
			{
				ID:              "B",
				ServiceRadiusNM: 500,
				Position:        models.Coordinate{EastingNM: 0, NorthingNM: 100},
			},
		},
	}

	result, err := analysis.AnalyzeScenario(scenario)

	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatal("nil result")
	}

	if result.NavigationSolution == nil {
		t.Fatal("nil navigation solution")
	}

	if result.SelectedPair == nil {
		t.Fatal("nil selected pair")
	}
}

func TestAnalyzeScenario_InsufficientStations(t *testing.T) {

	scenario := types.Scenario{
		Aircraft: models.Aircraft{
			Position: models.Coordinate{EastingNM: 0, NorthingNM: 0},
		},
		SigmaNM: 1,
		Stations: []models.DMEStation{
			{
				ID:              "A",
				ServiceRadiusNM: 500,
				Position:        models.Coordinate{EastingNM: 100, NorthingNM: 0},
			},
		},
	}

	_, err := analysis.AnalyzeScenario(scenario)

	if err == nil {
		t.Fatal("expected error for insufficient stations")
	}
}
