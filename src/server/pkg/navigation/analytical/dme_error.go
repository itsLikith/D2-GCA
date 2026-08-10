package analytical

import "math"

// SystemErrorNM is the fixed DME equipment system error σ_sis.
//
// Reference: IDEA.md §2.3 — "equipment system error σ_sis = 0.05 NM"
// per the AC-91-FS DME accuracy admission standard.
const SystemErrorNM = 0.05

// AirborneDMEError computes the airborne radio wave propagation
// environment error σ_air for a given range.
//
// Reference: IDEA.md §2.3:
//
//	σ_air = max(0.05, 0.125% × r)
//
// where r is the range to the DME station in nautical miles. This
// error grows with distance due to signal propagation effects, but
// has a minimum floor of 0.05 NM.
func AirborneDMEError(rangeNM float64) float64 {
	return math.Max(0.05, 0.00125*rangeNM)
}

// TotalDMEError computes the total single-DME ranging error σ_DME
// by combining the system error and airborne error.
//
// Reference: IDEA.md §2.3:
//
//	σ_DME = √(σ_sis² + σ_air²)
//
// where σ_sis = 0.05 NM (fixed equipment error) and σ_air is the
// range-dependent propagation error. At the maximum DME range of
// 130 NM, σ_DME ≈ 0.17 NM.
func TotalDMEError(rangeNM float64) float64 {
	airborneError := AirborneDMEError(rangeNM)

	return math.Sqrt(
		SystemErrorNM*SystemErrorNM + airborneError*airborneError,
	)
}
