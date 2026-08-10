package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/rnav-service/pkg/dto"
	apperrors "github.com/D2-GCA/PRISM/rnav-service/pkg/errors"
	"github.com/D2-GCA/PRISM/rnav-service/pkg/service"
	"github.com/D2-GCA/PRISM/rnav-service/pkg/validation"
)

// HandleRNAV2DAnalysis handles POST /api/v1/rnav/analyze requests.
//
// Performs a 2D DME/DME RNAV positioning analysis using the WLS
// algorithm described in IDEA.md §2.
func HandleRNAV2DAnalysis(c *fiber.Ctx) error {
	var req dto.AnalyzeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		)
	}

	if err := validation.ValidateAnalyzeRequest(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		)
	}

	result, err := service.RunRNAV2DAnalysis(&req)
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
