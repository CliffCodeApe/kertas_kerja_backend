package service

import (
	"fmt"
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"kertas_kerja/entity"
	"kertas_kerja/pkg/errs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type kertasKerjaServ struct {
	kertasKerjarRepo contract.KertasKerjaRepository
}

func implKertasKerjaService(repo *contract.Repository) contract.KertasKerjaService {
	return &kertasKerjaServ{
		kertasKerjarRepo: repo.KertasKerja,
	}
}

func (s *kertasKerjaServ) GetDataPembanding(req *dto.IdentitasKendaraan, tahap int) (*dto.KertasKerjaResponse, error) {
	// Konversi tahun dan kategori lokasi ke tipe data yang sesuai
	kategoriLokasi, _ := strconv.Atoi(req.KategoriLokasi)
	lokasiBersih := strings.TrimPrefix(strings.TrimSpace(req.LokasiObjek), "KPKNL ")

	// Query ke repository
	dataPembanding, err := s.kertasKerjarRepo.FindDataPembanding(
		req.MerekKendaraan,
		req.TipeKendaraan,
		req.TahunPembuatan,
		req.LokasiObjek,
		req.Provinsi,
		tahap,
	)
	if err != nil {
		return &dto.KertasKerjaResponse{
			Status:  "500",
			Message: "Gagal mencari data pembanding",
			Data:    dto.KertasKerjaData{},
		}, err
	}

	// Mapping hasil ke DTO response
	var pembandingList []dto.DataPembanding
	for _, lelang := range dataPembanding {
		pembanding := dto.DataPembanding{
			KodeLelang:     lelang.Kode,
			Merek:          lelang.Merek,
			Tipe:           lelang.Tipe,
			TahunPembuatan: lelang.TahunPembuatan,
			TahunTransaksi: lelang.TahunLelang,
			Lokasi:         lokasiBersih,
			KategoriLokasi: kategoriLokasi,
			HargaLelang:    float64(lelang.HargaLaku),
		}
		pembandingList = append(pembandingList, pembanding)
	}

	result := dto.KertasKerjaData{
		InputLelang: dto.IdentitasKendaraan{
			NamaObjek:           req.NamaObjek,
			LokasiObjek:         lokasiBersih,
			NUP:                 req.NUP,
			KategoriLokasi:      req.KategoriLokasi,
			MerekKendaraan:      req.MerekKendaraan,
			TipeKendaraan:       req.TipeKendaraan,
			TahunPembuatan:      req.TahunPembuatan,
			NomorPolisi:         req.NomorPolisi,
			DokumenKepemilikan:  req.DokumenKepemilikan,
			PemilikDokumen:      req.PemilikDokumen,
			JenisKendaraan:      req.JenisKendaraan,
			MasaBerlaku:         req.MasaBerlaku,
			PenggunaanKendaraan: req.PenggunaanKendaraan,
			Keterangan:          req.Keterangan,
			Warna:               req.Warna,
			BahanBakar:          req.BahanBakar,
			KondisiKendaraan:    req.KondisiKendaraan,
			Provinsi:            req.Provinsi,
		},
		DataPembanding: pembandingList,
	}

	return &dto.KertasKerjaResponse{
		Status:  "200",
		Message: "Berhasil mendapatkan data pembanding",
		Data:    result,
	}, nil
}

func (s *kertasKerjaServ) GetDataLelangByKode(kode string) (*dto.DataPembandingResponse, error) {
	lelang, err := s.kertasKerjarRepo.FindDataLelangByKode(kode)
	if err != nil {
		return nil, errs.ErrDataLelangNotFound
	}

	kategoriLokasi := GetKategoriLokasi(lelang.Kpknl)
	lokasiBersih := strings.TrimPrefix(strings.TrimSpace(lelang.Kpknl), "KPKNL ")

	return &dto.DataPembandingResponse{
		Status:  "200",
		Message: "Berhasil mendapatkan data lelang",
		Data: dto.DataPembanding{
			KodeLelang:     lelang.Kode,
			Merek:          lelang.Merek,
			Tipe:           lelang.Tipe,
			TahunPembuatan: lelang.TahunPembuatan,
			TahunTransaksi: lelang.TahunLelang,
			Lokasi:         lokasiBersih,
			KategoriLokasi: kategoriLokasi,
			HargaLelang:    float64(lelang.HargaLaku),
		},
	}, nil
}

