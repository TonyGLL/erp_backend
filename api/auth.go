package api

import (
	"net/http"

	"github.com/TonyGLL/erp_backend/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}

func (server *Server) AuthLogin(ctx *gin.Context) {
	var req AuthLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	passwordHashed, err := server.store.GetUserPassword(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Bad Credentials"})
		return
	}

	token, err := util.CreateToken(req.Username)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	response := AuthLoginResponse{
		Token: token,
	}

	ctx.JSON(http.StatusOK, response)
}