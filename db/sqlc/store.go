package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions and execute db queries and transaction
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (result TransferTxResult, err error)
}

// Store provides all functions and execute SQL queries and transaction
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

// execTX executes a function within a db transaction
func (store *SQLStore) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other.
// It creates a transfer record, adds account entries and updates accounts' balance within a single db transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (result TransferTxResult, err error) {
	err = store.execTX(ctx, func(q *Queries) error {
		//Create Transfer
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		//Create From-Entry
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		//Create To-Entry
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountId < arg.ToAccountId {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountId, -arg.Amount, arg.ToAccountId, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountId, arg.Amount, arg.FromAccountId, -arg.Amount)
		}

		if err != nil {
			return err
		}
		return nil
	})

	return
}

func addMoney(ctx context.Context, q *Queries, fromAccountId int64, fromAmount int64, toAccountId int64, toAmount int64) (fromAccount Account, toAccount Account, err error) {
	//Update From-Account balance
	fromAccount, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
		ID:     fromAccountId,
		Amount: fromAmount,
	})
	if err != nil {
		return
	}

	//Update To-Account balance
	toAccount, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
		ID:     toAccountId,
		Amount: toAmount,
	})

	return
}
