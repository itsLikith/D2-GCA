package dto

import (
	"github.com/D2-GCA/PRISM/pkg/simulation/types"
)

// CoverageResponse is the API response for coverage simulation.
type CoverageResponse struct {
	Points     []types.CoveragePointResult `json:"points"`
	Statistics types.CoverageStatistics    `json:"statistics"`
}

// ElevationSweepResponse is the API response for elevation sweep simulation.
type ElevationSweepResponse struct {
	AltitudeMode string                      `json:"altitudeMode"`
	Points       []types.ElevationSweepPoint `json:"points"`
}
