package analytical

import (
	"fmt"
	"math"
)

// DMEDMEHorizontalRMSError computes the closed-form 1σ RMS horizontal
// error for a DME/DME pair configuration in 2D.
//
// Reference: IDEA.md §2.3, Eq. 17:
//
//	σ_DME/DME = √(σ₁² + σ₂²) / sin(α)
//
// where σ₁ and σ₂ are the ranging errors of the two DME stations,
// and α is the true azimuth inclusion angle between them.
//
// This equation is derived from the closed-form expansion of the
// 2×2 covariance matrix (GᵀWG)⁻¹ in Eq. 16. The sin(α) term
// in the denominator shows that accuracy degrades as the inclusion
// angle approaches 0° or 180° (collinear geometry).
func DMEDMEHorizontalRMSError(
	sigma1NM float64,
	sigma2NM float64,
	inclusionAngleDeg float64,
) (float64, error) {

	if inclusionAngleDeg <= 0 || inclusionAngleDeg >= 180 {
		return 0, fmt.Errorf(
			"invalid inclusion angle: %.2f° (must be in (0°, 180°))",
			inclusionAngleDeg,
		)
	}

	inclusionAngleRad := inclusionAngleDeg * math.Pi / 180.0
	sinInclusionAngle := math.Sin(inclusionAngleRad)

	if math.Abs(sinInclusionAngle) < 1e-9 {
		return 0, fmt.Errorf("singular geometry: sin(α) ≈ 0")
	}

	return math.Sqrt(sigma1NM*sigma1NM+sigma2NM*sigma2NM) / sinInclusionAngle, nil
}
