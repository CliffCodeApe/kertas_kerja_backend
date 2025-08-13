package handler

import (
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"kertas_kerja/middleware"
	"kertas_kerja/pkg/errs"
	"kertas_kerja/pkg/token"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	app.GET("/all", middleware.MiddlewareLogin, middleware.MiddlewareSuperAdmin, k.GetAllRiwayatKertasKerja)
	app.GET("/", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.GetRiwayatKertasKerjaByUserID)
	app.POST("/:tahap", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.GetDataPembanding)
	app.GET("/lelang/:kode", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.GetDataLelangByKode)
	app.POST("/", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.SaveKertasKerja)
	app.DELETE("/:id", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.DeleteRiwayatKertasKerja)
	app.POST("/validasi/:id", middleware.MiddlewareLogin, middleware.MiddlewareAdmin, k.ValidasiKertasKerja)
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
	var req dto.IdentitasKendaraan
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

func (k *kertasKerjaController) GetAllRiwayatKertasKerja(ctx *gin.Context) {

	result, err := k.service.GetAllRiwayatKertasKerja()
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (k *kertasKerjaController) GetRiwayatKertasKerjaByUserID(ctx *gin.Context) {
	user, valid := ctx.Get("users")

	if !valid {
		ctx.JSON(http.StatusBadRequest, errs.ErrValid)
		return
	}

	userId := user.(*token.UserAuthToken)

	result, err := k.service.GetRiwayatKertasKerjaByUserID(userId.ID)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
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

func (k *kertasKerjaController) SaveKertasKerja(ctx *gin.Context) {
	user, valid := ctx.Get("users")

	if !valid {
		ctx.JSON(http.StatusBadRequest, errs.ErrValid)
		return
	}

	userId := user.(*token.UserAuthToken)

	var payload dto.IsiKertasKerjaRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "detail": err.Error()})
		return
	}

	response, err := k.service.SaveKertasKerja(&payload, userId.ID)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (k *kertasKerjaController) DeleteRiwayatKertasKerja(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "ID harus berupa angka",
			"error":   err.Error(),
		})
		return
	}

	response, err := k.service.DeleteRiwayatKertasKerja(id)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (k *kertasKerjaController) ValidasiKertasKerja(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	file, err := ctx.FormFile("pdf_file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File PDF tidak ditemukan"})
		return
	}

	// Ambil data riwayat untuk dapatkan nama file PDF lama
	riwayat, err := k.service.GetRiwayatKertasKerjaByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Data kertas kerja tidak ditemukan"})
		return
	}

	// Ambil nama file lama
	oldPdfPath := riwayat.PdfPath            // contoh: http://localhost:8080/assets/kertas_kerja/pdf/Kertas_Kerja_Yamaha_123456_10_08_2025_19_05_53.pdf
	oldFileName := filepath.Base(oldPdfPath) // Kertas_Kerja_Yamaha_123456_10_08_2025_19_05_53.pdf

	// Tambahkan "validasi" sebelum .pdf
	ext := filepath.Ext(oldFileName) // .pdf
	nameOnly := strings.TrimSuffix(oldFileName, ext)
	newFileName := nameOnly + "_validasi" + ext // Kertas_Kerja_Yamaha_123456_10_08_2025_19_05_53_validasi.pdf

	savePath := "assets/kertas_kerja/pdf/" + newFileName
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	// Update status validasi di service
	baseURL := os.Getenv("BASE_URL")
	pdfURL := baseURL + "/" + savePath
	resp, err := k.service.ValidasiKertasKerja(id, pdfURL)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
