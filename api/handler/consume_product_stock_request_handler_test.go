package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TamerB/ecommerce-stocks-tracker/api/handler"
	"github.com/TamerB/ecommerce-stocks-tracker/api/models"
	"github.com/TamerB/ecommerce-stocks-tracker/api/restapi/operations/stock"
	mockdb "github.com/TamerB/ecommerce-stocks-tracker/db/mock"
	db "github.com/TamerB/ecommerce-stocks-tracker/db/sqlc"
	"github.com/go-openapi/runtime"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProductSKUSuccess(t *testing.T) {
	t.Run("Response 200 if stock is found and request succeeds", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mockdb.NewMockStore(ctrl)

		params := db.UpdateStockTxParams{
			ProductSKU:  RandomString(12),
			CountryCode: RandomString(2),
			Quantity:    RandomInt64(),
		}

		store.EXPECT().ConsumeStockTx(context.Background(), params).Return(nil)

		response := httptest.NewRecorder()
		handler := handler.NewConsumeProductStockHandler(store)

		responder := handler.Handle(stock.ConsumeProductStockParams{
			Sku:     params.ProductSKU,
			Country: params.CountryCode,
			ConsumeProductStockParams: &models.ConsumeProductStockRequestBody{
				Quantity: &params.Quantity,
			},
		})

		responder.WriteResponse(response, runtime.TextProducer())
		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestUpdateProductSKUNotFound(t *testing.T) {
	t.Run("Response 404 if product is not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mockdb.NewMockStore(ctrl)

		params := db.UpdateStockTxParams{
			ProductSKU:  RandomString(12),
			CountryCode: RandomString(2),
			Quantity:    RandomInt64(),
		}

		store.EXPECT().ConsumeStockTx(context.Background(), params).Return(errors.New("stock not found"))

		response := httptest.NewRecorder()
		handler := handler.NewConsumeProductStockHandler(store)

		responder := handler.Handle(stock.ConsumeProductStockParams{
			Sku:     params.ProductSKU,
			Country: params.CountryCode,
			ConsumeProductStockParams: &models.ConsumeProductStockRequestBody{
				Quantity: &params.Quantity,
			},
		})

		responder.WriteResponse(response, runtime.TextProducer())
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestUpdateProductSKUMinus(t *testing.T) {
	t.Run("Response 400 if stock to be consumed is more than what exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mockdb.NewMockStore(ctrl)

		params := db.UpdateStockTxParams{
			ProductSKU:  RandomString(12),
			CountryCode: RandomString(2),
			Quantity:    RandomInt64(),
		}

		store.EXPECT().ConsumeStockTx(context.Background(), params).Return(errors.New("not enough stock"))

		response := httptest.NewRecorder()
		handler := handler.NewConsumeProductStockHandler(store)

		responder := handler.Handle(stock.ConsumeProductStockParams{
			Sku:     params.ProductSKU,
			Country: params.CountryCode,
			ConsumeProductStockParams: &models.ConsumeProductStockRequestBody{
				Quantity: &params.Quantity,
			},
		})

		responder.WriteResponse(response, runtime.TextProducer())
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
