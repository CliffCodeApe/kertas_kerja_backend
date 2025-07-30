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
	// tahunPenilaian := tahunPembuatan // Atau ambil dari field lain jika ada

	// Query ke repository
	dataPembanding, err := s.kertasKerjarRepo.FindDataPembanding(
		req.MerekKendaraan,
		req.TipeKendaraan,
		tahunPembuatan,
		req.LokasiObjek,
		// tahunPenilaian,
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

	return &dto.KertasKerjaResponse{
		Status:  "200",
		Message: "Berhasil mendapatkan data pembanding",
		Data:    result,
	}, nil
}
