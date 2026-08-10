package rnav_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/rnav-service/pkg/dto"
	"github.com/D2-GCA/PRISM/rnav-service/pkg/handlers"
)

func TestAnalyze3DHandler_Success(t *testing.T) {

	app := fiber.New()

	app.Post(
		"/analyze3d",
		handlers.HandleRNAV3DAnalysis,
	)

	reqBody := dto.Analyze3DRequest{
		Measurements: []dto.Measurement{
			{
				AzimuthDeg:   0,
				ElevationDeg: 5,
				SigmaNM:      0.0986,
			},
			{
				AzimuthDeg:   90,
				ElevationDeg: 5,
				SigmaNM:      0.0986,
			},
		},
		AltitudeMode: "RVSM",
	}

	bodyBytes, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		"POST",
		"/analyze3d",
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

	var respBody dto.Analyze3DResponse

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	if respBody.AltitudeMode != "RVSM" {
		t.Errorf(
			"expected RVSM, got %s",
			respBody.AltitudeMode,
		)
	}

	if respBody.HorizontalRMSNM <= 0 {
		t.Errorf(
			"expected positive horizontal RMS, got %f",
			respBody.HorizontalRMSNM,
		)
	}

	if respBody.VerticalRMSNM <= 0 {
		t.Errorf(
			"expected positive vertical RMS, got %f",
			respBody.VerticalRMSNM,
		)
	}
}

func TestAnalyze3DHandler_ValidationError(t *testing.T) {

	app := fiber.New()

	app.Post(
		"/analyze3d",
		handlers.HandleRNAV3DAnalysis,
	)

	reqBody := dto.Analyze3DRequest{
		Measurements: []dto.Measurement{
			{
				AzimuthDeg:   0,
				ElevationDeg: 5,
				SigmaNM:      0.0986,
			},
		},
		AltitudeMode: "INVALID_MODE",
	}

	bodyBytes, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		"POST",
		"/analyze3d",
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

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf(
			"expected status 400, got %d",
			resp.StatusCode,
		)
	}
}
