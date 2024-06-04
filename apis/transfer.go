package apis

import (
	. "backend-master-class/apis/requests"
	db "backend-master-class/db/sqlc"
	"backend-master-class/token"
	"database/sql"
	"errors"
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

	authPayload := ctx.MustGet(AUTHORIZATION_PAYLOAD).(*token.Payload)

	fromAccount, misMatch := server.isValidAccount(ctx, req.FromAccount, req.Currency)
	if misMatch {
		return
	}

	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, server.errorResponse(err))
		return
	}

	_, misMatch = server.isValidAccount(ctx, req.ToAccount, req.Currency)
	if misMatch {
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

func (server *Server) isValidAccount(ctx *gin.Context, accountId int64, currency int32) (db.Account, bool) {
	account, err := server.Store.GetAccount(ctx, accountId)
	fmt.Println(currency, accountId)
	if err != nil {
		status := http.StatusBadRequest
		if err == sql.ErrNoRows {
			status = http.StatusNotFound
		}
		ctx.JSON(status, server.errorResponse(err))
		return account, false
	}
	if account.CurrencyID != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %d vs %d", accountId, account.CurrencyID, currency)
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return account, false
	}

	return account, true
}
