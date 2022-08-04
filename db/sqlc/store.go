package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	ConsumeStockTx(context.Context, UpdateStockTxParams) error
}

// Store provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function withing a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// UpdateStockTxParams contains the input parameter of the stock update transaction
type UpdateStockTxParams struct {
	ProductSKU  string `json:"product_sku"`
	CountryCode string `json:"country_code"`
	Quantity    int64  `json:"quantity"`
}

// ConsumeStockTx performs a stock consuming transaction.
// It validates that the stock exists and will not be negative after reducing it.
func (store *SQLStore) ConsumeStockTx(ctx context.Context, arg UpdateStockTxParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		stock, err := q.GetStockByProductSKUAndCountryCodeForUpdate(ctx, GetStockByProductSKUAndCountryCodeForUpdateParams{
			Sku:     arg.ProductSKU,
			Country: arg.CountryCode,
		})
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				return errors.New("stock not found")
			}
			return err
		}

		if stock.ID < 0 {
			return errors.New("stock not found")
		}

		if stock.Quantity-arg.Quantity < 0 {
			return errors.New("not enough stock")
		}

		err = q.ConsumeStock(ctx, ConsumeStockParams{
			ID:       stock.ID,
			Quantity: arg.Quantity,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
