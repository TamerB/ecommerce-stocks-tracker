package handler

import (
	"github.com/TamerB/ecommerce-stocks-tracker/api/restapi/operations/health"
	"github.com/go-openapi/runtime/middleware"
)

// HealthzRequestHandler get healthz API
type HealthzRequestHandler struct{}

// NewHealthzRequestHandler creates a new http.Handler for the get healthz
func NewHealthzRequestHandler() *HealthzRequestHandler {
	return &HealthzRequestHandler{}
}

// Handle executing the request and returning a response
func (h *HealthzRequestHandler) Handle(params health.GetHealthzParams) middleware.Responder {
	response := health.NewGetHealthzOK()
	response.SetPayload("OK")
	return response
}
