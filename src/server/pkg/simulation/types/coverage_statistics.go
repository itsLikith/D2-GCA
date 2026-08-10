package types

// CoverageStatistics holds aggregate metrics from the coverage simulation.
//
// Reference: IDEA.md §2.3 — These statistics summarize which fraction
// of the evaluated airspace meets the RNAV-1 and RNAV-2 specifications,
// corresponding to the coverage maps in Figures 3–4.
type CoverageStatistics struct {
	// TotalPoints is the number of grid points evaluated.
	TotalPoints int `json:"totalPoints"`

	// AnalyzedPoints is the number of points where analysis succeeded
	// (i.e., at least two visible DME stations with valid geometry).
	AnalyzedPoints int `json:"analyzedPoints"`

	// RNAV1Points is the count of points meeting RNAV-1 specification.
	RNAV1Points int `json:"rnav1Points"`

	// RNAV2Points is the count of points meeting RNAV-2 specification.
	RNAV2Points int `json:"rnav2Points"`

	// RNAV1CoveragePercent is the percentage of analyzed points meeting RNAV-1.
	RNAV1CoveragePercent float64 `json:"rnav1CoveragePercent"`

	// RNAV2CoveragePercent is the percentage of analyzed points meeting RNAV-2.
	RNAV2CoveragePercent float64 `json:"rnav2CoveragePercent"`

	// AverageRMSNM is the mean 1σ RMS accuracy across all analyzed points.
	AverageRMSNM float64 `json:"averageRmsNM"`

	// AverageTwoSigmaNM is the mean 2σ accuracy across all analyzed points.
	AverageTwoSigmaNM float64 `json:"averageTwoSigmaNM"`
}
