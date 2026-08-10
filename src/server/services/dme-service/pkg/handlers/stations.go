package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/dme-service/pkg/db"
	"github.com/D2-GCA/PRISM/dme-service/pkg/dto"
	apperrors "github.com/D2-GCA/PRISM/dme-service/pkg/errors"
)

// HandleGetStations handles GET /api/v1/dme/stations requests.
func HandleGetStations(c *fiber.Ctx) error {
	database := db.GetDB()
	if database == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			apperrors.APIError{
				Code:    "DATABASE_NOT_INITIALIZED",
				Message: "Database connection is not initialized",
			},
		)
	}

	rows, err := database.Query("SELECT id, name, x, y, elevation_ft, service_volume_nm FROM dme_stations ORDER BY id")
	if err != nil {
		log.Printf("[DATABASE ERROR] Failed to query dme_stations: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(
			apperrors.APIError{
				Code:    "QUERY_ERROR",
				Message: "Failed to retrieve stations from database",
			},
		)
	}
	defer rows.Close()

	stations := make([]dto.StationDTO, 0)
	for rows.Next() {
		var s dto.StationDTO
		if err := rows.Scan(&s.ID, &s.Name, &s.X, &s.Y, &s.ElevationFt, &s.ServiceVolumeNM); err != nil {
			log.Printf("[DATABASE ERROR] Failed to scan row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(
				apperrors.APIError{
					Code:    "SCAN_ERROR",
					Message: "Failed to parse station row data",
				},
			)
		}
		stations = append(stations, s)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[DATABASE ERROR] Rows iterator error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(
			apperrors.APIError{
				Code:    "ITERATION_ERROR",
				Message: "Error iterating station records",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(stations)
}
