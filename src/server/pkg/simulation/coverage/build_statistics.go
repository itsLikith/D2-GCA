package coverage

import (
	"github.com/D2-GCA/PRISM/pkg/simulation/types"
)

// ComputeCoverageStatistics aggregates the per-point results into
// summary statistics for the coverage simulation.
func ComputeCoverageStatistics(pointResults []types.CoveragePointResult) types.CoverageStatistics {
	stats := types.CoverageStatistics{
		TotalPoints: len(pointResults),
	}

	var sumRMSNM float64
	var sumTwoSigmaNM float64

	for _, result := range pointResults {
		stats.AnalyzedPoints++

		sumRMSNM += result.RMSAccuracyNM
		sumTwoSigmaNM += result.TwoSigmaNM

		if result.RNAV1 {
			stats.RNAV1Points++
		}
		if result.RNAV2 {
			stats.RNAV2Points++
		}
	}

	if stats.AnalyzedPoints > 0 {
		analyzedFloat := float64(stats.AnalyzedPoints)

		stats.RNAV1CoveragePercent = float64(stats.RNAV1Points) / analyzedFloat * 100
		stats.RNAV2CoveragePercent = float64(stats.RNAV2Points) / analyzedFloat * 100
		stats.AverageRMSNM = sumRMSNM / analyzedFloat
		stats.AverageTwoSigmaNM = sumTwoSigmaNM / analyzedFloat
	}

	return stats
}
