package compliance

import "github.com/D2-GCA/PRISM/pkg/navigation/types"

// CheckRNAV2Compliance checks if the achieved 2σ accuracy meets the
// RNAV-2 specification.
//
// Reference: IDEA.md §1 — "RNAV-2 specification is planned for use
// on en-route segments. The RNAV-2 specification requires that the
// 2σ positioning error of the area navigation system does not exceed
// 2.0 NM during the entire flight phase."
//
// Per ICAO Annex 10, the flight technical error (2σ) allowed for
// RNAV-2 operation is ≤ 1.0 NM, resulting in a maximum navigation
// positioning error of 1.73 NM (95%) per §2.3.
func CheckRNAV2Compliance(twoSigmaAccuracyNM float64) *types.ComplianceCheckResult {
	return &types.ComplianceCheckResult{
		SpecificationType:  "RNAV-2",
		RequiredAccuracyNM: 2.0,
		AchievedAccuracyNM: twoSigmaAccuracyNM,
		IsCompliant:        twoSigmaAccuracyNM <= 2.0,
	}
}
