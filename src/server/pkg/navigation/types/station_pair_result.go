package types

import "github.com/D2-GCA/PRISM/pkg/models"

// StationPairResult holds the selected DME station pair along with
// its geometric quality metrics.
//
// Reference: IDEA.md §4 — "If multiple DME signals are received in
// the effective airspace, two DME stations that make the RNAV accuracy
// high can be selected according to Eq. 20."
type StationPairResult struct {
	// PrimaryStation is the first DME station in the selected pair.
	PrimaryStation models.DMEStation

	// SecondaryStation is the second DME station in the selected pair.
	SecondaryStation models.DMEStation

	// InclusionAngleDeg is the true azimuth angle α between the two
	// stations as seen from the aircraft, in degrees.
	InclusionAngleDeg float64

	// GeometryScore measures how close the inclusion angle is to the
	// ideal 90° (lower is better). A score of 0 means perfect geometry.
	GeometryScore float64
}
