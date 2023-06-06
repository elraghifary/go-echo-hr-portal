// Package sql implements some sql utility functions
package sql

import (
	"context"
	"database/sql"

	"github.com/elraghifary/go-echo-hr-portal/cmd/identifier"
)

// Tx represents the interface to execute multiple database operation
type Tx struct {
	db  *sql.DB
	key identifier.TxKey
}

// NewTx creates a new MySQL transaction
func NewTx(db *sql.DB, key identifier.TxKey) *Tx {
	return &Tx{
		db:  db,
		key: key,
	}
}

// ExecTx runs the fn with the transaction inside the context
func (t *Tx) ExecTx(ctx context.Context, fn func(fnctx context.Context) error) error {
	tx, err := t.db.Begin()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}

		return err
	}

	ctx = context.WithValue(ctx, t.key, tx)

	err = fn(ctx)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Returns key to access transaction from ctx
func (t *Tx) GetTxKey() identifier.TxKey {
	return t.key
}
