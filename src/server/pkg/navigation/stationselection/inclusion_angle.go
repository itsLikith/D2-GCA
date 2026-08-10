package stationselection

import (
	"math"

	"github.com/D2-GCA/PRISM/pkg/models"
)

// ComputeInclusionAngle calculates the true azimuth inclusion angle α
// between two DME stations as seen from the aircraft position.
//
// Reference: IDEA.md §2.3 — The inclusion angle α is the angle between
// the two direction vectors from the aircraft to each DME station.
// This angle directly affects positioning accuracy through Eq. 17:
//
//	σ = √(σ₁² + σ₂²) / sin(α)
//
// The angle is computed using the dot product formula:
//
//	α = arccos(v₁·v₂ / (|v₁|·|v₂|))
//
// where v₁ and v₂ are vectors from the aircraft to each station.
func ComputeInclusionAngle(
	aircraft models.Aircraft,
	stationA models.DMEStation,
	stationB models.DMEStation,
) float64 {

	// Vector from aircraft to station A: (ΔE_A, ΔN_A).
	vectorToA_Easting := stationA.Position.EastingNM - aircraft.Position.EastingNM
	vectorToA_Northing := stationA.Position.NorthingNM - aircraft.Position.NorthingNM

	// Vector from aircraft to station B: (ΔE_B, ΔN_B).
	vectorToB_Easting := stationB.Position.EastingNM - aircraft.Position.EastingNM
	vectorToB_Northing := stationB.Position.NorthingNM - aircraft.Position.NorthingNM

	// Dot product: v₁·v₂ = ΔE_A·ΔE_B + ΔN_A·ΔN_B.
	dotProduct := vectorToA_Easting*vectorToB_Easting +
		vectorToA_Northing*vectorToB_Northing

	// Magnitudes: |v₁| and |v₂|.
	magnitudeA := math.Sqrt(
		vectorToA_Easting*vectorToA_Easting +
			vectorToA_Northing*vectorToA_Northing,
	)

	magnitudeB := math.Sqrt(
		vectorToB_Easting*vectorToB_Easting +
			vectorToB_Northing*vectorToB_Northing,
	)

	// cos(α) = v₁·v₂ / (|v₁|·|v₂|), clamped to [-1, 1].
	cosAngle := dotProduct / (magnitudeA * magnitudeB)
	if cosAngle > 1 {
		cosAngle = 1
	}
	if cosAngle < -1 {
		cosAngle = -1
	}

	return math.Acos(cosAngle) * 180 / math.Pi
}
