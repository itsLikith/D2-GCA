package types

// PairAnalysisResult holds the positioning accuracy analysis output
// for a specific DME/DME station pair configuration.
//
// This is the result of applying the closed-form Eq. 17 to a pair
// of DME stations with known ranging errors and inclusion angle.
type PairAnalysisResult struct {
	// InclusionAngleDeg is the true azimuth angle α between the two
	// DME stations relative to the aircraft, in degrees.
	InclusionAngleDeg float64

	// RMSNM is the 1σ RMS horizontal positioning error (NM) per Eq. 17.
	RMSNM float64

	// TwoSigmaNM is the 2σ (95%) horizontal positioning error (NM),
	// used for RNAV compliance checking.
	TwoSigmaNM float64

	// RNAV1 indicates if the flight technical error meets RNAV-1 threshold.
	RNAV1 bool

	// RNAV2 indicates if the flight technical error meets RNAV-2 threshold.
	RNAV2 bool
}
