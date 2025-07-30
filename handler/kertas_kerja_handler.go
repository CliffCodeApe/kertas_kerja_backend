package handler

import (
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kertasKerjaController struct {
	service contract.KertasKerjaService
}

func (k *kertasKerjaController) getPrefix() string {
	return "/kertas-kerja"
}

func (k *kertasKerjaController) initService(service *contract.Service) {
	k.service = service.KertasKerja
}

func (k *kertasKerjaController) initRoute(app *gin.RouterGroup) {
	app.POST("/", k.GetDataPembanding)
}

func (k *kertasKerjaController) GetDataPembanding(ctx *gin.Context) {
	var req dto.KertasKerjaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	response, err := k.service.GetDataPembanding(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Gagal mencari data pembanding",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
