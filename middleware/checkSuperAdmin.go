package middleware

import (
	"fmt"
	"net/http"

	token2 "kertas_kerja/pkg/token"

	"github.com/gin-gonic/gin"
)

func MiddlewareSuperAdmin(ctx *gin.Context) {
	user, exists := ctx.Get("users")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		ctx.Abort()
		return
	}

	fmt.Printf("role user: %v\n", user)

	// Pastikan user struct punya field Role
	u, ok := user.(*token2.UserAuthToken)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid user type"})
		ctx.Abort()
		return
	}

	if u.Role != "superadmin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Super Admin only"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
