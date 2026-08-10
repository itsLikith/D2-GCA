package types

// CoverageReport holds the complete output of a coverage simulation,
// including per-point results and aggregate statistics.
type CoverageReport struct {
	// Points contains the accuracy result at each evaluated grid point.
	Points []CoveragePointResult

	// Statistics contains aggregate metrics (coverage percentages, averages).
	Statistics CoverageStatistics
}
