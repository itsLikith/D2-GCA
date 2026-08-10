package geometry_test

import (
	"math"
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/geometry"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
)

func TestBuildGeometryMatrix2D_Dimensions(t *testing.T) {
	input := []types.Measurement{
		{AzimuthDeg: 0},
		{AzimuthDeg: 90},
	}

	g := geometry.BuildGeometryMatrix2D(input)
	r, c := g.Matrix.Dims()

	if r != 2 {
		t.Fatalf("expected 2 rows, got %d", r)
	}
	if c != 2 {
		t.Fatalf("expected 2 cols, got %d", c)
	}
}

func TestBuildGeometryMatrix2D_Values(t *testing.T) {
	input := []types.Measurement{
		{AzimuthDeg: 0},
		{AzimuthDeg: 90},
	}

	g := geometry.BuildGeometryMatrix2D(input)

	// α=0°: sin(0)=0, cos(0)=1
	if math.Abs(g.Matrix.At(0, 0)-0.0) > 1e-9 {
		t.Errorf("G[0,0] expected 0, got %f", g.Matrix.At(0, 0))
	}
	if math.Abs(g.Matrix.At(0, 1)-1.0) > 1e-9 {
		t.Errorf("G[0,1] expected 1, got %f", g.Matrix.At(0, 1))
	}

	// α=90°: sin(90)=1, cos(90)=0
	if math.Abs(g.Matrix.At(1, 0)-1.0) > 1e-9 {
		t.Errorf("G[1,0] expected 1, got %f", g.Matrix.At(1, 0))
	}
	if math.Abs(g.Matrix.At(1, 1)-0.0) > 1e-9 {
		t.Errorf("G[1,1] expected 0, got %f", g.Matrix.At(1, 1))
	}
}

func TestBuildGeometryMatrix3D_Dimensions(t *testing.T) {
	input := []types.Measurement{
		{AzimuthDeg: 0, ElevationDeg: 5},
		{AzimuthDeg: 90, ElevationDeg: 10},
	}

	g := geometry.BuildGeometryMatrix3D(input)
	r, c := g.Matrix.Dims()

	if r != 2 {
		t.Fatalf("expected 2 rows, got %d", r)
	}
	if c != 3 {
		t.Fatalf("expected 3 cols, got %d", c)
	}
}

func TestBuildGeometryMatrix3D_ZeroElevation(t *testing.T) {
	// At θ=0°: cos(0)=1, sin(0)=0
	// Row should equal [sin(α), cos(α), 0]
	input := []types.Measurement{
		{AzimuthDeg: 45, ElevationDeg: 0},
	}

	g := geometry.BuildGeometryMatrix3D(input)

	sinAz := math.Sin(45 * math.Pi / 180)
	cosAz := math.Cos(45 * math.Pi / 180)

	if math.Abs(g.Matrix.At(0, 0)-sinAz) > 1e-9 {
		t.Errorf("expected sin(45°), got %f", g.Matrix.At(0, 0))
	}
	if math.Abs(g.Matrix.At(0, 1)-cosAz) > 1e-9 {
		t.Errorf("expected cos(45°), got %f", g.Matrix.At(0, 1))
	}
	if math.Abs(g.Matrix.At(0, 2)-0.0) > 1e-9 {
		t.Errorf("expected 0 for elevation component, got %f", g.Matrix.At(0, 2))
	}
}

func TestBuildGeometryMatrix3DWithAltimeter_Dimensions(t *testing.T) {
	input := []types.Measurement{
		{AzimuthDeg: 0, ElevationDeg: 5},
		{AzimuthDeg: 90, ElevationDeg: 10},
	}

	g := geometry.BuildGeometryMatrix3DWithAltimeter(input)
	r, c := g.Matrix.Dims()

	// 2 DME rows + 1 altimeter row = 3
	if r != 3 {
		t.Fatalf("expected 3 rows, got %d", r)
	}
	if c != 3 {
		t.Fatalf("expected 3 cols, got %d", c)
	}
}

func TestBuildGeometryMatrix3DWithAltimeter_LastRow(t *testing.T) {
	input := []types.Measurement{
		{AzimuthDeg: 30, ElevationDeg: 5},
	}

	g := geometry.BuildGeometryMatrix3DWithAltimeter(input)

	r, _ := g.Matrix.Dims()
	lastRow := r - 1

	// Altimeter row: [0, 0, 1]
	if g.Matrix.At(lastRow, 0) != 0 {
		t.Errorf("altimeter row col 0 should be 0")
	}
	if g.Matrix.At(lastRow, 1) != 0 {
		t.Errorf("altimeter row col 1 should be 0")
	}
	if g.Matrix.At(lastRow, 2) != 1 {
		t.Errorf("altimeter row col 2 should be 1")
	}
}
