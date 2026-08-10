package types

// ComplianceCheckResult holds the outcome of checking a positioning
// accuracy value against an RNAV specification threshold.
type ComplianceCheckResult struct {
	// SpecificationType identifies the RNAV specification (e.g., "RNAV-1").
	SpecificationType string

	// RequiredAccuracyNM is the maximum allowed 2σ positioning error (NM).
	RequiredAccuracyNM float64

	// AchievedAccuracyNM is the actual 2σ positioning error achieved (NM).
	AchievedAccuracyNM float64

	// IsCompliant is true if AchievedAccuracyNM ≤ RequiredAccuracyNM.
	IsCompliant bool
}
