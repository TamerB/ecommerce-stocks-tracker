package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/TamerB/ecommerce-stocks-tracker/api/models"
	"github.com/TamerB/ecommerce-stocks-tracker/api/restapi/operations/stock"
	db "github.com/TamerB/ecommerce-stocks-tracker/db/sqlc"
	"github.com/go-openapi/runtime/middleware"
)

// ConsumeProductStockHandler consumes product's country stock API
type ConsumeProductStockHandler struct {
	store db.Store
}

// NewConsumeProductStockHandler creates new http.Handler for ConsumeProductStockHandler
func NewConsumeProductStockHandler(store db.Store) *ConsumeProductStockHandler {
	return &ConsumeProductStockHandler{
		store: store,
	}
}

// Handle executing the request and returning a response
func (h *ConsumeProductStockHandler) Handle(params stock.ConsumeProductStockParams) middleware.Responder {
	consumeResponse, responder := h.ConsumeProductStock(params)

	if consumeResponse == nil {
		return responder
	}

	responseOK := responder.(*stock.ConsumeProductStockOK)
	responseOK.SetPayload(consumeResponse)
	return responseOK
}

// ConsumeProductStock is seperated from Handle function for testing
func (h *ConsumeProductStockHandler) ConsumeProductStock(params stock.ConsumeProductStockParams) (*models.BaseResponse, middleware.Responder) {
	var errMessage models.BaseResponse
	if params.Sku == "" {
		log.Println("error", errorSKUNotPassed)
		errMessage = getErrorMessage(http.StatusBadRequest, errorSKUNotPassed)
		res := stock.NewConsumeProductStockBadRequest()
		res.SetPayload(&errMessage)
		return nil, res
	}

	if params.Country == "" {
		log.Println("error", errorCountryNotPassed)
		errMessage = getErrorMessage(http.StatusBadRequest, errorCountryNotPassed)
		res := stock.NewConsumeProductStockBadRequest()
		res.SetPayload(&errMessage)
		return nil, res
	}

	if params.ConsumeProductStockParams == nil || params.ConsumeProductStockParams.Quantity == nil {
		log.Println("error", errorQuantityNotPassed)
		errMessage = getErrorMessage(http.StatusBadRequest, errorQuantityNotPassed)
		res := stock.NewConsumeProductStockBadRequest()
		res.SetPayload(&errMessage)
		return nil, res
	}

	err := h.store.ConsumeStockTx(context.Background(), db.UpdateStockTxParams{
		ProductSKU:  params.Sku,
		CountryCode: params.Country,
		Quantity:    *params.ConsumeProductStockParams.Quantity,
	})

	if err != nil {
		var res *stock.ConsumeProductStockDefault
		log.Println("error", err)
		switch err.Error() {
		case "not enough stock":
			errMessage = getErrorMessage(http.StatusConflict, err.Error())
			res = stock.NewConsumeProductStockDefault(400)
			break
		case "stock not found":
			errMessage = getErrorMessage(http.StatusNotFound, err.Error())
			res = stock.NewConsumeProductStockDefault(404)
			break
		default:
			errMessage = getErrorMessage(http.StatusInternalServerError, "Something went wrong")
			res = stock.NewConsumeProductStockDefault(500)
		}
		res.SetPayload(&errMessage)
		return nil, res
	}

	stockConsumeResponse := &models.BaseResponse{
		Success: true,
		Messages: []*models.Message{
			{
				Number: 200,
				Type:   "Success",
				Text:   "Stock was consumed successfully",
			},
		},
	}

	return stockConsumeResponse, stock.NewConsumeProductStockOK()
}
