package middlewares

import (
	"net/http"
	"strings"

	"github.com/TonyGLL/erp_backend/util"
	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	token := strings.Split(authorization, " ")[1]
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	_, err := util.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
