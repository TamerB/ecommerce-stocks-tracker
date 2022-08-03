package handler_test

import (
	"context"
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

func TestGetProductBySKUSuccess(t *testing.T) {
	t.Run("Response 200 if product is found and request succeeds", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := mockdb.NewMockStore(ctrl)

		randomProduct := db.Product{
			ID:        RandomInt64(),
			Sku:       RandomString(12),
			Name:      RandomString(20),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		store.EXPECT().GetProductBySKU(context.Background(), randomProduct.Sku).Return(randomProduct, nil)

		response := httptest.NewRecorder()
		handler := handler.NewGetProductRequestHandler(store)

		responder := handler.Handle(product.GetProductBySkuParams{Sku: randomProduct.Sku})

		responder.WriteResponse(response, runtime.TextProducer())

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestGetProductBySKUNotPassed(t *testing.T) {
	t.Run("Response 400 if sku parameter is not passed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := mockdb.NewMockStore(ctrl)
		response := httptest.NewRecorder()
		handler := handler.NewGetProductRequestHandler(store)

		responder := handler.Handle(product.GetProductBySkuParams{})

		responder.WriteResponse(response, runtime.TextProducer())
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestGetProductBySKUNotFound(t *testing.T) {
	t.Run("Response 404 if product is not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := mockdb.NewMockStore(ctrl)

		sku := RandomString(12)

		store.EXPECT().GetProductBySKU(context.Background(), sku).Return(db.Product{}, errors.New("Product not found"))

		response := httptest.NewRecorder()
		handler := handler.NewGetProductRequestHandler(store)

		responder := handler.Handle(product.GetProductBySkuParams{Sku: sku})

		responder.WriteResponse(response, runtime.TextProducer())

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}
