package types

// CoveragePointResult holds the positioning accuracy analysis at a
// single grid point in the coverage simulation.
type CoveragePointResult struct {
	// EastingNM is the E coordinate of the grid point (NM).
	EastingNM float64 `json:"eastingNM"`

	// NorthingNM is the N coordinate of the grid point (NM).
	NorthingNM float64 `json:"northingNM"`

	// InclusionAngleDeg is the inclusion angle α of the selected
	// DME pair at this point (degrees).
	InclusionAngleDeg float64 `json:"inclusionAngleDeg"`

	// RMSAccuracyNM is the 1σ RMS horizontal error at this point (NM).
	RMSAccuracyNM float64 `json:"rmsAccuracyNM"`

	// TwoSigmaNM is the 2σ (95%) horizontal error at this point (NM).
	TwoSigmaNM float64 `json:"twoSigmaNM"`

	// RNAV1 indicates RNAV-1 compliance at this point.
	RNAV1 bool `json:"rnav1"`

	// RNAV2 indicates RNAV-2 compliance at this point.
	RNAV2 bool `json:"rnav2"`
}
