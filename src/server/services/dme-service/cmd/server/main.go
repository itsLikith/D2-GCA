// DME Service — Provides DME ranging error model and pair accuracy
// analysis endpoints for the PRISM system.
//
// Reference: IDEA.md §2.3 — Implements the AC-91-FS DME accuracy
// standard and the closed-form DME/DME horizontal accuracy (Eq. 17/18).
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/D2-GCA/PRISM/dme-service/pkg/config"
	"github.com/D2-GCA/PRISM/dme-service/pkg/db"
	"github.com/D2-GCA/PRISM/dme-service/pkg/handlers"
)

func main() {
	cfg := config.Load()

	// Initialize Database connection pool
	db.Connect(cfg.DatabaseURL)

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api/v1")
	dme := api.Group("/dme")

	// Health check endpoint.
	dme.Get("/health", handlers.Health)

	// GET /api/v1/dme/stations — Retrieve all DME stations from database.
	dme.Get("/stations", handlers.HandleGetStations)

	// POST /api/v1/dme/error — Compute DME error model (§2.3).
	dme.Post("/error", handlers.HandleDMEErrorModel)

	// POST /api/v1/dme/pair-accuracy — Compute DME/DME pair accuracy (§2.3, Eq. 17/18).
	dme.Post("/pair-accuracy", handlers.HandleDMEPairAccuracy)

	log.Fatal(app.Listen(":" + cfg.Port))
}
