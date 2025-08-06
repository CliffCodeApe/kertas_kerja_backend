package handler

import (
	"kertas_kerja/contract"
	"kertas_kerja/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service contract.UserService
}

func (a *userController) getPrefix() string {
	return "/user"
}

func (a *userController) initService(service *contract.Service) {
	a.service = service.User
}

func (a *userController) initRoute(app *gin.RouterGroup) {
	app.POST("/validate/:userID", middleware.MiddlewareLogin, middleware.MiddlewareSuperAdmin, a.ValidateUser)
	app.GET("/all", middleware.MiddlewareLogin, middleware.MiddlewareSuperAdmin, a.GetAllUsers)
}

func (a *userController) ValidateUser(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	result, err := a.service.ValidateUser(userID)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(result.StatusCode, result)

}

func (a *userController) GetAllUsers(ctx *gin.Context) {
	result, err := a.service.GetAllUsers()
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
