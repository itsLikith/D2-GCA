package types

// AltitudeMode represents the source of altitude information
// for DME/DME RNAV 3D positioning.
//
// Reference: IDEA.md §3.1 — "According to the different sources of
// altitude information, there are currently three operating modes for
// DME/DME RNAV navigation."
type AltitudeMode int

const (
	// AltitudeModeRVSM uses a Reduced Vertical Separation Minimum
	// barometric altimeter. The 1σ error is 20 meters (~0.0108 NM).
	// Reference: IDEA.md §3.2
	AltitudeModeRVSM AltitudeMode = iota

	// AltitudeModeFixed assumes the aircraft's altitude is accurately
	// known, so the altitude error σ₃ = 0.
	// Reference: IDEA.md §3.3
	AltitudeModeFixed

	// AltitudeModeCVSM uses a Conventional Vertical Separation Minimum
	// barometric altimeter. The 1σ error is 50 meters (~0.0270 NM).
	// Reference: IDEA.md §3.4
	AltitudeModeCVSM
)

// String returns the human-readable name for this altitude mode.
func (m AltitudeMode) String() string {
	switch m {
	case AltitudeModeRVSM:
		return "RVSM"
	case AltitudeModeFixed:
		return "FIXED"
	case AltitudeModeCVSM:
		return "CVSM"
	default:
		return "UNKNOWN"
	}
}
