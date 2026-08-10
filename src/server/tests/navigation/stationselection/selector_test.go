package stationselection_test

import (
	"math"
	"testing"

	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/navigation/stationselection"
)

func TestSelectBestPair_OrthogonalStations(t *testing.T) {

	aircraft := models.Aircraft{
		Position: models.Coordinate{EastingNM: 0, NorthingNM: 0},
	}

	stations := []models.DMEStation{
		{
			ID:              "A",
			ServiceRadiusNM: 200,
			Position:        models.Coordinate{EastingNM: 10, NorthingNM: 0},
		},
		{
			ID:              "B",
			ServiceRadiusNM: 200,
			Position:        models.Coordinate{EastingNM: 0, NorthingNM: 10},
		},
	}

	result, err := stationselection.SelectBestStationPair(
		aircraft,
		stations,
	)

	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatal("nil result")
	}

	// Inclusion angle should be 90°
	if math.Abs(result.InclusionAngleDeg-90.0) > 0.1 {
		t.Errorf("expected 90°, got %.1f°",
			result.InclusionAngleDeg)
	}
}

func TestSelectBestPair_NoVisibleStations(t *testing.T) {

	aircraft := models.Aircraft{
		Position: models.Coordinate{EastingNM: 0, NorthingNM: 0},
	}

	stations := []models.DMEStation{
		{
			ID:              "A",
			ServiceRadiusNM: 1, // Too small
			Position:        models.Coordinate{EastingNM: 100, NorthingNM: 0},
		},
		{
			ID:              "B",
			ServiceRadiusNM: 1,
			Position:        models.Coordinate{EastingNM: 0, NorthingNM: 100},
		},
	}

	_, err := stationselection.SelectBestStationPair(
		aircraft,
		stations,
	)

	if err == nil {
		t.Fatal("expected error for no visible stations")
	}
}

func TestSelectBestPair_CollinearStationsRejected(t *testing.T) {

	aircraft := models.Aircraft{
		Position: models.Coordinate{EastingNM: 0, NorthingNM: 0},
	}

	// Both stations roughly in same direction → angle ≈ 0°
	stations := []models.DMEStation{
		{
			ID:              "A",
			ServiceRadiusNM: 200,
			Position:        models.Coordinate{EastingNM: 10, NorthingNM: 0},
		},
		{
			ID:              "B",
			ServiceRadiusNM: 200,
			Position:        models.Coordinate{EastingNM: 20, NorthingNM: 0},
		},
	}

	_, err := stationselection.SelectBestStationPair(
		aircraft,
		stations,
	)

	if err == nil {
		t.Fatal("expected error for collinear stations (angle < 30°)")
	}
}
