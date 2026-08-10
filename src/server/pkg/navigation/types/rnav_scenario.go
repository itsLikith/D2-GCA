package types

import "github.com/D2-GCA/PRISM/pkg/models"

// RNAVScenario extends the base Scenario with 3D parameters required
// for the three RNAV operation accuracy models described in the paper.
//
// Reference: IDEA.md §3.1 — "To continuously obtain the three-dimensional
// positioning information of the aircraft's RNAV navigation, the airborne
// RNAV system must also receive altitude information from other sensors."
//
// The three operating modes are:
//   - DME/DME + RVSM barometric altimeter (§3.2, σ₃ = 20m)
//   - DME/DME + fixed altitude flight    (§3.3, σ₃ = 0)
//   - DME/DME + CVSM barometric altimeter (§3.4, σ₃ = 50m)
type RNAVScenario struct {
	// Aircraft is the airborne platform.
	Aircraft models.Aircraft

	// Stations is the list of available DME ground beacon stations.
	Stations []models.DMEStation

	// AltitudeMode specifies the source of altitude information,
	// which determines the altitude sensor error σ₃.
	AltitudeMode AltitudeMode

	// AltitudeSigmaNM is the 1σ error of the altitude sensor in NM.
	// For RVSM: ~0.0108 NM (20m), Fixed: 0, CVSM: ~0.0270 NM (50m).
	AltitudeSigmaNM float64
}
