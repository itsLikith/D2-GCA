package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/rnav-service/pkg/dto"
	apperrors "github.com/D2-GCA/PRISM/rnav-service/pkg/errors"
	"github.com/D2-GCA/PRISM/rnav-service/pkg/service"
	"github.com/D2-GCA/PRISM/rnav-service/pkg/validation"
)

// HandleRNAV3DAnalysis handles POST /api/v1/rnav/analyze3d requests.
//
// Performs a 3D DME/DME RNAV positioning analysis with altitude sensor
// integration, implementing the operation accuracy model from IDEA.md §3.1.
func HandleRNAV3DAnalysis(c *fiber.Ctx) error {
	var req dto.Analyze3DRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		)
	}

	if err := validation.ValidateAnalyze3DRequest(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		)
	}

	result, err := service.RunRNAV3DAnalysis(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			apperrors.APIError{
				Code:    "ANALYSIS_FAILED",
				Message: err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
