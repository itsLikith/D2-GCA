package compliance

import "github.com/D2-GCA/PRISM/pkg/navigation/types"

// CheckRNAV1Compliance checks if the achieved 2σ accuracy meets the
// RNAV-1 specification.
//
// Reference: IDEA.md §1 — "RNAV-1 specification is planned to be
// mandatory in terminal areas. The RNAV-1 specification requires
// that the 2σ positioning error of the system does not exceed 1.0 NM
// during the entire flight phase."
//
// The RNAV-1 threshold accounts for both navigation facility system
// error and flight technical error. Per ICAO Annex 10, the flight
// technical error (2σ) allowed for RNAV-1 operation is ≤ 0.5 NM,
// resulting in a maximum navigation positioning error of 0.866 NM (§2.3).
func CheckRNAV1Compliance(twoSigmaAccuracyNM float64) *types.ComplianceCheckResult {
	return &types.ComplianceCheckResult{
		SpecificationType:  "RNAV-1",
		RequiredAccuracyNM: 1.0,
		AchievedAccuracyNM: twoSigmaAccuracyNM,
		IsCompliant:        twoSigmaAccuracyNM <= 1.0,
	}
}
