package elevation_test

import (
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"
	navtypes "github.com/D2-GCA/PRISM/pkg/navigation/types"
	"github.com/D2-GCA/PRISM/pkg/simulation/elevation"
	simtypes "github.com/D2-GCA/PRISM/pkg/simulation/types"
)

func TestRunSweep_RVSM(t *testing.T) {

	config := simtypes.ElevationSweepConfig{
		DMESigma1NM:           analytical.ICAODMESigmaNM,
		DMESigma2NM:           analytical.ICAODMESigmaNM,
		AltitudeSensorSigmaNM: analytical.RVSMSigmaNM,
		AltitudeMode:          navtypes.AltitudeModeRVSM,

		InclusionAnglesDeg: []float64{30, 60, 90, 120, 150},

		ElevationMinDeg:  0,
		ElevationMaxDeg:  20,
		ElevationStepDeg: 5,
	}

	result, err := elevation.RunElevationAngleSweep(config)

	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatal("nil result")
	}

	if result.AltitudeMode != "RVSM" {
		t.Errorf("expected RVSM, got %s", result.AltitudeMode)
	}

	// 5 angles × 5 elevation steps (0,5,10,15,20) = 25 points
	expectedPoints := 5 * 5

	if len(result.Points) != expectedPoints {
		t.Errorf("expected %d points, got %d",
			expectedPoints, len(result.Points))
	}

	// Verify all 2σ values are positive
	for _, p := range result.Points {

		if p.HorizontalRMS2SigmaNM <= 0 {
			t.Errorf("non-positive 2σ at θ=%.0f° α=%.0f°",
				p.ElevationDeg, p.InclusionAngleDeg)
		}
	}
}

func TestRunSweep_Fixed(t *testing.T) {

	config := simtypes.ElevationSweepConfig{
		DMESigma1NM:           analytical.ICAODMESigmaNM,
		DMESigma2NM:           analytical.ICAODMESigmaNM,
		AltitudeSensorSigmaNM: analytical.FixedAltitudeSigmaNM,
		AltitudeMode:          navtypes.AltitudeModeFixed,

		InclusionAnglesDeg: []float64{90},

		ElevationMinDeg:  0,
		ElevationMaxDeg:  10,
		ElevationStepDeg: 5,
	}

	result, err := elevation.RunElevationAngleSweep(config)

	if err != nil {
		t.Fatal(err)
	}

	if result.AltitudeMode != "FIXED" {
		t.Errorf("expected FIXED, got %s", result.AltitudeMode)
	}

	// All points at α=90° should have 2σ < 1.0 NM (RNAV-1 compliant)
	for _, p := range result.Points {

		if p.HorizontalRMS2SigmaNM >= 1.0 {
			t.Errorf("2σ=%.3f exceeds 1.0 NM at θ=%.0f° α=90°",
				p.HorizontalRMS2SigmaNM, p.ElevationDeg)
		}
	}
}

func TestRunSweep_CVSM(t *testing.T) {

	config := simtypes.ElevationSweepConfig{
		DMESigma1NM:           analytical.ICAODMESigmaNM,
		DMESigma2NM:           analytical.ICAODMESigmaNM,
		AltitudeSensorSigmaNM: analytical.CVSMSigmaNM,
		AltitudeMode:          navtypes.AltitudeModeCVSM,

		InclusionAnglesDeg: []float64{90},

		ElevationMinDeg:  0,
		ElevationMaxDeg:  10,
		ElevationStepDeg: 5,
	}

	result, err := elevation.RunElevationAngleSweep(config)

	if err != nil {
		t.Fatal(err)
	}

	if result.AltitudeMode != "CVSM" {
		t.Errorf("expected CVSM, got %s", result.AltitudeMode)
	}

	// CVSM has larger σ₃ than RVSM, so errors should be larger
	if len(result.Points) == 0 {
		t.Fatal("expected points")
	}
}
