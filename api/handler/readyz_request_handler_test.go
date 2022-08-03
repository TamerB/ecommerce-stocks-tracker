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

func TestReadyzRequestHandler(t *testing.T) {
	t.Run("response with 200 and ok in body", func(t *testing.T) {
		handler := handler.NewReadyzRequestHandler()
		response := httptest.NewRecorder()

		responder := handler.Handle(health.GetReadyzParams{})

		responder.WriteResponse(response, runtime.TextProducer())

		assert.Equal(t, response.Code, http.StatusOK)
		assert.Equal(t, response.Body.String(), "OK")
	})
}
