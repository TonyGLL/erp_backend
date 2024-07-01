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

	header := strings.Split(authorization, " ")
	if len(header) != 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	token := header[1]

	_, err := util.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
