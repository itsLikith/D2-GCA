package analytical

// Paper Section 3 altitude sensor error standards.
//
// These constants define the 1σ error for each altitude
// information source used in DME/DME RNAV 3D positioning.

const (
	// MetersPerNM is the conversion factor from meters to NM.
	MetersPerNM = 1852.0

	// RVSMSigmaMeters is the 1σ error of the RVSM barometric
	// altimeter: 20 meters (Paper Section 3.2).
	RVSMSigmaMeters = 20.0

	// RVSMSigmaNM is the RVSM error in nautical miles.
	RVSMSigmaNM = RVSMSigmaMeters / MetersPerNM

	// CVSMSigmaMeters is the 1σ error of the CVSM barometric
	// altimeter: 50 meters (Paper Section 3.4).
	CVSMSigmaMeters = 50.0

	// CVSMSigmaNM is the CVSM error in nautical miles.
	CVSMSigmaNM = CVSMSigmaMeters / MetersPerNM

	// FixedAltitudeSigmaNM is the altitude error for fixed
	// altitude flight: 0 (Paper Section 3.3).
	FixedAltitudeSigmaNM = 0.0

	// ICAODMESigmaNM is the ICAO standard DME accuracy used
	// in the paper's simulation: 0.0986 NM.
	ICAODMESigmaNM = 0.0986

	// MaxDMERangeNM is the maximum DME ranging distance: 130 NM.
	MaxDMERangeNM = 130.0

	// MaxDMESigmaNM is the DME ranging error at max range.
	// σ = √(σ_sis² + σ_air²) where σ_sis=0.05, σ_air=max(0.05, 0.00125*130)=0.1625
	// → √(0.0025 + 0.0264) ≈ 0.17 NM
	MaxDMESigmaNM = 0.17
)
