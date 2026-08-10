package engine_test

import (
	"math"
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/engine"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

func TestCompute2DNavigationSolution(t *testing.T) {
	measurements := []types.Measurement{
		{AzimuthDeg: 0, SigmaNM: 1},
		{AzimuthDeg: 90, SigmaNM: 1},
	}

	result, err := engine.Compute2DNavigationSolution(
		measurements,
		[]float64{10, 20},
	)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("nil result")
	}
	if result.PositionEstimate == nil {
		t.Fatal("nil position estimate")
	}

	// RMS should be √(1+1) = √2 ≈ 1.414
	expectedRMS := math.Sqrt(2.0)
	if math.Abs(result.RMSAccuracyNM-expectedRMS) > 0.01 {
		t.Errorf("expected RMS ≈ %.3f, got %.3f",
			expectedRMS, result.RMSAccuracyNM)
	}
}

func TestCompute3DNavigationSolution_RVSM(t *testing.T) {
	measurements := []types.Measurement{
		{AzimuthDeg: 0, ElevationDeg: 5, SigmaNM: 0.0986},
		{AzimuthDeg: 90, ElevationDeg: 5, SigmaNM: 0.0986},
	}

	altSigma := 20.0 / 1852.0 // RVSM: 20m in NM

	result, err := engine.Compute3DNavigationSolution(
		measurements,
		altSigma,
		true,
	)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("nil result")
	}
	if result.RMSAccuracyNM <= 0 {
		t.Errorf("expected positive RMS, got %f", result.RMSAccuracyNM)
	}
	if math.Abs(result.VerticalAccuracyNM-altSigma) > 0.001 {
		t.Errorf("expected vertical ≈ %f, got %f",
			altSigma, result.VerticalAccuracyNM)
	}

	expected2Sigma := 2.0 * result.RMSAccuracyNM
	if math.Abs(result.TwoSigmaNM-expected2Sigma) > 1e-9 {
		t.Errorf("2σ mismatch")
	}
}

func TestCompute3DNavigationSolution_FixedAltitude(t *testing.T) {
	measurements := []types.Measurement{
		{AzimuthDeg: 0, ElevationDeg: 5, SigmaNM: 0.0986},
		{AzimuthDeg: 90, ElevationDeg: 5, SigmaNM: 0.0986},
	}

	fixedSigma := 1e-6

	result, err := engine.Compute3DNavigationSolution(
		measurements,
		fixedSigma,
		true,
	)
	if err != nil {
		t.Fatal(err)
	}
	if result.RMSAccuracyNM <= 0 {
		t.Errorf("expected positive RMS, got %f", result.RMSAccuracyNM)
	}
	if result.VerticalAccuracyNM > 0.001 {
		t.Errorf("expected near-zero vertical, got %f",
			result.VerticalAccuracyNM)
	}
}
