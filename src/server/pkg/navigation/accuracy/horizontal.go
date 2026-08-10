package accuracy

import (
	"fmt"
	"math"
)

// HorizontalAccuracy3DClosedForm computes the closed-form horizontal
// RMS accuracy for DME/DME + altimeter RNAV operation.
//
// Reference: IDEA.md §3.2, Eq. 23:
//
//	σ_AZ = √(σ₁²/cos²θ₁ + σ₂²/cos²θ₂ + σ₃²(tan²θ₁ + tan²θ₂ - 2·tanθ₁·tanθ₂·cosα)) / sinα
//
// Parameters:
//   - sigma1NM, sigma2NM: DME ranging errors σ₁, σ₂ (NM)
//   - sigma3NM: altitude sensor error σ₃ (NM); 0 for fixed altitude (§3.3)
//   - theta1Deg, theta2Deg: elevation angles θ₁, θ₂ of DME stations (degrees)
//   - inclusionAngleDeg: true azimuth inclusion angle α between the
//     two DME stations relative to the aircraft (degrees)
//
// Returns the 1σ horizontal RMS accuracy in NM.
//
// This function provides a direct analytical solution as an alternative
// to the matrix-based approach of Eq. 20. Both methods yield identical
// results for the symmetric 2-station case.
func HorizontalAccuracy3DClosedForm(
	sigma1NM float64,
	sigma2NM float64,
	sigma3NM float64,
	theta1Deg float64,
	theta2Deg float64,
	inclusionAngleDeg float64,
) (float64, error) {

	if inclusionAngleDeg <= 0 || inclusionAngleDeg >= 180 {
		return 0, fmt.Errorf(
			"invalid inclusion angle: %.2f° (must be in (0°, 180°))",
			inclusionAngleDeg,
		)
	}

	inclusionAngleRad := inclusionAngleDeg * math.Pi / 180.0
	theta1Rad := theta1Deg * math.Pi / 180.0
	theta2Rad := theta2Deg * math.Pi / 180.0

	sinInclusionAngle := math.Sin(inclusionAngleRad)
	if math.Abs(sinInclusionAngle) < 1e-12 {
		return 0, fmt.Errorf("singular geometry: sin(α) ≈ 0")
	}

	cosTheta1 := math.Cos(theta1Rad)
	cosTheta2 := math.Cos(theta2Rad)
	if math.Abs(cosTheta1) < 1e-12 || math.Abs(cosTheta2) < 1e-12 {
		return 0, fmt.Errorf("singular geometry: cos(θ) ≈ 0 (elevation ≈ 90°)")
	}

	// Term 1: σ₁²/cos²θ₁ — DME station 1 contribution.
	dmeStation1Term := (sigma1NM * sigma1NM) / (cosTheta1 * cosTheta1)

	// Term 2: σ₂²/cos²θ₂ — DME station 2 contribution.
	dmeStation2Term := (sigma2NM * sigma2NM) / (cosTheta2 * cosTheta2)

	// Term 3: σ₃²(tan²θ₁ + tan²θ₂ - 2·tanθ₁·tanθ₂·cosα) — altimeter contribution.
	tanTheta1 := math.Tan(theta1Rad)
	tanTheta2 := math.Tan(theta2Rad)
	cosInclusionAngle := math.Cos(inclusionAngleRad)

	altimeterTerm := sigma3NM * sigma3NM *
		(tanTheta1*tanTheta1 +
			tanTheta2*tanTheta2 -
			2*tanTheta1*tanTheta2*cosInclusionAngle)

	numerator := math.Sqrt(dmeStation1Term + dmeStation2Term + altimeterTerm)

	return numerator / sinInclusionAngle, nil
}
