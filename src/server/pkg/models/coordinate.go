// Package models defines the core domain types used throughout
// the PRISM DME/DME RNAV positioning system.
//
// These types model the physical entities described in the paper:
// aircraft positions, DME ground stations, and coordinate geometry.
package models

// Coordinate represents a 2D position in the horizontal plane using
// Easting (E) and Northing (N) coordinates, measured in nautical miles.
//
// Reference: IDEA.md §2.1 — The aircraft position is denoted (E, N),
// and each DME station position is (Eᵢ, Nᵢ). All horizontal distance
// computations (Eq. 2) operate on these coordinates.
type Coordinate struct {
	// EastingNM is the position along the East axis in nautical miles.
	// Corresponds to the "E" component in the paper's notation.
	EastingNM float64

	// NorthingNM is the position along the North axis in nautical miles.
	// Corresponds to the "N" component in the paper's notation.
	NorthingNM float64
}