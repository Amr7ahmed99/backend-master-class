package util

import (
	"backend-master-class/db/connection"
	db "backend-master-class/db/sqlc"
	"context"
)

func IsSupportedCurrency(currency int32) bool {
	store := db.NewStore(connection.ConnectionDB)
	if _, err := store.GetCurrency(context.Background(), int64(currency)); err != nil {
		return false
	}
	return true
}