func (s *kertasKerjaServ) GetAllRiwayatKertasKerja() (*dto.GetAllRiwayatKertasKerjaResponse, error) {
	riwayat, err := s.kertasKerjarRepo.GetAllRiwayatKertasKerja()
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan riwayat kertas kerja: %w", err)
	}

	var result []dto.RiwayatKertasKerjaData
	for _, item := range riwayat {
		result = append(result, dto.RiwayatKertasKerjaData{
			ID:        item.ID,
			UserID:    item.UserID,
			NamaObjek: item.NamaObjek,
			PdfPath:   item.PdfPath,
			KodeKL:    item.KodeKL,
		})
	}

	response := &dto.GetAllRiwayatKertasKerjaResponse{
		Status:  "200",
		Message: "Berhasil mendapatkan riwayat kertas kerja",
		Data:    result,
	}

	return response, nil
}

func (s *kertasKerjaServ) GetRiwayatKertasKerjaByUserID(userID uint64) (*dto.GetAllRiwayatKertasKerjaResponse, error) {
	riwayat, err := s.kertasKerjarRepo.GetRiwayatKertasKerjaByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan riwayat kertas kerja: %w", err)
	}

	var result []dto.RiwayatKertasKerjaData
	for _, item := range riwayat {
		result = append(result, dto.RiwayatKertasKerjaData{
			ID:        item.ID,
			UserID:    item.UserID,
			NamaObjek: item.NamaObjek,
			PdfPath:   item.PdfPath,
			KodeKL:    item.KodeKL,
		})
	}

	response := &dto.GetAllRiwayatKertasKerjaResponse{
		Status:  "200",
		Message: "Berhasil mendapatkan riwayat kertas kerja",
		Data:    result,
	}

	return response, nil
}

func (s *kertasKerjaServ) GetRiwayatKertasKerjaByID(id uint64) (*dto.RiwayatKertasKerjaData, error) {
	riwayat, err := s.kertasKerjarRepo.GetRiwayatKertasKerjaByID(id)
	if err != nil {
		return nil, fmt.Errorf("data riwayat tidak ditemukan: %w", err)
	}

	result := &dto.RiwayatKertasKerjaData{
		ID:         riwayat.ID,
		UserID:     riwayat.UserID,
		NamaObjek:  riwayat.NamaObjek,
		PdfPath:    riwayat.PdfPath,
		ExcelPath:  riwayat.ExcelPath,
		IsVerified: riwayat.IsVerified,
		KodeKL:     riwayat.KodeKL,
	}

	return result, nil
}

func (s *kertasKerjaServ) DeleteRiwayatKertasKerja(id uint64) (*dto.DeleteRiwayatKertasKerjaResponse, error) {
	riwayat, err := s.kertasKerjarRepo.GetRiwayatKertasKerjaByID(id)
	if err != nil {
		return nil, fmt.Errorf("data riwayat tidak ditemukan: %w", err)
	}

	// 2. Hapus file PDF dan Excel (jika ada)
	removeFileFromURL(riwayat.PdfPath)
	removeFileFromURL(riwayat.ExcelPath)

	// 3. Hapus data dari database
	err = s.kertasKerjarRepo.DeleteRiwayatKertasKerja(id)
	if err != nil {
		return nil, fmt.Errorf("gagal menghapus riwayat kertas kerja: %w", err)
	}

	return &dto.DeleteRiwayatKertasKerjaResponse{
		Status:  "200",
		Message: "Berhasil menghapus riwayat kertas kerja",
	}, nil
}

