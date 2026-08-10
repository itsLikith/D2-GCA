// Package coverage implements spatial coverage simulation for
// DME/DME RNAV positioning accuracy analysis.
//
// It evaluates positioning accuracy across a grid of points to
// determine where RNAV-1 and RNAV-2 specifications are met,
// producing the coverage maps described in IDEA.md §2.3 (Figures 3–4).
package coverage

import (
	"github.com/D2-GCA/PRISM/pkg/models"
	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"
	"github.com/D2-GCA/PRISM/pkg/navigation/geometry"
	"github.com/D2-GCA/PRISM/pkg/navigation/stationselection"
	"github.com/D2-GCA/PRISM/pkg/simulation/types"
)

// AnalyzeSinglePoint performs the complete DME/DME RNAV accuracy
// analysis at a single grid point.
//
// Reference: IDEA.md §2.3 — At each point, the function:
//  1. Selects the best DME station pair (§4)
//  2. Computes DME ranging errors using the ICAO error model (§2.3)
//  3. Applies the closed-form accuracy equation (Eq. 17/18)
//  4. Checks RNAV-1 and RNAV-2 compliance
//
// This produces one data point for the coverage maps in Figures 3–4.
func AnalyzeSinglePoint(
	aircraft models.Aircraft,
	stations []models.DMEStation,
) (*types.CoveragePointResult, error) {

	// Step 1: Select the best DME pair meeting the three conditions (§4).
	selectedPair, err := stationselection.SelectBestStationPair(
		aircraft, stations,
	)
	if err != nil {
		return nil, err
	}

	// Step 2: Compute range to each station (Eq. 2) and apply ICAO error model.
	rangeToPrimaryNM := geometry.HorizontalDistance(
		aircraft.Position,
		selectedPair.PrimaryStation.Position,
	)
	rangeToSecondaryNM := geometry.HorizontalDistance(
		aircraft.Position,
		selectedPair.SecondaryStation.Position,
	)

	// DME error model per AC-91-FS: σ = √(σ_sis² + σ_air²) (§2.3).
	sigmaPrimaryNM := analytical.TotalDMEError(rangeToPrimaryNM)
	sigmaSecondaryNM := analytical.TotalDMEError(rangeToSecondaryNM)

	// Step 3: Closed-form accuracy analysis (Eq. 17/18).
	pairResult, err := analytical.AnalyzeStationPair(
		selectedPair.InclusionAngleDeg,
		sigmaPrimaryNM,
		sigmaSecondaryNM,
	)
	if err != nil {
		return nil, err
	}

	return &types.CoveragePointResult{
		EastingNM:         aircraft.Position.EastingNM,
		NorthingNM:        aircraft.Position.NorthingNM,
		InclusionAngleDeg: selectedPair.InclusionAngleDeg,
		RMSAccuracyNM:     pairResult.RMSNM,
		TwoSigmaNM:        pairResult.TwoSigmaNM,
		RNAV1:             pairResult.RNAV1,
		RNAV2:             pairResult.RNAV2,
	}, nil
}
