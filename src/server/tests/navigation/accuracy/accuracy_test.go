package accuracy_test

import (
	"math"
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/accuracy"
)

func TestHorizontalAccuracy3DClosedForm_FixedAltitude(t *testing.T) {
	// Fixed: σ₃=0, Eq. 24: σ = √(σ₁²/cos²θ + σ₂²/cos²θ)/sin(α)
	sigma := 0.0986
	theta := 5.0
	alpha := 90.0

	rms, err := accuracy.HorizontalAccuracy3DClosedForm(
		sigma, sigma, 0, theta, theta, alpha,
	)
	if err != nil {
		t.Fatal(err)
	}

	thetaRad := theta * math.Pi / 180
	cosTheta := math.Cos(thetaRad)

	expected := math.Sqrt(2*sigma*sigma/(cosTheta*cosTheta)) /
		math.Sin(alpha*math.Pi/180)

	if math.Abs(rms-expected) > 1e-9 {
		t.Errorf("expected %.6f, got %.6f", expected, rms)
	}
}

func TestHorizontalAccuracy3DClosedForm_InvalidAngle(t *testing.T) {
	_, err := accuracy.HorizontalAccuracy3DClosedForm(
		0.1, 0.1, 0, 0, 0, 0,
	)
	if err == nil {
		t.Fatal("expected error for α=0")
	}
}
