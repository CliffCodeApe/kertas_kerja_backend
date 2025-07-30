package service

import (
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"strconv"
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

func (s *kertasKerjaServ) GetDataPembanding(req *dto.KertasKerjaRequest) (*dto.KertasKerjaResponse, error) {
	// Konversi tahun dan kategori lokasi ke tipe data yang sesuai
	tahunPembuatan, _ := strconv.Atoi(req.TahunPembuatan)
	kategoriLokasi, _ := strconv.Atoi(req.KategoriLokasi)

	// Query ke repository
	dataPembanding, err := s.kertasKerjarRepo.FindDataPembanding(
		req.MerekKendaraan,
		req.TipeKendaraan,
		tahunPembuatan,
		req.LokasiObjek,
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
			Lokasi:         lelang.Kpknl,
			KategoriLokasi: kategoriLokasi,
			HargaLelang:    float64(lelang.HargaLaku),
		}
		pembandingList = append(pembandingList, pembanding)
	}

	result := dto.KertasKerjaData{
		InputLelang:    *req,
		DataPembanding: pembandingList,
	}

	err = IsiDataInputKeExcel(&result.InputLelang)
	if err != nil {
		return &dto.KertasKerjaResponse{
			Status:  "500",
			Message: "Gagal mengisi data ke Excel",
			Data:    dto.KertasKerjaData{},
		}, err
	}

	return &dto.KertasKerjaResponse{
		Status:  "200",
		Message: "Berhasil mendapatkan data pembanding",
		Data:    result,
	}, nil
}

func IsiDataInputKeExcel(input *dto.KertasKerjaRequest) (err error) {
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

	savePath := "assets/kertas_kerja/Kertas_Kerja_" + time.Now().Format("20060102150405") + ".xlsx"
	// Simpan file hasil
	return f.SaveAs(savePath)
	// return savePath, err
}
