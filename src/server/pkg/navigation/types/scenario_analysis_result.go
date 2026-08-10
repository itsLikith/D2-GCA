package types

// ScenarioAnalysisResult holds the complete output of analyzing a
// DME/DME RNAV scenario, including the selected station pair and
// the resulting navigation solution.
type ScenarioAnalysisResult struct {
	// SelectedPair contains the chosen DME station pair and its geometry.
	SelectedPair *StationPairResult

	// NavigationSolution contains the WLS positioning result, accuracy,
	// and RNAV compliance assessment.
	NavigationSolution *NavigationSolution
}
