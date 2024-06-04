package apis

import (
	. "backend-master-class/apis/requests"
	db "backend-master-class/db/sqlc"
	"backend-master-class/enums"
	"backend-master-class/token"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

const (
	AUTHORIZATION_PAYLOAD = "authorization_payload"
)

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AUTHORIZATION_PAYLOAD).(*token.Payload)
	account, err := server.Store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:      authPayload.Username,
		Balance:    0,
		CurrencyID: req.Currency,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case enums.FOREIGN_KEY_VIOLATION, enums.UNIQUE_VIOLATION:
				ctx.JSON(http.StatusForbidden, server.errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) getAccount(ctx *gin.Context) {

	var getAccountReq GetAccountRequest

	if err := ctx.ShouldBindUri(&getAccountReq); err != nil {
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}

	account, err := server.Store.GetAccount(ctx, getAccountReq.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("account doesn't exist")
			ctx.JSON(http.StatusNotFound, server.errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AUTHORIZATION_PAYLOAD).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, server.errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req ListAccountRequest
	var err error

	err = ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AUTHORIZATION_PAYLOAD).(*token.Payload)
	accounts, err := server.Store.ListAccounts(ctx, db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var req GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}

	var updateReq UpdateAccountRequest
	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AUTHORIZATION_PAYLOAD).(*token.Payload)
	account, err := server.Store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("account doesn't exist")
			ctx.JSON(http.StatusNotFound, server.errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}

	if account.Owner != authPayload.Username {
		err = errors.New("account doesn't belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, server.errorResponse(err))
		return
	}

	updatedAccount, err := server.Store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID:         req.ID,
		Balance:    updateReq.Balance,
		CurrencyID: updateReq.Currency,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedAccount)
}
