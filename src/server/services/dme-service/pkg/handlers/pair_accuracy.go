package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/dme-service/pkg/dto"
	apperrors "github.com/D2-GCA/PRISM/dme-service/pkg/errors"
	"github.com/D2-GCA/PRISM/dme-service/pkg/service"
	"github.com/D2-GCA/PRISM/dme-service/pkg/validation"
)

// HandleDMEPairAccuracy handles POST /api/v1/dme/pair-accuracy requests.
//
// It computes the horizontal positioning accuracy for a specific DME/DME
// station pair configuration using the closed-form Eq. 17/18 from
// IDEA.md §2.3.
func HandleDMEPairAccuracy(c *fiber.Ctx) error {
	var req dto.PairAccuracyRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		)
	}

	if err := validation.ValidatePairAccuracyRequest(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		)
	}

	result, err := service.ComputeDMEPairAccuracy(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			apperrors.APIError{
				Code:    "COMPUTATION_FAILED",
				Message: err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
