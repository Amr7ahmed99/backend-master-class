package api

import (
	. "backend-master-class/api/request_params"
	db "backend-master-class/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createTransfer(ctx *gin.Context) {
	var req CreateTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}

	if match := server.isValidAccount(ctx, req.FromAccount, req.Currency); !match {
		return
	}

	if match := server.isValidAccount(ctx, req.ToAccount, req.Currency); !match {
		return
	}

	result, err := server.Store.TransferTx(ctx, db.TransferTxParams{
		FromAccountId: req.FromAccount,
		ToAccountId:   req.ToAccount,
		Amount:        req.Amount,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) isValidAccount(ctx *gin.Context, accountId int64, currency int32) bool {
	account, err := server.Store.GetAccount(ctx, accountId)
	fmt.Println(currency, accountId)
	if err != nil {
		status := http.StatusBadRequest
		if err == sql.ErrNoRows {
			status = http.StatusNotFound
		}
		ctx.JSON(status, server.errorResponse(err))
		return false
	}
	if account.CurrencyID != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %d vs %d", accountId, account.CurrencyID, currency)
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return false
	}

	return true
}
