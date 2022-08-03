package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/TamerB/ecommerce-stocks-tracker/api/models"
	"github.com/TamerB/ecommerce-stocks-tracker/api/restapi/operations/product"
	db "github.com/TamerB/ecommerce-stocks-tracker/db/sqlc"
)

// GetProductRequestHandler get product API
type GetProductRequestHandler struct {
	store db.Store
}

// NewGetProductRequestHandler creates a new http.Handler for getting product
func NewGetProductRequestHandler(store db.Store) *GetProductRequestHandler {
	return &GetProductRequestHandler{
		store: store,
	}
}

// Handle executing the request and returning a response
func (h *GetProductRequestHandler) Handle(params product.GetProductBySkuParams) middleware.Responder {
	productResponse, responder := h.GetProductBySKU(params)

	if productResponse == nil {
		return responder
	}

	response := &models.BaseResponse{
		Success:  true,
		Messages: []*models.Message{{Number: 200, Type: "Success", Text: "Product details retrieved successfully"}},
		Data:     productResponse,
	}

	responseOK := responder.(*product.GetProductBySkuOK)
	responseOK.SetPayload(response)
	return responseOK
}

// GetProductBySKU is seperated from Handle function for testing
func (h *GetProductRequestHandler) GetProductBySKU(params product.GetProductBySkuParams) (*models.Product, middleware.Responder) {
	var errMessage models.BaseResponse

	if params.Sku == "" {
		log.Println("error", errorSKUNotPassed)
		errMessage = getErrorMessage(http.StatusBadRequest, errorSKUNotPassed)
		res := product.NewGetProductBySkuBadRequest()
		res.SetPayload(&errMessage)
		return nil, res
	}

	dbProduct, err := h.store.GetProductBySKU(context.Background(), params.Sku)

	if err != nil {
		log.Println("error", err)
		if strings.HasSuffix(err.Error(), "connection refused") {
			errMessage = getErrorMessage(http.StatusInternalServerError, dbConnectionRefused)
		} else {
			errMessage = getErrorMessage(http.StatusNotFound, fmt.Sprintf(errorProductNotFound, params.Sku))
		}
		res := product.NewGetProductBySkuNotFound()
		res.SetPayload(&errMessage)
		return nil, res
	}

	if dbProduct.ID <= 0 {
		log.Println("error", fmt.Sprintf("Product with SKU (%s) not found", params.Sku))
		errMessage = getErrorMessage(http.StatusNotFound, fmt.Sprintf(errorProductNotFound, params.Sku))
		res := product.NewGetProductBySkuNotFound()
		res.SetPayload(&errMessage)
		return nil, res
	}

	createdAt := strfmt.DateTime(dbProduct.CreatedAt)
	updatedAt := strfmt.DateTime(dbProduct.UpdatedAt)

	productResponse := &models.Product{
		ID:        &dbProduct.ID,
		Sku:       &dbProduct.Sku,
		Name:      &dbProduct.Name,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}
	return productResponse, product.NewGetProductBySkuOK()
}
