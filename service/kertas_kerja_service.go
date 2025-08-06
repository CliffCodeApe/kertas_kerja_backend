package service

import (
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"os/exec"
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

func (s *kertasKerjaServ) GetDataPembanding(req *dto.KertasKerjaRequest, tahap int) (*dto.KertasKerjaResponse, error) {
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
		InputLelang: dto.KertasKerjaRequest{
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

func (s *kertasKerjaServ) SaveKertasKerjaToExcel(input *dto.KertasKerjaRequest, pembandingList *[]dto.DataPembanding) error {
	return IsiDataInputKeExcel(input, pembandingList)
}

func (s *kertasKerjaServ) GetDataLelangByKode(kode string) (*dto.DataPembandingResponse, error) {
	lelang, err := s.kertasKerjarRepo.FindDataLelangByKode(kode)
	if err != nil {
		return nil, err
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

// func (s *kertasKerjaServ) InsertRiwayatKertasKerja(payload *dto.RiwayatKertasKerjaRequest) (*dto.RiwayatKertasKerjaResponse, error) {

// }

func IsiDataInputKeExcel(input *dto.KertasKerjaRequest, dataPembanding *[]dto.DataPembanding) (err error) {
	f, err := excelize.OpenFile("assets/template/Kertas_Kerja_template.xlsx")
	if err != nil {
		return err
	}

	sheet := "Simulasi Di Jawa" // misal: Sheet1
	// fmt.Printf("Sheet: %s\n", sheet)

	// Isian berdasarkan posisi sel (sesuai contoh Excel di atas)
	f.SetCellValue(sheet, "E11", input.NamaObjek)
	f.SetCellValue(sheet, "K11", input.NUP)

	f.SetCellValue(sheet, "E12", input.LokasiObjek)
	f.SetCellValue(sheet, "K12", input.KategoriLokasi)

	switch input.JenisKendaraan {
	case "Roda 4 atau lebih":
		f.SetCellValue(sheet, "I13", "✓")
	case "Roda 2":
		f.SetCellValue(sheet, "E13", "✓")
	}

	f.SetCellValue(sheet, "E14", input.MerekKendaraan)
	f.SetCellValue(sheet, "E15", input.TipeKendaraan)

	f.SetCellValue(sheet, "E17", input.NomorPolisi)

	switch input.DokumenKepemilikan {
	case "BPKB":
		f.SetCellValue(sheet, "E18", "✓")
	case "STNK":
		f.SetCellValue(sheet, "I18", "✓")
	case "Lainnya":
		f.SetCellValue(sheet, "L18", "✓")
	case "Tidak ada":
		f.SetCellValue(sheet, "P18", "✓")
	}

	f.SetCellValue(sheet, "E19", input.PemilikDokumen)

	switch input.MasaBerlaku {
	case "Masih Berlaku":
		f.SetCellValue(sheet, "E20", "✓")
	case "Habis Masa Berlaku":
		f.SetCellValue(sheet, "I20", "✓")
	}

	f.SetCellValue(sheet, "E22", input.PenggunaanKendaraan)
	f.SetCellValue(sheet, "E23", input.Keterangan)

	f.SetCellValue(sheet, "E26", input.Warna)
	f.SetCellValue(sheet, "E27", input.TahunPembuatan)
	f.SetCellValue(sheet, "E28", input.BahanBakar)

	// Centang kondisi kendaraan berdasarkan nilai
	switch input.KondisiKendaraan {
	case "0.5":
		f.SetCellValue(sheet, "M26", "✓")
	case "0.6":
		f.SetCellValue(sheet, "M27", "✓")
	case "0.7":
		f.SetCellValue(sheet, "M28", "✓")
	}

	// Isi data pembanding
	if len(*dataPembanding) > 0 {
		for i, pembanding := range *dataPembanding {
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

	savePath := "assets/kertas_kerja/Kertas_Kerja_" + input.NamaObjek + "_" + input.NUP + "_" + time.Now().Format("02_01_2006_15_04_05") + ".xlsx"
	// Simpan file hasil
	f.SaveAs(savePath)

	excelFilename := "Kertas_Kerja_" + time.Now().Format("02_01_2006") + ".xlsx"
	excelPath := "assets/kertas_kerja/" + excelFilename
	pdfDir := "assets/kertas_kerja/"
	err = ConvertExcelToPDF(excelPath, pdfDir)
	if err != nil {
		return err
	}

	return nil
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

func ConvertExcelToPDF(excelPath, pdfDir string) error {
	cmd := exec.Command("libreoffice", "--headless", "--convert-to", "pdf", excelPath, "--outdir", pdfDir)
	return cmd.Run()
}
