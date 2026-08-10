package covariance_test

import (
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/covariance"
	"github.com/D2-GCA/PRISM/pkg/navigation/geometry"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
	"github.com/D2-GCA/PRISM/pkg/navigation/wls"
)

func TestComputeFromGeometryAndWeights_2D_OrthogonalStations(t *testing.T) {
	measurements := []types.Measurement{
		{AzimuthDeg: 0, SigmaNM: 1},
		{AzimuthDeg: 90, SigmaNM: 1},
	}

	g := geometry.BuildGeometryMatrix2D(measurements)
	w, err := wls.BuildWeightMatrix(measurements)
	if err != nil {
		t.Fatal(err)
	}

	cov, err := covariance.ComputeFromGeometryAndWeights(g, w)
	if err != nil {
		t.Fatal(err)
	}

	if cov == nil {
		t.Fatal("nil covariance")
	}

	r, c := cov.Matrix.Dims()
	if r != 2 || c != 2 {
		t.Fatalf("expected 2x2, got %dx%d", r, c)
	}

	// For orthogonal directions with equal sigma=1:
	// cov = I (identity)
	u11 := cov.Matrix.At(0, 0)
	u22 := cov.Matrix.At(1, 1)

	if u11 < 0.99 || u11 > 1.01 {
		t.Errorf("expected u11≈1, got %f", u11)
	}
	if u22 < 0.99 || u22 > 1.01 {
		t.Errorf("expected u22≈1, got %f", u22)
	}
}

func TestComputeFromGeometryAndWeights_3D_WithAltimeter(t *testing.T) {
	measurements := []types.Measurement{
		{AzimuthDeg: 0, ElevationDeg: 5, SigmaNM: 0.1},
		{AzimuthDeg: 90, ElevationDeg: 5, SigmaNM: 0.1},
	}

	g := geometry.BuildGeometryMatrix3DWithAltimeter(measurements)

	allMeasurements := append(
		measurements,
		types.Measurement{SigmaNM: 0.01},
	)

	w, err := wls.BuildWeightMatrix(allMeasurements)
	if err != nil {
		t.Fatal(err)
	}

	cov, err := covariance.ComputeFromGeometryAndWeights(g, w)
	if err != nil {
		t.Fatal(err)
	}

	r, c := cov.Matrix.Dims()
	if r != 3 || c != 3 {
		t.Fatalf("expected 3x3, got %dx%d", r, c)
	}
}
