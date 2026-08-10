package types

// AccuracyResult3D holds the decomposed horizontal and vertical
// accuracy from a 3×3 covariance matrix.
//
// Reference: IDEA.md §3.1, Eq. 21:
//
//	σ_AZ = √(u₁₁ + u₂₂)  — horizontal RMS
//	σ_EL = √(u₃₃)         — vertical RMS
type AccuracyResult3D struct {
	// HorizontalRMSNM is the 1σ horizontal RMS accuracy (NM).
	HorizontalRMSNM float64

	// VerticalRMSNM is the 1σ vertical RMS accuracy (NM).
	VerticalRMSNM float64
}
