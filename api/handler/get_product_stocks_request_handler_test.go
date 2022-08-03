package handler_test

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/TamerB/ecommerce-stocks-tracker/api/handler"
	"github.com/TamerB/ecommerce-stocks-tracker/api/restapi/operations/product"
	mockdb "github.com/TamerB/ecommerce-stocks-tracker/db/mock"
	db "github.com/TamerB/ecommerce-stocks-tracker/db/sqlc"
	"github.com/go-openapi/runtime"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductStocksBySKUSuccess(t *testing.T) {
	t.Run("Response 200 if product is found and request succeeds", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := mockdb.NewMockStore(ctrl)

		id := RandomInt64()
		sku := RandomString(12)
		name := RandomString(20)
		createdAt := time.Now()
		updatedAt := time.Now()

		randomProduct := []db.ListProductStocksBySKURow{
			{
				ID:        id,
				Sku:       sku,
				Name:      name,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				Country:   sql.NullString{String: RandomString(2), Valid: true},
				Quantity:  sql.NullInt64{Int64: RandomInt64(), Valid: true},
			},
			{
				ID:        id,
				Sku:       sku,
				Name:      name,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				Country:   sql.NullString{String: RandomString(2), Valid: true},
				Quantity:  sql.NullInt64{Int64: RandomInt64(), Valid: true},
			},
			{
				ID:        id,
				Sku:       sku,
				Name:      name,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				Country:   sql.NullString{String: RandomString(2), Valid: true},
				Quantity:  sql.NullInt64{Int64: RandomInt64(), Valid: true},
			},
		}

		store.EXPECT().ListProductStocksBySKU(context.Background(), randomProduct[0].Sku).Return(randomProduct, nil)

		response := httptest.NewRecorder()
		handler := handler.NewGetProductStocksRequestHandler(store)

		responder := handler.Handle(product.GetProductStocksBySkuParams{Sku: randomProduct[0].Sku})

		responder.WriteResponse(response, runtime.TextProducer())

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestGetProductStocksBySKUNotPassed(t *testing.T) {
	t.Run("Response 400 if sku parameter is not passed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := mockdb.NewMockStore(ctrl)
		response := httptest.NewRecorder()
		handler := handler.NewGetProductStocksRequestHandler(store)

		responder := handler.Handle(product.GetProductStocksBySkuParams{})

		responder.WriteResponse(response, runtime.TextProducer())
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestGetProductStocksBySKUNotFound(t *testing.T) {
	t.Run("Response 404 if product is not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := mockdb.NewMockStore(ctrl)

		sku := RandomString(12)

		store.EXPECT().ListProductStocksBySKU(context.Background(), sku).Return([]db.ListProductStocksBySKURow{}, errors.New("Product not found"))

		response := httptest.NewRecorder()
		handler := handler.NewGetProductStocksRequestHandler(store)

		responder := handler.Handle(product.GetProductStocksBySkuParams{Sku: sku})

		responder.WriteResponse(response, runtime.TextProducer())

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}
