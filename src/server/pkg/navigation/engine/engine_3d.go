package engine

import (
	"github.com/D2-GCA/PRISM/pkg/navigation/accuracy"
	"github.com/D2-GCA/PRISM/pkg/navigation/compliance"
	"github.com/D2-GCA/PRISM/pkg/navigation/covariance"
	"github.com/D2-GCA/PRISM/pkg/navigation/geometry"
	"github.com/D2-GCA/PRISM/pkg/navigation/types"
	"github.com/D2-GCA/PRISM/pkg/navigation/wls"
)

// Compute3DNavigationSolution performs a 3D DME/DME RNAV positioning
// analysis per IDEA.md §3.1.
//
// This extends the 2D analysis (§2) by incorporating altitude information
// from an external sensor (barometric altimeter or fixed altitude).
//
// The pipeline:
//  1. Build the 3D geometry matrix G (Eq. 19), optionally with an
//     altimeter pseudo-observation row [0,0,1] (Eq. 22)
//  2. Build weight matrix W including the altimeter's σ₃ (Eq. 13)
//  3. Compute covariance: cov(x̂) = (GᵀWG)⁻¹ (Eq. 20)
//  4. Extract horizontal and vertical accuracy (Eq. 21):
//     σ_AZ = √(u₁₁+u₂₂), σ_EL = √(u₃₃)
//  5. Check RNAV compliance
//
// When includeAltimeter is true, an extra row [0,0,1] is appended to
// the G matrix, and the altimeter's sigma is included in W. This models
// the three operating modes:
//   - DME/DME + RVSM altimeter (§3.2, σ₃ = 20m)
//   - DME/DME + CVSM altimeter (§3.4, σ₃ = 50m)
//   - DME/DME + fixed altitude  (§3.3, σ₃ ≈ 0)
func Compute3DNavigationSolution(
	measurements []types.Measurement,
	altimeterSigmaNM float64,
	includeAltimeter bool,
) (*types.NavigationSolution, error) {

	// ──────────────────────────────────────────────────
	// Step 1: Build 3D Geometry Matrix (Eq. 19 or Eq. 22)
	// ──────────────────────────────────────────────────

	var geometryMatrix *types.GeometryMatrix

	if includeAltimeter {
		// With altimeter: G has n+1 rows, last row is [0,0,1] (Eq. 22).
		geometryMatrix = geometry.BuildGeometryMatrix3DWithAltimeter(measurements)
	} else {
		// Without altimeter: G has n rows (Eq. 19).
		geometryMatrix = geometry.BuildGeometryMatrix3D(measurements)
	}

	// ──────────────────────────────────────────────────
	// Step 2: Build Weight Matrix W (Eq. 13)
	//
	// Include the altimeter as an additional measurement
	// with its own σ₃ in the weight matrix.
	// ──────────────────────────────────────────────────

	allMeasurements := make([]types.Measurement, len(measurements))
	copy(allMeasurements, measurements)

	if includeAltimeter {
		allMeasurements = append(allMeasurements, types.Measurement{
			SigmaNM: altimeterSigmaNM,
		})
	}

	weightMatrix, err := wls.BuildWeightMatrix(allMeasurements)
	if err != nil {
		return nil, err
	}

	// ──────────────────────────────────────────────────
	// Step 3: Compute Covariance cov(x̂) = (GᵀWG)⁻¹ (Eq. 20)
	// ──────────────────────────────────────────────────

	covarianceMatrix, err := covariance.ComputeFromGeometryAndWeights(
		geometryMatrix, weightMatrix,
	)
	if err != nil {
		return nil, err
	}

	// ──────────────────────────────────────────────────
	// Step 4: Extract Horizontal and Vertical Accuracy (Eq. 21)
	// ──────────────────────────────────────────────────

	accuracyResult, err := accuracy.HorizontalAndVerticalRMSFromCovariance3D(
		covarianceMatrix,
	)
	if err != nil {
		return nil, err
	}

	twoSigmaNM := accuracy.ConvertToTwoSigma(accuracyResult.HorizontalRMSNM)

	// ──────────────────────────────────────────────────
	// Step 5: Check RNAV Compliance
	// ──────────────────────────────────────────────────

	rnav1Result := compliance.CheckRNAV1Compliance(twoSigmaNM)
	rnav2Result := compliance.CheckRNAV2Compliance(twoSigmaNM)

	return &types.NavigationSolution{
		Covariance:         covarianceMatrix,
		RMSAccuracyNM:      accuracyResult.HorizontalRMSNM,
		VerticalAccuracyNM: accuracyResult.VerticalRMSNM,
		TwoSigmaNM:         twoSigmaNM,
		RNAV1:              rnav1Result.IsCompliant,
		RNAV2:              rnav2Result.IsCompliant,
	}, nil
}