// Helper untuk menghapus file dari path URL
func removeFileFromURL(fileURL string) {
	// Hilangkan BASE_URL jika ada
	baseURL := os.Getenv("BASE_URL")
	filePath := strings.TrimPrefix(fileURL, baseURL)
	filePath = strings.TrimPrefix(filePath, "/") // pastikan tidak ada slash di depan
	_ = os.Remove(filePath)                      // abaikan error jika file tidak ada
}

func (s *kertasKerjaServ) SaveKertasKerja(payload *dto.IsiKertasKerjaRequest, userID uint64) (*dto.RiwayatKertasKerjaResponse, error) {
	// Simpan ke Excel
	excelPath, err := IsiDataLelangKeExcel(payload)
	if err != nil {
		return nil, fmt.Errorf("gagal simpan excel: %w", err)
	}

	// Convert ke PDF
	pdfPath, err := ConvertExcelToPDF(excelPath)
	if err != nil {
		return nil, fmt.Errorf("gagal konversi ke PDF: %w", err)
	}

	// Simpan ke DB
	riwayat := &entity.KertasKerja{
		UserID:    userID,
		NamaObjek: payload.InputLelang.NamaObjek,
		PdfPath:   pdfPath,
		ExcelPath: excelPath,
	}

	err = s.kertasKerjarRepo.InsertRiwayatKertasKerja(riwayat)
	if err != nil {
		return nil, err
	}

	response := &dto.RiwayatKertasKerjaResponse{
		Status:  "200",
		Message: "Berhasil menyimpan riwayat kertas kerja",
		Data: dto.RiwayatKertasKerjaData{
			UserID:     riwayat.UserID,
			NamaObjek:  riwayat.NamaObjek,
			PdfPath:    riwayat.PdfPath,
			ExcelPath:  riwayat.ExcelPath,
			IsVerified: riwayat.IsVerified,
		},
	}

	return response, nil
}

func (s *kertasKerjaServ) ValidasiKertasKerja(id uint64, pdfPath string) (*dto.ValidasiKertasKerjaResponse, error) {
	err := s.kertasKerjarRepo.ValidasiKertasKerja(id, pdfPath)
	if err != nil {
		return nil, fmt.Errorf("gagal validasi kertas kerja: %w", err)
	}
	return &dto.ValidasiKertasKerjaResponse{
		Status:  "200",
		Message: "Kertas kerja berhasil divalidasi",
	}, nil
}

