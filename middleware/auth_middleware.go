package middleware

import (
	"net/http"
	"strings"

	token2 "kertas_kerja/pkg/token"

	"github.com/gin-gonic/gin"
)

func MiddlewareLogin(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	parts := strings.Split(bearerToken, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required and must be Bearer token"})
		ctx.Abort()
		return
	}

	token := parts[1]

	user, err := token2.ValidateAccessToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Set("users", user)

	ctx.Next()
}
