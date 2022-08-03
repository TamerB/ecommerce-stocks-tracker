package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TamerB/ecommerce-stocks-tracker/api/handler"
	"github.com/TamerB/ecommerce-stocks-tracker/api/restapi/operations/health"
	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
)

func TestHealthzRequestHandler(t *testing.T) {
	t.Run("response with 200 and ok in body", func(t *testing.T) {
		handler := handler.NewHealthzRequestHandler()
		response := httptest.NewRecorder()

		responder := handler.Handle(health.GetHealthzParams{})

		responder.WriteResponse(response, runtime.TextProducer())

		assert.Equal(t, response.Code, http.StatusOK)
		assert.Equal(t, response.Body.String(), "OK")
	})
}
