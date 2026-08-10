package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/D2-GCA/PRISM/gateway/pkg/config"
	"github.com/D2-GCA/PRISM/gateway/pkg/handlers"
	"github.com/D2-GCA/PRISM/gateway/pkg/proxy"
)

func main() {

	cfg := config.Load()

	app := fiber.New()

	app.Use(logger.New())

	//--------------------------------------------------
	// Health
	//--------------------------------------------------

	app.Get(
		"/api/v1/gateway/health",
		handlers.Health,
	)

	//--------------------------------------------------
	// Route: /api/v1/dme/* → DME Service
	//--------------------------------------------------

	app.All(
		"/api/v1/dme/*",
		proxy.Forward(cfg.DMEServiceURL),
	)

	//--------------------------------------------------
	// Route: /api/v1/rnav/* → RNAV Service
	//--------------------------------------------------

	app.All(
		"/api/v1/rnav/*",
		proxy.Forward(cfg.RNAVServiceURL),
	)

	//--------------------------------------------------
	// Route: /api/v1/simulation/* → Simulation Service
	//--------------------------------------------------

	app.All(
		"/api/v1/simulation/*",
		proxy.Forward(cfg.SimulationServiceURL),
	)

	log.Fatal(
		app.Listen(":" + cfg.Port),
	)
}
