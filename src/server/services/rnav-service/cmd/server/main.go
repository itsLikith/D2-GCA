// RNAV Service — Provides 2D and 3D DME/DME RNAV positioning analysis
// endpoints for the PRISM system.
//
// Reference: IDEA.md §2 (2D analysis) and §3 (3D analysis with altitude).
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/D2-GCA/PRISM/rnav-service/pkg/config"
	"github.com/D2-GCA/PRISM/rnav-service/pkg/handlers"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api/v1")
	rnav := api.Group("/rnav")

	// Health check endpoint.
	rnav.Get("/health", handlers.Health)

	// POST /api/v1/rnav/analyze — 2D WLS positioning analysis (§2).
	rnav.Post("/analyze", handlers.HandleRNAV2DAnalysis)

	// POST /api/v1/rnav/analyze3d — 3D positioning with altitude (§3).
	rnav.Post("/analyze3d", handlers.HandleRNAV3DAnalysis)

	cfg := config.Load()
	log.Fatal(app.Listen(":" + cfg.Port))
}
