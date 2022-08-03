// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetProductBySkuParams creates a new GetProductBySkuParams object
//
// There are no default values defined in the spec.
func NewGetProductBySkuParams() GetProductBySkuParams {

	return GetProductBySkuParams{}
}

// GetProductBySkuParams contains all the bound params for the get product by sku operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetProductBySku
type GetProductBySkuParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Product SKU
	  Required: true
	  Max Length: 12
	  Pattern: ^[a-z0-9]+$
	  In: path
	*/
	Sku string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetProductBySkuParams() beforehand.
func (o *GetProductBySkuParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSku, rhkSku, _ := route.Params.GetOK("sku")
	if err := o.bindSku(rSku, rhkSku, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSku binds and validates parameter Sku from path.
func (o *GetProductBySkuParams) bindSku(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Sku = raw

	if err := o.validateSku(formats); err != nil {
		return err
	}

	return nil
}

// validateSku carries on validations for parameter Sku
func (o *GetProductBySkuParams) validateSku(formats strfmt.Registry) error {

	if err := validate.MaxLength("sku", "path", o.Sku, 12); err != nil {
		return err
	}

	if err := validate.Pattern("sku", "path", o.Sku, `^[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}
