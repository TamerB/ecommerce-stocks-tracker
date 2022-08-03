package handler

import (
	"github.com/TamerB/ecommerce-stocks-tracker/api/models"
)

const (
	errorSKUNotPassed      = "Bad request: Product SKU is required"
	errorCountryNotPassed  = "Bad request: Country param is required"
	errorQuantityNotPassed = "Bad request: Quantity body param is required"
	errorFileNotPassed     = "Bad request: CSV file is required"
	errorProductNotFound   = "Product with SKU (%s) not found"
	errDefault             = "Internal server error"
	dbConnectionRefused    = "Connect: DB Connection refused"
	dbSomethingWentWrong   = "DB: Something went wrong"
)

func getErrorMessage(code int, message string) models.BaseResponse {
	return models.BaseResponse{
		Success: false,
		Errors: []*models.Error{
			{
				Number: int64(code),
				Text:   message,
			},
		},
	}
}
