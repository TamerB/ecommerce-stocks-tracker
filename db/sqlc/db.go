// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.consumeStockStmt, err = db.PrepareContext(ctx, consumeStock); err != nil {
		return nil, fmt.Errorf("error preparing query ConsumeStock: %w", err)
	}
	if q.getProductBySKUStmt, err = db.PrepareContext(ctx, getProductBySKU); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductBySKU: %w", err)
	}
	if q.getProductStockCountBySKUStmt, err = db.PrepareContext(ctx, getProductStockCountBySKU); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductStockCountBySKU: %w", err)
	}
	if q.getStockByProductSKUAndCountryCodeForUpdateStmt, err = db.PrepareContext(ctx, getStockByProductSKUAndCountryCodeForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetStockByProductSKUAndCountryCodeForUpdate: %w", err)
	}
	if q.listProductStocksBySKUStmt, err = db.PrepareContext(ctx, listProductStocksBySKU); err != nil {
		return nil, fmt.Errorf("error preparing query ListProductStocksBySKU: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.consumeStockStmt != nil {
		if cerr := q.consumeStockStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing consumeStockStmt: %w", cerr)
		}
	}
	if q.getProductBySKUStmt != nil {
		if cerr := q.getProductBySKUStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductBySKUStmt: %w", cerr)
		}
	}
	if q.getProductStockCountBySKUStmt != nil {
		if cerr := q.getProductStockCountBySKUStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductStockCountBySKUStmt: %w", cerr)
		}
	}
	if q.getStockByProductSKUAndCountryCodeForUpdateStmt != nil {
		if cerr := q.getStockByProductSKUAndCountryCodeForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStockByProductSKUAndCountryCodeForUpdateStmt: %w", cerr)
		}
	}
	if q.listProductStocksBySKUStmt != nil {
		if cerr := q.listProductStocksBySKUStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProductStocksBySKUStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                              DBTX
	tx                                              *sql.Tx
	consumeStockStmt                                *sql.Stmt
	getProductBySKUStmt                             *sql.Stmt
	getProductStockCountBySKUStmt                   *sql.Stmt
	getStockByProductSKUAndCountryCodeForUpdateStmt *sql.Stmt
	listProductStocksBySKUStmt                      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                            tx,
		tx:                            tx,
		consumeStockStmt:              q.consumeStockStmt,
		getProductBySKUStmt:           q.getProductBySKUStmt,
		getProductStockCountBySKUStmt: q.getProductStockCountBySKUStmt,
		getStockByProductSKUAndCountryCodeForUpdateStmt: q.getStockByProductSKUAndCountryCodeForUpdateStmt,
		listProductStocksBySKUStmt:                      q.listProductStocksBySKUStmt,
	}
}
