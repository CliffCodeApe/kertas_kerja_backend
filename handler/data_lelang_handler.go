package handler

import (
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"kertas_kerja/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dataLelangController struct {
	service contract.DataLelangService
}

func (k *dataLelangController) getPrefix() string {
	return "/data-lelang"
}

func (k *dataLelangController) initService(service *contract.Service) {
	k.service = service.DataLelang
}

func (k *dataLelangController) initRoute(app *gin.RouterGroup) {
	app.POST("/tambah-data", middleware.MiddlewareLogin, middleware.MiddlewareSuperAdmin, k.InsertDataLelang)
}

func (k *dataLelangController) InsertDataLelang(ctx *gin.Context) {
	var req dto.DataLelang
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	response, err := k.service.InsertDataLelang(&req)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
