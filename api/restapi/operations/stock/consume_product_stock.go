// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConsumeProductStockHandlerFunc turns a function with the right signature into a consume product stock handler
type ConsumeProductStockHandlerFunc func(ConsumeProductStockParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConsumeProductStockHandlerFunc) Handle(params ConsumeProductStockParams) middleware.Responder {
	return fn(params)
}

// ConsumeProductStockHandler interface for that can handle valid consume product stock params
type ConsumeProductStockHandler interface {
	Handle(ConsumeProductStockParams) middleware.Responder
}

// NewConsumeProductStock creates a new http.Handler for the consume product stock operation
func NewConsumeProductStock(ctx *middleware.Context, handler ConsumeProductStockHandler) *ConsumeProductStock {
	return &ConsumeProductStock{Context: ctx, Handler: handler}
}

/* ConsumeProductStock swagger:route PUT /products/{sku}/stocks/{country} Stock consumeProductStock

Consumes a product's stock

Consumes a product's stock by SKU and country

*/
type ConsumeProductStock struct {
	Context *middleware.Context
	Handler ConsumeProductStockHandler
}

func (o *ConsumeProductStock) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewConsumeProductStockParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
