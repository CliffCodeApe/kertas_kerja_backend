package handler

import (
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"kertas_kerja/pkg/errs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authController struct {
	service contract.AuthService
}

func (a *authController) getPrefix() string {
	return "/auth"
}

func (a *authController) initService(service *contract.Service) {
	a.service = service.Auth
}

func (a *authController) initRoute(app *gin.RouterGroup) {
	app.POST("/login", a.Login)
	app.GET("/refreshToken", a.RefreshToken)
}

func (a *authController) Login(ctx *gin.Context) {
	var payload dto.AuthLoginRequest
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, errs.ErrRequestBody)
		return
	}

	result, err := a.service.Login(&payload)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (a *authController) RefreshToken(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	parts := strings.Split(bearerToken, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, errs.NewUnauthorizedError("Authorization token is required and must be Bearer token"))
		return
	}

	token := parts[1]

	result, err := a.service.RefreshToken(token)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(result.StatusCode, result)

}
