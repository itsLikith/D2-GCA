// Simulation Service — Provides spatial coverage analysis and
// elevation angle sweep simulation endpoints for the PRISM system.
//
// Reference: IDEA.md §2.3 (coverage maps) and §3 (elevation sweep).
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/D2-GCA/PRISM/simulation-service/pkg/config"
	"github.com/D2-GCA/PRISM/simulation-service/pkg/handlers"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api/v1")
	sim := api.Group("/simulation")

	// Health check endpoint.
	sim.Get("/health", handlers.Health)

	// POST /api/v1/simulation/coverage — Spatial coverage analysis (§2.3).
	sim.Post("/coverage", handlers.HandleCoverageSimulation)

	// POST /api/v1/simulation/elevation — Elevation angle sweep (§3.2–3.4).
	sim.Post("/elevation", handlers.HandleElevationSweep)

	cfg := config.Load()
	log.Fatal(app.Listen(":" + cfg.Port))
}
