package simulation_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/simulation-service/pkg/dto"
	"github.com/D2-GCA/PRISM/simulation-service/pkg/handlers"
)

func TestCoverageHandler_Success(t *testing.T) {

	app := fiber.New()

	app.Post(
		"/coverage",
		handlers.HandleCoverageSimulation,
	)

	reqBody := dto.CoverageRequest{
		Stations: []dto.StationDTO{
			{
				ID:              "DME1",
				X:               0.0,
				Y:               50.0,
				ElevationFt:     100.0,
				ServiceVolumeNM: 130.0,
			},
			{
				ID:              "DME2",
				X:               50.0,
				Y:               0.0,
				ElevationFt:     100.0,
				ServiceVolumeNM: 130.0,
			},
		},
		MinX:       -10.0,
		MaxX:       10.0,
		MinY:       -10.0,
		MaxY:       10.0,
		GridStepNM: 5.0,
	}

	bodyBytes, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		"POST",
		"/coverage",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := app.Test(
		req,
		-1,
	)

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			resp.StatusCode,
		)
	}

	var respBody dto.CoverageResponse

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	if len(respBody.Points) == 0 {
		t.Error("expected point results in coverage output")
	}
}

func TestElevationSweepHandler_Success(t *testing.T) {

	app := fiber.New()

	app.Post(
		"/elevation",
		handlers.HandleElevationSweep,
	)

	reqBody := dto.ElevationSweepRequest{
		Sigma1NM:           0.0986,
		Sigma2NM:           0.0986,
		AltitudeMode:       "RVSM",
		InclusionAnglesDeg: []float64{90.0},
		ElevationMinDeg:    0.0,
		ElevationMaxDeg:    10.0,
		ElevationStepDeg:   5.0,
	}

	bodyBytes, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		"POST",
		"/elevation",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := app.Test(
		req,
		-1,
	)

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			resp.StatusCode,
		)
	}

	var respBody dto.ElevationSweepResponse

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	if respBody.AltitudeMode != "RVSM" {
		t.Errorf(
			"expected RVSM, got %s",
			respBody.AltitudeMode,
		)
	}

	if len(respBody.Points) == 0 {
		t.Error("expected sweep points in output")
	}
}
