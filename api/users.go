package api

import (
	. "backend-master-class/api/request_params"
	db "backend-master-class/db/sqlc"
	"backend-master-class/enums"
	"backend-master-class/util"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}

	user, err := server.Store.CreateUser(ctx, db.CreateUserParams{
		Username:       req.Username,
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
		Email:          req.Email,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case enums.UNIQUE_VIOLATION:
				ctx.JSON(http.StatusForbidden, server.errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) getUser(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}
	user, err := server.Store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("user doesn't exist")
			ctx.JSON(http.StatusNotFound, server.errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req ListUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}

	users, err := server.Store.ListUser(ctx, db.ListUserParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
