package analytical

// IsValidInclusionAngle checks whether the inclusion angle between two
// DME stations meets the minimum operational constraint.
//
// Reference: IDEA.md §4, Conclusion condition 2:
//
//	"The true azimuth inclusion angle of the two DME stations relative
//	 to the aircraft must meet the condition of 30° ≤ α ≤ 150°."
//
// This constraint comes from the simulation analysis in §2.3, which shows
// that within the 130 NM service volume, the airspace with inclusion angles
// between 30°–150° meets the RNAV-2 specification.
func IsValidInclusionAngle(inclusionAngleDeg float64) bool {
	return inclusionAngleDeg >= 30 && inclusionAngleDeg <= 150
}

// IsPreferredRNAV1Angle checks whether the inclusion angle falls within
// the preferred range for RNAV-1 compliance.
//
// Reference: IDEA.md §2.3 — "to meet RNAV-1 operation, the airspace is
// reduced to an inclusion angle of 40°–140°." This tighter range ensures
// the positioning error remains below the RNAV-1 threshold of 0.866 NM.
func IsPreferredRNAV1Angle(inclusionAngleDeg float64) bool {
	return inclusionAngleDeg >= 40 && inclusionAngleDeg <= 140
}
