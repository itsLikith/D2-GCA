package analytical

// IsRNAV2FlightTechCompliant checks if the 2σ error meets the flight
// technical error threshold for RNAV-2 operation.
//
// Reference: IDEA.md §2.3 — Per ICAO Annex 10, "the flight technical
// error (2σ) allowed for RNAV-2 operation is not greater than 1.0 NM."
//
// Note: This is the flight technical error component, which is distinct
// from the total positioning error threshold (2.0 NM) checked by the
// compliance package. The total positioning error includes both the
// navigation facility system error and the flight technical error.
func IsRNAV2FlightTechCompliant(twoSigmaNM float64) bool {
	return twoSigmaNM <= 1.0
}

// IsRNAV1FlightTechCompliant checks if the 2σ error meets the flight
// technical error threshold for RNAV-1 operation.
//
// Reference: IDEA.md §2.3 — Per ICAO Annex 10, "the flight technical
// error (2σ) allowed for RNAV-1 operation is not greater than 0.5 NM."
func IsRNAV1FlightTechCompliant(twoSigmaNM float64) bool {
	return twoSigmaNM <= 0.5
}
