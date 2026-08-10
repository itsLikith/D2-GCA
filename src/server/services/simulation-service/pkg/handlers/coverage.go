package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/simulation-service/pkg/dto"
	apperrors "github.com/D2-GCA/PRISM/simulation-service/pkg/errors"
	"github.com/D2-GCA/PRISM/simulation-service/pkg/service"
	"github.com/D2-GCA/PRISM/simulation-service/pkg/validation"
)

// HandleCoverageSimulation handles POST /api/v1/simulation/coverage requests.
//
// Runs a spatial coverage simulation evaluating DME/DME RNAV accuracy
// across a grid, producing coverage maps per IDEA.md §2.3.
func HandleCoverageSimulation(c *fiber.Ctx) error {
	var req dto.CoverageRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		)
	}

	if err := validation.ValidateCoverageRequest(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		)
	}

	result, err := service.RunCoverageSimulation(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			apperrors.APIError{
				Code:    "SIMULATION_FAILED",
				Message: err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
