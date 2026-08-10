package types

import "github.com/D2-GCA/PRISM/pkg/models"

// Scenario defines the inputs for a 2D DME/DME RNAV positioning analysis.
//
// Reference: IDEA.md §2 — A scenario consists of an aircraft position,
// a set of available DME ground stations, and a uniform ranging error.
// The system selects the best station pair and computes the horizontal
// positioning accuracy using the WLS algorithm.
type Scenario struct {
	// Aircraft is the airborne platform whose position is being determined.
	Aircraft models.Aircraft

	// Stations is the list of available DME ground beacon stations.
	// At least two stations must be visible for positioning (IDEA.md §2.1).
	Stations []models.DMEStation

	// SigmaNM is the uniform 1σ ranging error applied to all DME
	// measurements, in nautical miles.
	SigmaNM float64
}
