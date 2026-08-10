package analytical_test

import (
	"math"
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"
)

func TestDMEDMERMSError_Equation17(t *testing.T) {

	// Paper Eq. 17: σ = √(σ₁²+σ₂²)/sin(α)
	// With σ₁=σ₂=0.17, α=90°:
	// σ = √(0.0289+0.0289)/1 = √0.0578 ≈ 0.2404

	rms, err := analytical.DMEDMEHorizontalRMSError(0.17, 0.17, 90.0)

	if err != nil {
		t.Fatal(err)
	}

	expected := math.Sqrt(0.17*0.17+0.17*0.17) / math.Sin(90.0*math.Pi/180.0)

	if math.Abs(rms-expected) > 1e-6 {
		t.Errorf("expected %.6f, got %.6f", expected, rms)
	}
}

func TestDMEDMERMSError_Equation17_30Degrees(t *testing.T) {

	// At α=30°: sin(30°) = 0.5
	// σ = √(0.17²+0.17²)/0.5

	rms, err := analytical.DMEDMEHorizontalRMSError(0.17, 0.17, 30.0)

	if err != nil {
		t.Fatal(err)
	}

	expected := math.Sqrt(2*0.17*0.17) / math.Sin(30.0*math.Pi/180.0)

	if math.Abs(rms-expected) > 1e-6 {
		t.Errorf("expected %.6f, got %.6f", expected, rms)
	}

	// 2σ at 30° with max range error
	twoSigma := 2 * rms

	// Paper states 30°-150° meets RNAV-2 (≤1.73 NM)
	if twoSigma > 1.73 {
		t.Logf("2σ=%.3f exceeds RNAV-2 max (1.73 NM) at 30°", twoSigma)
	}
}

func TestDMEDMERMSError_InvalidAngle(t *testing.T) {

	_, err := analytical.DMEDMEHorizontalRMSError(0.17, 0.17, 0.0)

	if err == nil {
		t.Fatal("expected error for angle=0")
	}

	_, err = analytical.DMEDMEHorizontalRMSError(0.17, 0.17, 180.0)

	if err == nil {
		t.Fatal("expected error for angle=180")
	}
}

func TestDMEDMERMSError_SymmetryProperty(t *testing.T) {

	// RMS should be symmetric: swap σ₁,σ₂
	rms1, _ := analytical.DMEDMEHorizontalRMSError(0.1, 0.2, 60.0)
	rms2, _ := analytical.DMEDMEHorizontalRMSError(0.2, 0.1, 60.0)

	if math.Abs(rms1-rms2) > 1e-12 {
		t.Errorf("symmetry violated: %f != %f", rms1, rms2)
	}
}
