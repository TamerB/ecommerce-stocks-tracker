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

// GetProductStocksRequestHandler get product stocks API
type GetProductStocksRequestHandler struct {
	store db.Store
}

// NewGetProductStocksRequestHandler creates a new http.Handler for getting product stocks
func NewGetProductStocksRequestHandler(store db.Store) *GetProductStocksRequestHandler {
	return &GetProductStocksRequestHandler{
		store: store,
	}
}

// Handle executing the request and returning a response
func (h *GetProductStocksRequestHandler) Handle(params product.GetProductStocksBySkuParams) middleware.Responder {
	productResponse, responder := h.GetProductStocksBySKU(params)

	if productResponse == nil {
		return responder
	}

	response := &models.BaseResponse{
		Success:  true,
		Messages: []*models.Message{{Number: 200, Type: "Success", Text: "Product Stocks details retrieved successfully"}},
		Data:     productResponse,
	}

	responseOK := responder.(*product.GetProductStocksBySkuOK)
	responseOK.SetPayload(response)
	return responseOK
}

// GetProductBySKU is seperated from Handle function for testing
func (h *GetProductStocksRequestHandler) GetProductStocksBySKU(params product.GetProductStocksBySkuParams) (*models.ProductStocks, middleware.Responder) {
	var errMessage models.BaseResponse

	if params.Sku == "" {
		log.Println("error", errorSKUNotPassed)
		errMessage = getErrorMessage(http.StatusBadRequest, errorSKUNotPassed)
		res := product.NewGetProductStocksBySkuBadRequest()
		res.SetPayload(&errMessage)
		return nil, res
	}

	dbProduct, err := h.store.ListProductStocksBySKU(context.Background(), params.Sku)

	if err != nil {
		log.Println("error", err)
		if strings.HasSuffix(err.Error(), "connection refused") {
			errMessage = getErrorMessage(http.StatusInternalServerError, dbConnectionRefused)
		} else {
			errMessage = getErrorMessage(http.StatusNotFound, fmt.Sprintf(errorProductNotFound, params.Sku))
		}
		res := product.NewGetProductStocksBySkuNotFound()
		res.SetPayload(&errMessage)
		return nil, res
	}

	if len(dbProduct) == 0 {
		log.Println("error", fmt.Sprintf("Product with SKU (%s) not found", params.Sku))
		errMessage = getErrorMessage(http.StatusNotFound, fmt.Sprintf(errorProductNotFound, params.Sku))
		res := product.NewGetProductStocksBySkuNotFound()
		res.SetPayload(&errMessage)
		return nil, res
	}

	createdAt := strfmt.DateTime(dbProduct[0].CreatedAt)
	updatedAt := strfmt.DateTime(dbProduct[0].UpdatedAt)

	productResponse := &models.ProductStocks{
		ID:        &dbProduct[0].ID,
		Sku:       &dbProduct[0].Sku,
		Name:      &dbProduct[0].Name,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Stocks:    []*models.CountryStock{},
	}

	for _, val := range dbProduct {
		if !val.Country.Valid || !val.Quantity.Valid {
			continue
		}
		productResponse.Stocks = append(productResponse.Stocks, &models.CountryStock{
			Country:  &val.Country.String,
			Quantity: &val.Quantity.Int64,
		})
	}
	return productResponse, product.NewGetProductStocksBySkuOK()
}
