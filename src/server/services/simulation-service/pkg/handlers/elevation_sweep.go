package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/simulation-service/pkg/dto"
	apperrors "github.com/D2-GCA/PRISM/simulation-service/pkg/errors"
	"github.com/D2-GCA/PRISM/simulation-service/pkg/service"
	"github.com/D2-GCA/PRISM/simulation-service/pkg/validation"
)

// HandleElevationSweep handles POST /api/v1/simulation/elevation requests.
//
// Runs an elevation angle sweep simulation, computing horizontal accuracy
// vs. elevation angle for a given altitude mode (IDEA.md §3.2–3.4).
func HandleElevationSweep(c *fiber.Ctx) error {
	var req dto.ElevationSweepRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		)
	}

	if err := validation.ValidateElevationSweepRequest(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			apperrors.APIError{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		)
	}

	result, err := service.RunElevationAngleSweep(&req)
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
