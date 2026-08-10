package dme_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/dme-service/pkg/dto"
	"github.com/D2-GCA/PRISM/dme-service/pkg/handlers"
)

func TestErrorModelHandler_Success(t *testing.T) {

	app := fiber.New()

	app.Post(
		"/error",
		handlers.HandleDMEErrorModel,
	)

	reqBody := dto.ErrorModelRequest{
		RangeNM: 50.0,
	}

	bodyBytes, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		"POST",
		"/error",
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

	var respBody dto.ErrorModelResponse

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	if respBody.RangeNM != 50.0 {
		t.Errorf(
			"expected range 50.0, got %f",
			respBody.RangeNM,
		)
	}

	if respBody.TotalErrorNM <= 0 {
		t.Errorf(
			"expected positive total error, got %f",
			respBody.TotalErrorNM,
		)
	}
}

func TestPairAccuracyHandler_Success(t *testing.T) {

	app := fiber.New()

	app.Post(
		"/pair-accuracy",
		handlers.HandleDMEPairAccuracy,
	)

	reqBody := dto.PairAccuracyRequest{
		Range1NM:          40.0,
		Range2NM:          60.0,
		InclusionAngleDeg: 90.0,
	}

	bodyBytes, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		"POST",
		"/pair-accuracy",
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

	var respBody dto.PairAccuracyResponse

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	if respBody.InclusionAngleDeg != 90.0 {
		t.Errorf(
			"expected angle 90.0, got %f",
			respBody.InclusionAngleDeg,
		)
	}

	if respBody.RMSNM <= 0 {
		t.Errorf(
			"expected positive RMS, got %f",
			respBody.RMSNM,
		)
	}
}
