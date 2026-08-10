// Package types defines the core data structures used throughout the
// navigation computation pipeline for DME/DME RNAV positioning analysis.
//
// These types directly correspond to the mathematical constructs described
// in the paper: geometry matrices (G), weight matrices (W), covariance
// matrices, observation vectors, measurements, and positioning solutions.
package types

import (
	"gonum.org/v1/gonum/mat"
)

// GeometryMatrix wraps the direction vector matrix G that relates
// DME station observations to the aircraft's position error.
//
// Reference: IDEA.md §2.1, Eq. 6 — For 2D positioning:
//
//	G = [sinα₁ cosα₁; sinα₂ cosα₂; ...; sinαₙ cosαₙ]
//
// Reference: IDEA.md §3.1, Eq. 19 — For 3D positioning:
//
//	G = [cosθ₁·sinα₁ cosθ₁·cosα₁ sinθ₁; ...; cosθₙ·sinαₙ cosθₙ·cosαₙ sinθₙ]
//
// where αᵢ is the azimuth angle and θᵢ is the elevation angle of the
// i-th DME station relative to the aircraft.
type GeometryMatrix struct {
	Matrix *mat.Dense
}

// WeightMatrix wraps the accuracy weight factor matrix W that gives
// higher weight to DME stations with smaller ranging errors.
//
// Reference: IDEA.md §2.2, Eq. 13:
//
//	W = R⁻¹ = diag(σ₁⁻², σ₂⁻², ..., σₙ⁻²)
//
// where σᵢ is the ranging error of the i-th DME station.
type WeightMatrix struct {
	Matrix *mat.Dense
}

// CovarianceMatrix wraps the error covariance matrix of the WLS
// position estimate.
//
// Reference: IDEA.md §2.2, Eq. 14:
//
//	cov(x̂) = (GᵀWG)⁻¹
//
// The diagonal elements u₁₁ and u₂₂ represent the variance of the
// Easting and Northing errors respectively. For 3D, u₃₃ represents
// the vertical error variance.
type CovarianceMatrix struct {
	Matrix *mat.Dense
}
