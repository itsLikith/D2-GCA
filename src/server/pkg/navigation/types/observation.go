package types

// ObservationVector wraps the vector b of distance observation errors
// from the RNAV equipment to n DME beacon stations.
//
// Reference: IDEA.md §2.1, Eq. 5:
//
//	b = Gx + v
//
// where b = [Δr₁, Δr₂, ..., Δrₙ]ᵀ is the horizontal projection of
// the distance observation errors, and v is the random noise vector.
type ObservationVector struct {
	Values []float64
}
