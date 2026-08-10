package gateway_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"github.com/D2-GCA/PRISM/gateway/pkg/handlers"
)

func TestHealthHandler_Success(t *testing.T) {

	app := fiber.New()

	app.Get(
		"/api/v1/gateway/health",
		handlers.Health,
	)

	req := httptest.NewRequest(
		"GET",
		"/api/v1/gateway/health",
		nil,
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

	var respBody map[string]string

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	if respBody["status"] != "UP" {
		t.Errorf(
			"expected status UP, got %s",
			respBody["status"],
		)
	}

	if respBody["service"] != "gateway" {
		t.Errorf(
			"expected service gateway, got %s",
			respBody["service"],
		)
	}
}
