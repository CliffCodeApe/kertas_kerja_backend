package handler

import (
	"kertas_kerja/contract"
	"kertas_kerja/dto"
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
	// Super
	app.PATCH("/validate/:userID", middleware.MiddlewareLogin, middleware.MiddlewareSuperAdmin, a.ValidateUser)
	app.PATCH("/invalidate/:userID", middleware.MiddlewareLogin, middleware.MiddlewareSuperAdmin, a.InValidateUser)
	app.PATCH("/change-role/:userID", middleware.MiddlewareLogin, middleware.MiddlewareSuperAdmin, a.ChangeUserRole)
	app.GET("/all", middleware.MiddlewareLogin, middleware.MiddlewareSuperAdmin, a.GetAllUsers)
	app.DELETE("/:userID", middleware.MiddlewareLogin, middleware.MiddlewareSuperAdmin, a.DeleteUser)
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

func (a *userController) InValidateUser(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	result, err := a.service.InValidateUser(userID)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(result.StatusCode, result)

}

func (a *userController) DeleteUser(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	result, err := a.service.DeleteUser(userID)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(result.StatusCode, result)

}

func (a *userController) ChangeUserRole(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	var payload dto.ChangeUserRoleRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		handlerError(ctx, err)
		return
	}

	result, err := a.service.ChangeUserRole(&payload, userID)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(result.StatusCode, result)

}
