package types

// AccuracyResult holds the RMS and 2σ horizontal positioning accuracy
// from a DME/DME RNAV analysis.
//
// Reference: IDEA.md §2.2, Eq. 12 — σ = √(u₁₁+u₂₂), and the 2σ
// value is used for RNAV specification compliance checking.
type AccuracyResult struct {
	// RMSAccuracyNM is the 1σ RMS horizontal error in nautical miles.
	RMSAccuracyNM float64

	// TwoSigmaNM is the 2σ (95%) horizontal error in nautical miles.
	TwoSigmaNM float64
}
