package analytical_test

import (
	"math"
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"
)

func TestTotalDMEError_MinimumRange(t *testing.T) {

	// At short range, σ_air = max(0.05, 0.00125*r)
	// For r = 10 NM: 0.00125*10 = 0.0125 < 0.05 → σ_air = 0.05
	// σ = √(0.05² + 0.05²) = √0.005 ≈ 0.0707

	sigma := analytical.TotalDMEError(10.0)
	expected := math.Sqrt(0.05*0.05 + 0.05*0.05)

	if math.Abs(sigma-expected) > 1e-6 {
		t.Errorf("expected %.6f, got %.6f", expected, sigma)
	}
}

func TestTotalDMEError_MaxRange(t *testing.T) {

	// At 130 NM: σ_air = max(0.05, 0.00125*130) = 0.1625
	// σ = √(0.05² + 0.1625²) ≈ 0.17

	sigma := analytical.TotalDMEError(130.0)
	expected := math.Sqrt(0.05*0.05 + 0.1625*0.1625)

	if math.Abs(sigma-expected) > 0.001 {
		t.Errorf("expected ≈%.3f, got %.3f", expected, sigma)
	}
}

func TestAirborneDMEError_Threshold(t *testing.T) {

	// Breakpoint at r = 0.05/0.00125 = 40 NM
	below := analytical.AirborneDMEError(30.0) // 0.00125*30=0.0375 < 0.05
	above := analytical.AirborneDMEError(60.0) // 0.00125*60=0.075 > 0.05

	if below != 0.05 {
		t.Errorf("below threshold: expected 0.05, got %f", below)
	}

	expected := 0.00125 * 60.0

	if math.Abs(above-expected) > 1e-9 {
		t.Errorf("above threshold: expected %f, got %f", expected, above)
	}
}
