package handler

import (
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"kertas_kerja/middleware"
	"net/http"
	"strconv"

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
	app.POST("/:tahap", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.GetDataPembanding)
	app.GET("/lelang/:kode", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.GetDataLelangByKode)
	app.POST("/save", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.SaveKertasKerjaToExcel)
	app.POST("/", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.InsertRiwayatKertasKerja)
}

func (k *kertasKerjaController) GetDataPembanding(ctx *gin.Context) {
	tahapStr := ctx.Param("tahap")
	tahap, err := strconv.Atoi(tahapStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Tahap harus berupa angka",
			"error":   err.Error(),
		})
		return
	}
	var req dto.KertasKerjaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	response, err := k.service.GetDataPembanding(&req, tahap)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (k *kertasKerjaController) GetDataLelangByKode(ctx *gin.Context) {
	kode := ctx.Param("kode")
	if kode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Kode lelang tidak boleh kosong",
		})
		return
	}

	response, err := k.service.GetDataLelangByKode(kode)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (k *kertasKerjaController) SaveKertasKerjaToExcel(ctx *gin.Context) {
	var req dto.KertasKerjaData // berisi input lelang + list data pembanding
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	response, err := k.service.SaveKertasKerja(&req.InputLelang, &req.DataPembanding)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (k *kertasKerjaController) InsertRiwayatKertasKerja(ctx *gin.Context) {
	var req dto.RiwayatKertasKerjaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	response, err := k.service.InsertRiwayatKertasKerja(&req)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