func IsiDataLelangKeExcel(data *dto.IsiKertasKerjaRequest) (savePath string, err error) {
	f, err := excelize.OpenFile("assets/template/Kertas_Kerja_template.xlsx")
	if err != nil {
		return "", err
	}

	sheet := "kertas_kerja" // misal: Sheet1

	// Identitas Kendaraan
	f.SetCellValue(sheet, "E11", data.InputLelang.NamaObjek)
	f.SetCellValue(sheet, "K11", data.InputLelang.NUP)

	f.SetCellValue(sheet, "E12", data.InputLelang.LokasiObjek)
	f.SetCellValue(sheet, "K12", data.InputLelang.KategoriLokasi)

	switch data.InputLelang.JenisKendaraan {
	case "Roda 4 atau lebih":
		f.SetCellValue(sheet, "I13", "✓")
	case "Roda 2":
		f.SetCellValue(sheet, "E13", "✓")
	}

	f.SetCellValue(sheet, "E14", data.InputLelang.MerekKendaraan)
	f.SetCellValue(sheet, "E15", data.InputLelang.TipeKendaraan)

	f.SetCellValue(sheet, "E17", data.InputLelang.NomorPolisi)

	// DokumenKepemilikan sekarang array, centang semua yang sesuai
	for _, dok := range data.InputLelang.DokumenKepemilikan {
		switch dok {
		case "BPKB":
			f.SetCellValue(sheet, "E18", "✓")
		case "STNK":
			f.SetCellValue(sheet, "I18", "✓")
		case "Lainnya":
			f.SetCellValue(sheet, "L18", "✓")
		case "Tidak ada":
			f.SetCellValue(sheet, "P18", "✓")
		}
	}

	f.SetCellValue(sheet, "E19", data.InputLelang.PemilikDokumen)

	switch data.InputLelang.MasaBerlaku {
	case "Masih Berlaku":
		f.SetCellValue(sheet, "E20", "✓")
	case "Habis Masa Berlaku":
		f.SetCellValue(sheet, "I20", "✓")
	}

	f.SetCellValue(sheet, "E22", data.InputLelang.PenggunaanKendaraan)
	f.SetCellValue(sheet, "E23", data.InputLelang.Keterangan)

	f.SetCellValue(sheet, "E26", data.InputLelang.Warna)
	f.SetCellValue(sheet, "E27", data.InputLelang.TahunPembuatan)
	f.SetCellValue(sheet, "E28", data.InputLelang.BahanBakar)

	switch data.InputLelang.KondisiKendaraan {
	case "0.5":
		f.SetCellValue(sheet, "M26", "✓")
	case "0.6":
		f.SetCellValue(sheet, "M27", "✓")
	case "0.7":
		f.SetCellValue(sheet, "M28", "✓")
	}

	// Isi data pembanding
	if len(data.DataPembanding) > 0 {
		for i, pembanding := range data.DataPembanding {
			row := 31 + i // Mulai dari baris 32 untuk data pembanding
			f.SetCellValue(sheet, "B"+strconv.Itoa(row), pembanding.Merek)
			f.SetCellValue(sheet, "D"+strconv.Itoa(row), pembanding.KodeLelang)
			f.SetCellValue(sheet, "G"+strconv.Itoa(row), pembanding.Tipe)
			f.SetCellValue(sheet, "I"+strconv.Itoa(row), pembanding.HargaLelang)
			f.SetCellValue(sheet, "K"+strconv.Itoa(row), pembanding.TahunTransaksi)
			f.SetCellValue(sheet, "N"+strconv.Itoa(row), pembanding.Lokasi)
			f.SetCellValue(sheet, "O"+strconv.Itoa(row), pembanding.KategoriLokasi)
			f.SetCellValue(sheet, "P"+strconv.Itoa(row), pembanding.TahunPembuatan)
		}
	}

	// Data Penyesuaian
	if len(data.DataPenyesuaian) > 0 {
		for i, penyesuaian := range data.DataPenyesuaian {
			row := 39 + i // Mulai dari baris setelah data pembanding
			f.SetCellValue(sheet, "C"+strconv.Itoa(row), penyesuaian.DataHasilLelang)
			f.SetCellValue(sheet, "E"+strconv.Itoa(row), penyesuaian.Tipe)
			f.SetCellValue(sheet, "G"+strconv.Itoa(row), penyesuaian.Merek)
			f.SetCellValue(sheet, "H"+strconv.Itoa(row), penyesuaian.Waktu)
			f.SetCellValue(sheet, "J"+strconv.Itoa(row), penyesuaian.Lokasi)
			f.SetCellValue(sheet, "K"+strconv.Itoa(row), penyesuaian.TahunPembuatan)
			f.SetCellValue(sheet, "N"+strconv.Itoa(row), penyesuaian.Total)
			f.SetCellValue(sheet, "P"+strconv.Itoa(row), penyesuaian.NilaiTaksiran)
		}
	}

	f.SetCellValue(sheet, "N49", data.FaktorKondisi)

	// Hasil Taksiran

	if data.HasilTaksiran.TotalNilaiTaksiran != 0 || data.HasilTaksiran.RataRataNilaiTaksiran != 0 ||
		data.HasilTaksiran.TaksiranNilaiLimitLelang != 0 || data.HasilTaksiran.Pembulatan != 0 {
		f.SetCellValue(sheet, "P"+strconv.Itoa(47), data.HasilTaksiran.TotalNilaiTaksiran)
		f.SetCellValue(sheet, "P"+strconv.Itoa(48), data.HasilTaksiran.RataRataNilaiTaksiran)
		f.SetCellValue(sheet, "P"+strconv.Itoa(49), data.HasilTaksiran.TaksiranNilaiLimitLelang)
		f.SetCellValue(sheet, "P"+strconv.Itoa(50), data.HasilTaksiran.Pembulatan)
	}

	baseURL := os.Getenv("BASE_URL")
	savePath = "assets/kertas_kerja/Kertas_Kerja_" + data.InputLelang.NamaObjek + "_" + data.InputLelang.NUP + "_" + time.Now().Format("02_01_2006_15_04_05") + ".xlsx"
	excelPath := baseURL + "/" + savePath
	// Simpan file hasil
	f.SaveAs(savePath)

	return excelPath, err
}

