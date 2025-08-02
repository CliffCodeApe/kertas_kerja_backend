package middleware

import (
	"net/http"

	token2 "kertas_kerja/pkg/token"

	"github.com/gin-gonic/gin"
)

func MiddlewareAdmin(ctx *gin.Context) {
	user, exists := ctx.Get("users")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		ctx.Abort()
		return
	}

	// Pastikan user struct punya field Role
	u, ok := user.(*token2.UserAuthToken)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid user type"})
		ctx.Abort()
		return
	}

	if u.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Access denied: admin only"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
