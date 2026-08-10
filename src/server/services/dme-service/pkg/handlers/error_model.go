package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/dme-service/pkg/dto"
	apperrors "github.com/D2-GCA/PRISM/dme-service/pkg/errors"
	"github.com/D2-GCA/PRISM/dme-service/pkg/service"
	"github.com/D2-GCA/PRISM/dme-service/pkg/validation"
)

// HandleDMEErrorModel handles POST /api/v1/dme/error requests.
//
// It computes the DME error model components (system, airborne, total)
// for a given range, implementing the AC-91-FS error standard from
// IDEA.md §2.3.
func HandleDMEErrorModel(c *fiber.Ctx) error {
	var req dto.ErrorModelRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		)
	}

	if err := validation.ValidateErrorModelRequest(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		)
	}

	result := service.ComputeDMEErrorModel(&req)

	return c.Status(fiber.StatusOK).JSON(result)
}