func ConvertExcelToPDF(excelPath string) (pdfPath string, err error) {

	pdfDir := "assets/kertas_kerja/pdf"
	if err := os.MkdirAll(pdfDir, 0755); err != nil {
		return "", err
	}

	done := make(chan error, 1)
	go func() {
		cmd := exec.Command("libreoffice", "--headless", "--convert-to", "pdf", excelPath, "--outdir", pdfDir)
		done <- cmd.Run()
	}()

	err = <-done
	if err != nil {
		return "", err
	}

	baseURL := os.Getenv("BASE_URL")
	pdfPath = filepath.Join(pdfDir, strings.ReplaceAll(filepath.Base(excelPath), ".xlsx", ".pdf"))
	pdfPath = baseURL + "/" + pdfPath
	return pdfPath, nil
}

// Daftar lokasi per kategori
var kategori1 = []string{
	"Malang", "Medan", "Solo", "Semarang", "Yogyakarta", "Pelembang", "Denpasar",
	"Pekanbaru", "Makassar", "Cirebon", "Balikpapan", "Batam",
}

var kategori2 = []string{
	"Bekasi", "Tanggerang", "Bogor", "Jakarta", "Sidoarjo", "Surabaya", "Bandung", "Serang",
}

var kategori3 = []string{
	"Purwakarta", "Bandar Lampung", "Banjarmasin", "Tasikmalaya", "Madiun", "Purwokerto", "Padang",
	"Banda Aceh", "Tegal", "Pekalongan", "Jember", "Singaraja", "Bukittinggi", "Samarinda", "Metro",
	"Mataram", "Jambi", "Pematang Siantar", "Pontianak", "Kisaran", "Ljhoksumawe", "Dumai", "Bengkulu",
	"Pemekasan", "Manado", "Lahat", "Palangkaraya", "Bontang", "Pare-Pare", "Kendari", "Kupang", "Ambon",
	"Padang Sideumpuan", "Gorontalo", "Palu", "Jayapura", "Singkawang", "Palopo", "Sorong",
	"Pangkal Pinang", "Mamuju", "Tarakan", "Bima", "Pangkalan Bun", "Ternate", "Biak",
}

func GetKategoriLokasi(rawLokasi string) int {
	lokasi := strings.TrimPrefix(strings.TrimSpace(rawLokasi), "KPKNL ")
	if strings.HasPrefix(lokasi, "Jakarta") {
		lokasi = "Jakarta"
	}
	for _, loc := range kategori1 {
		if strings.EqualFold(loc, lokasi) {
			return 1
		}
	}
	for _, loc := range kategori2 {
		if strings.EqualFold(loc, lokasi) {
			return 2
		}
	}
	for _, loc := range kategori3 {
		if strings.EqualFold(loc, lokasi) {
			return 3
		}
	}
	return 0 // Tidak terklasifikasi
}
