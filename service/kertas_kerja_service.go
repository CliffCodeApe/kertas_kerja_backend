package service

import (
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"strconv"
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
	tahunPenilaian := tahunPembuatan // Atau ambil dari field lain jika ada

	// Query ke repository
	dataPembanding, err := s.kertasKerjarRepo.FindDataPembanding(
		req.MerekKendaraan,
		req.TipeKendaraan,
		tahunPembuatan,
		req.LokasiObjek,
		kategoriLokasi,
		tahunPenilaian,
	)
	if err != nil {
		return &dto.KertasKerjaResponse{
			Status:  "500",
			Message: "Gagal mencari data pembanding",
			Data:    dto.KertasKerjaData{},
		}, err
	}

	// Mapping hasil ke DTO response
	var pembanding dto.DataPembanding
	if len(dataPembanding) > 0 {
		lelang := dataPembanding[0] // Ambil data pertama sebagai contoh
		pembanding = dto.DataPembanding{
			KodeLelang:     lelang.Kode,
			Merek:          lelang.Merek,
			Tipe:           lelang.Tipe,
			TahunPembuatan: lelang.TahunPembuatan,
			TahunTransaksi: lelang.TahunLelang,
			Lokasi:         lelang.Kpknl,
			KategoriLokasi: lelang.KategoriLokasi,
			HargaLelang:    float64(lelang.HargaLaku),
		}
	}

	result := dto.KertasKerjaData{
		InputLelang:    *req,
		DataPembanding: []dto.DataPembanding{pembanding},
	}

	return &dto.KertasKerjaResponse{
		Status:  "200",
		Message: "Berhasil mendapatkan data pembanding",
		Data:    result,
	}, nil
}
