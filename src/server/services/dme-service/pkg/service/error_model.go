package service

import (
	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"

	"github.com/D2-GCA/PRISM/dme-service/pkg/dto"
)

// ComputeDMEErrorModel calculates the DME error components for a given range.
//
// Reference: IDEA.md §2.3 — The DME accuracy standard per AC-91-FS defines:
//   - System error: σ_sis = 0.05 NM
//   - Airborne error: σ_air = max(0.05, 0.125% × range)
//   - Total: σ_DME = √(σ_sis² + σ_air²)
func ComputeDMEErrorModel(req *dto.ErrorModelRequest) *dto.ErrorModelResponse {
	airborneErrorNM := analytical.AirborneDMEError(req.RangeNM)
	totalErrorNM := analytical.TotalDMEError(req.RangeNM)

	return &dto.ErrorModelResponse{
		RangeNM:         req.RangeNM,
		SystemErrorNM:   analytical.SystemErrorNM,
		AirborneErrorNM: airborneErrorNM,
		TotalErrorNM:    totalErrorNM,
	}
}
