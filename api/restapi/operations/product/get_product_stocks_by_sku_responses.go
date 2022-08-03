// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/TamerB/ecommerce-stocks-tracker/api/models"
)

// GetProductStocksBySkuOKCode is the HTTP code returned for type GetProductStocksBySkuOK
const GetProductStocksBySkuOKCode int = 200

/*GetProductStocksBySkuOK successful operation

swagger:response getProductStocksBySkuOK
*/
type GetProductStocksBySkuOK struct {

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewGetProductStocksBySkuOK creates GetProductStocksBySkuOK with default headers values
func NewGetProductStocksBySkuOK() *GetProductStocksBySkuOK {

	return &GetProductStocksBySkuOK{}
}

// WithPayload adds the payload to the get product stocks by sku o k response
func (o *GetProductStocksBySkuOK) WithPayload(payload *models.BaseResponse) *GetProductStocksBySkuOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get product stocks by sku o k response
func (o *GetProductStocksBySkuOK) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductStocksBySkuOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProductStocksBySkuBadRequestCode is the HTTP code returned for type GetProductStocksBySkuBadRequest
const GetProductStocksBySkuBadRequestCode int = 400

/*GetProductStocksBySkuBadRequest Invalid SKU

swagger:response getProductStocksBySkuBadRequest
*/
type GetProductStocksBySkuBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewGetProductStocksBySkuBadRequest creates GetProductStocksBySkuBadRequest with default headers values
func NewGetProductStocksBySkuBadRequest() *GetProductStocksBySkuBadRequest {

	return &GetProductStocksBySkuBadRequest{}
}

// WithPayload adds the payload to the get product stocks by sku bad request response
func (o *GetProductStocksBySkuBadRequest) WithPayload(payload *models.BaseResponse) *GetProductStocksBySkuBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get product stocks by sku bad request response
func (o *GetProductStocksBySkuBadRequest) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductStocksBySkuBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProductStocksBySkuNotFoundCode is the HTTP code returned for type GetProductStocksBySkuNotFound
const GetProductStocksBySkuNotFoundCode int = 404

/*GetProductStocksBySkuNotFound Product not found

swagger:response getProductStocksBySkuNotFound
*/
type GetProductStocksBySkuNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewGetProductStocksBySkuNotFound creates GetProductStocksBySkuNotFound with default headers values
func NewGetProductStocksBySkuNotFound() *GetProductStocksBySkuNotFound {

	return &GetProductStocksBySkuNotFound{}
}

// WithPayload adds the payload to the get product stocks by sku not found response
func (o *GetProductStocksBySkuNotFound) WithPayload(payload *models.BaseResponse) *GetProductStocksBySkuNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get product stocks by sku not found response
func (o *GetProductStocksBySkuNotFound) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductStocksBySkuNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetProductStocksBySkuDefault Return response with error other than the defined ones

swagger:response getProductStocksBySkuDefault
*/
type GetProductStocksBySkuDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewGetProductStocksBySkuDefault creates GetProductStocksBySkuDefault with default headers values
func NewGetProductStocksBySkuDefault(code int) *GetProductStocksBySkuDefault {
	if code <= 0 {
		code = 500
	}

	return &GetProductStocksBySkuDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get product stocks by sku default response
func (o *GetProductStocksBySkuDefault) WithStatusCode(code int) *GetProductStocksBySkuDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get product stocks by sku default response
func (o *GetProductStocksBySkuDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get product stocks by sku default response
func (o *GetProductStocksBySkuDefault) WithPayload(payload *models.BaseResponse) *GetProductStocksBySkuDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get product stocks by sku default response
func (o *GetProductStocksBySkuDefault) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductStocksBySkuDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
