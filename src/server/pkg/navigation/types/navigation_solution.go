package types

// NavigationSolution holds the complete output of a DME/DME RNAV
// positioning analysis, including the estimated position, accuracy
// metrics, and RNAV specification compliance results.
//
// Reference: IDEA.md §2.2 — The solution is derived from the WLS
// algorithm (Eq. 8), with accuracy computed from the covariance
// matrix (Eq. 14), and compliance checked against RNAV-1 and
// RNAV-2 specifications (§1).
type NavigationSolution struct {
	// PositionEstimate holds the WLS-estimated position error [ΔE, ΔN].
	PositionEstimate *PositionEstimate

	// Covariance is the error covariance matrix cov(x̂) = (GᵀWG)⁻¹
	// from which accuracy metrics are extracted (IDEA.md §2.2, Eq. 14).
	Covariance *CovarianceMatrix

	// RMSAccuracyNM is the 1σ RMS horizontal error σ = √(u₁₁+u₂₂)
	// in nautical miles (IDEA.md §2.2, Eq. 12).
	RMSAccuracyNM float64

	// VerticalAccuracyNM is the 1σ vertical error σ_EL = √(u₃₃)
	// in nautical miles (IDEA.md §3.1, Eq. 21). Zero for 2D-only analysis.
	VerticalAccuracyNM float64

	// TwoSigmaNM is the 2σ (95%) horizontal positioning error,
	// which is compared against the RNAV specification thresholds.
	TwoSigmaNM float64

	// RNAV1 indicates whether the positioning accuracy meets the
	// RNAV-1 specification (2σ ≤ 1.0 NM) per IDEA.md §1.
	RNAV1 bool

	// RNAV2 indicates whether the positioning accuracy meets the
	// RNAV-2 specification (2σ ≤ 2.0 NM) per IDEA.md §1.
	RNAV2 bool
}
