package api

import (
	"net/http"

	db "github.com/TonyGLL/erp_backend/db/sql"
	"github.com/gin-gonic/gin"
)

type getRolesRequest struct {
	Page  int32 `form:"page" binding:"required,min=1"`
	Limit int32 `form:"limit" binding:"required,min=5,max=10"`
}

type getRolesResponse struct {
	Total int64     `json:"total"`
	Page  int32     `json:"page"`
	Roles []db.Role `json:"roles"`
}

// Get Roles	godoc
// @Summary Get Roles
// @Description Get all roles
// @Tags Roles
// @Accept json
// @Produce application/json
// @Param	request	query	getRolesRequest true "Query Params"
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @Security JWT
// @Success 200 {object} getRolesResponse
// @Failure		401			{string}	gin.H	"StatusUnauthorized"
// @Failure		500			{string}	gin.H	"StatusInternalServerError"
// @Router /roles [get]
func (server *Server) getRoles(ctx *gin.Context) {
	var req getRolesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetRolesParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	roles, err := server.store.GetRoles(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	total, err := server.store.CountRoles(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := getRolesResponse{
		Total: total,
		Page:  req.Page,
		Roles: roles,
	}

	ctx.JSON(http.StatusOK, response)
}
