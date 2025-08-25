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
	"regexp"
	"strconv"
	"strings"
	"time"

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
	// Admin
	app.GET("/all", middleware.MiddlewareLogin, middleware.MiddlewareAdmin, k.GetAllRiwayatKertasKerja)
	app.PATCH("/validasi/:id", middleware.MiddlewareLogin, middleware.MiddlewareAdmin, k.ValidasiKertasKerja)

	// Satker
	app.GET("/", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.GetRiwayatKertasKerjaByUserID)
	app.POST("/:tahap", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.GetDataPembanding)
	app.GET("/lelang/:kode", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.GetDataLelangByKode)
	app.POST("/", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.SaveKertasKerja)
	app.DELETE("/:id", middleware.MiddlewareLogin, middleware.MiddlewareUser, k.DeleteRiwayatKertasKerja)
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

	riwayat, err := k.service.GetRiwayatKertasKerjaByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Data kertas kerja tidak ditemukan"})
		return
	}

	oldPdfPath := riwayat.PdfPath
	oldFileName := filepath.Base(oldPdfPath)

	ext := filepath.Ext(oldFileName) // .pdf
	nameOnly := strings.TrimSuffix(oldFileName, ext)

	// Cari dan ganti bagian datetime di nama file
	// Pola: _dd_mm_yyyy_hh_mm_ss atau _validasi.pdf
	// Contoh: Kertas_Kerja_Yamaha_123456_10_08_2025_19_05_53.pdf
	// atau:   Kertas_Kerja_Yamaha_123456_10_08_2025_19_05_53_validasi.pdf

	// Regex untuk datetime: _dd_mm_yyyy_hh_mm_ss
	re := regexp.MustCompile(`_\d{2}_\d{2}_\d{4}_\d{2}_\d{2}_\d{2}`)
	nowStr := "_" + time.Now().Format("02_01_2006_15_04_05")

	// Hilangkan datetime lama
	nameOnly = re.ReplaceAllString(nameOnly, nowStr)

	// Pastikan ada "_validasi" di akhir nama (sebelum .pdf)
	if !strings.HasSuffix(nameOnly, "_validasi") {
		nameOnly += "_validasi"
	}

	newFileName := nameOnly + ext
	savePath := "assets/kertas_kerja/pdf/" + newFileName

	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	baseURL := os.Getenv("BASE_URL")
	pdfURL := baseURL + "/" + savePath
	resp, err := k.service.ValidasiKertasKerja(id, pdfURL)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
