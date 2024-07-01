package api

import (
	"database/sql"
	"net/http"

	db "github.com/TonyGLL/erp_backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

/* GET USER */
type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

/* GET USERS */
type getUsersRequest struct {
	Page  int32 `form:"page" binding:"required,min=1"`
	Limit int32 `form:"limit" binding:"required,min=5,max=10"`
}

type getUsersResponse struct {
	Total int64           `json:"total"`
	Page  int32           `json:"page"`
	Users []db.GetUserRow `json:"users"`
}

func (server *Server) getUsers(ctx *gin.Context) {
	var req getUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetUsersParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	users, err := server.store.GetUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	total, err := server.store.CountUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := getUsersResponse{
		Total: total,
		Page:  req.Page,
		Users: users,
	}

	ctx.JSON(http.StatusOK, response)
}

type TestResponse struct {
	Ok bool `json:"ok"`
}

/* CREATE USER */
func (server *Server) createUser(ctx *gin.Context) {
	response := TestResponse{
		Ok: true,
	}
	ctx.JSON(http.StatusCreated, response)
}

/* UPDATE USER */
func (server *Server) updateUser(ctx *gin.Context) {
	response := TestResponse{
		Ok: true,
	}
	ctx.JSON(http.StatusNoContent, response)
}

/* DELETE USER */
func (server *Server) deleteUser(ctx *gin.Context) {
	response := TestResponse{
		Ok: true,
	}
	ctx.JSON(http.StatusNoContent, response)
}
