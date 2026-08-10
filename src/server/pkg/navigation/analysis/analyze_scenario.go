package analysis

import (
	"fmt"

	"github.com/D2-GCA/PRISM/pkg/navigation/engine"
	"github.com/D2-GCA/PRISM/pkg/navigation/stationselection"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

// AnalyzeScenario performs a complete DME/DME RNAV positioning analysis
// for a given scenario, including automatic station pair selection.
//
// Reference: IDEA.md — This function implements the full paper workflow:
//  1. Select the best DME station pair (§4):
//     - Filter visible stations (condition 1)
//     - Check inclusion angle 30°–150° (condition 2)
//     - Choose pair closest to 90° for optimal accuracy
//  2. Generate measurements with azimuth, range, and sigma
//  3. Run the 2D WLS navigation pipeline (§2):
//     - Build G matrix (Eq. 6), W matrix (Eq. 13)
//     - Solve WLS (Eq. 8)
//     - Compute accuracy (Eq. 12, 14)
//     - Check RNAV compliance
func AnalyzeScenario(scenario types.Scenario) (*types.ScenarioAnalysisResult, error) {

	// ──────────────────────────────────────────────────
	// Step 1: Select Best DME Station Pair (§4)
	// ──────────────────────────────────────────────────

	selectedPair, err := stationselection.SelectBestStationPair(
		scenario.Aircraft,
		scenario.Stations,
	)
	if err != nil {
		return nil, fmt.Errorf("station selection failed: %w", err)
	}

	// ──────────────────────────────────────────────────
	// Step 2: Generate Measurements
	// ──────────────────────────────────────────────────

	measurements := stationselection.GenerateMeasurementsForPair(
		scenario.Aircraft,
		selectedPair,
		scenario.SigmaNM,
	)

	// ──────────────────────────────────────────────────
	// Step 3: Build Observation Vector
	//
	// For the accuracy analysis, the observed ranges are used
	// directly as the observation vector b.
	// ──────────────────────────────────────────────────

	observations := make([]float64, 0, len(measurements))
	for _, measurement := range measurements {
		observations = append(observations, measurement.RangeNM)
	}

	// ──────────────────────────────────────────────────
	// Step 4: Run Navigation Engine (§2)
	// ──────────────────────────────────────────────────

	navigationSolution, err := engine.Compute2DNavigationSolution(
		measurements,
		observations,
	)
	if err != nil {
		return nil, fmt.Errorf("navigation solution failed: %w", err)
	}

	return &types.ScenarioAnalysisResult{
		SelectedPair:       selectedPair,
		NavigationSolution: navigationSolution,
	}, nil
}