package handler

import (
	"github.com/TamerB/ecommerce-stocks-tracker/api/restapi/operations/health"
	"github.com/go-openapi/runtime/middleware"
)

// ReadyzRequestHandler get readyz API
type ReadyzRequestHandler struct{}

// NewReadyzRequestHandler creates a new http.Handler for the get readyz
func NewReadyzRequestHandler() *ReadyzRequestHandler {
	return &ReadyzRequestHandler{}
}

// Handle executing the request and returning a response
func (h *ReadyzRequestHandler) Handle(params health.GetReadyzParams) middleware.Responder {
	response := health.NewGetReadyzOK()
	response.SetPayload("OK")
	return response
}
