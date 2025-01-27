package private

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initVerifyRoutes(root fiber.Router) {
	compiler := root.Group("/compiler")
	compiler.Post("/template", h.RunTemplate)
}

// @Tags Codex-Compiler
// @Summary Run Template
// @Description This is a template for an endpoint created to make requests to Codex-Compiler.
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/compiler/template [post]
func (h *PrivateHandler) RunTemplate(c *fiber.Ctx) error {
	// TODO: This is my request template for all.

	// URL to send GET request to
	url := "http://nginx/compiler-api/v1/private/run"
	sessionID := c.Cookies("session_id")

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response.Response(500, "Error creating GET request", err)
	}

	// Add the Codex-Compiler header
	req.Header.Add("Codex-Compiler", "b77759141fc85bf31e75b1d9c48bbe67")

	// Add the session_id cookie to the request
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: sessionID,
	})

	// Create an HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Return error response if request execution fails
		return response.Response(500, "Error making GET request", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Return error response if reading body fails
		return response.Response(500, "Error reading response body", nil)
	}

	var data response.BaseResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error decoding session data:", err)
		return serviceErrors.NewServiceErrorWithMessage(500, "Error decoding session data")
	}

	return response.Response(data.StatusCode, data.Message, data.Data)
}
