package types

import "github.com/D2-GCA/PRISM/pkg/models"

// CoverageSimulationScenario defines the inputs for a spatial
// coverage simulation.
//
// Reference: IDEA.md §2.3 — The coverage simulation evaluates
// positioning accuracy across a grid of points to produce coverage
// maps (Figures 3–4).
type CoverageSimulationScenario struct {
	// Stations is the list of DME ground stations to evaluate.
	Stations []models.DMEStation

	// MinEastingNM is the western boundary of the grid (NM).
	MinEastingNM float64

	// MaxEastingNM is the eastern boundary of the grid (NM).
	MaxEastingNM float64

	// MinNorthingNM is the southern boundary of the grid (NM).
	MinNorthingNM float64

	// MaxNorthingNM is the northern boundary of the grid (NM).
	MaxNorthingNM float64

	// GridStepNM is the spacing between grid points (NM).
	GridStepNM float64
}
