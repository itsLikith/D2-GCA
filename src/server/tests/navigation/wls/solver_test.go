package wls_test

import (
	"math"
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/geometry"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
	"github.com/D2-GCA/PRISM/pkg/navigation/wls"
)

func TestSolve_OrthogonalStations(t *testing.T) {

	measurements := []types.Measurement{
		{
			AzimuthDeg: 0,
			SigmaNM:    1,
		},
		{
			AzimuthDeg: 90,
			SigmaNM:    1,
		},
	}

	g := geometry.BuildGeometryMatrix2D(
		measurements,
	)

	w, err := wls.BuildWeightMatrix(
		measurements,
	)

	if err != nil {
		t.Fatal(err)
	}

	b := &types.ObservationVector{
		Values: []float64{10, 20},
	}

	solution, err := wls.Solve(g, w, b)

	if err != nil {
		t.Fatal(err)
	}

	if solution == nil {
		t.Fatal("nil solution")
	}

	// With α₁=0° and α₂=90°, G = [[0,1],[1,0]]
	// x̂ = (GᵀWG)⁻¹ GᵀWb
	// For identity W: x̂ = G⁻¹ b = [20, 10]
	if math.Abs(solution.EastingNM-20.0) > 1e-9 {
		t.Errorf("expected EastingNM=20, got %f", solution.EastingNM)
	}

	if math.Abs(solution.NorthingNM-10.0) > 1e-9 {
		t.Errorf("expected NorthingNM=10, got %f", solution.NorthingNM)
	}
}

func TestSolve_ObservationMismatch(t *testing.T) {

	measurements := []types.Measurement{
		{AzimuthDeg: 0, SigmaNM: 1},
		{AzimuthDeg: 90, SigmaNM: 1},
	}

	g := geometry.BuildGeometryMatrix2D(measurements)
	w, _ := wls.BuildWeightMatrix(measurements)

	b := &types.ObservationVector{
		Values: []float64{10}, // Wrong count
	}

	_, err := wls.Solve(g, w, b)

	if err == nil {
		t.Fatal("expected error for mismatched observations")
	}
}
