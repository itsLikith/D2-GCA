package wls_test

import (
	"math"
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/types"
	"github.com/D2-GCA/PRISM/pkg/navigation/wls"
)

func TestBuildWeightMatrix_CorrectDiagonal(t *testing.T) {

	input := []types.Measurement{
		{SigmaNM: 0.5},
		{SigmaNM: 1.0},
	}

	w, err := wls.BuildWeightMatrix(input)

	if err != nil {
		t.Fatal(err)
	}

	// W = diag(1/σ²) = diag(4, 1)
	v1 := w.Matrix.At(0, 0)
	v2 := w.Matrix.At(1, 1)

	if math.Abs(v1-4.0) > 1e-9 {
		t.Fatalf("expected 4.0, got %f", v1)
	}

	if math.Abs(v2-1.0) > 1e-9 {
		t.Fatalf("expected 1.0, got %f", v2)
	}

	// Off-diagonal should be zero
	if w.Matrix.At(0, 1) != 0 {
		t.Fatal("off-diagonal should be zero")
	}
}

func TestBuildWeightMatrix_ZeroSigmaError(t *testing.T) {

	input := []types.Measurement{
		{SigmaNM: 0},
	}

	_, err := wls.BuildWeightMatrix(input)

	if err == nil {
		t.Fatal("expected error for zero sigma")
	}
}

func TestBuildWeightMatrix_NegativeSigmaError(t *testing.T) {

	input := []types.Measurement{
		{SigmaNM: -1.0},
	}

	_, err := wls.BuildWeightMatrix(input)

	if err == nil {
		t.Fatal("expected error for negative sigma")
	}
}
