package service

import (
	"encoding/csv"
	"fmt"
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"kertas_kerja/entity"
	"os"
)

type dataLelangServ struct {
	DataLelangRepo contract.DataLelangRepository
}

func implDataLelangService(repo *contract.Repository) contract.DataLelangService {
	return &dataLelangServ{
		DataLelangRepo: repo.DataLelang,
	}
}

func (d *dataLelangServ) InsertDataLelang(payload *dto.DataLelang) (*dto.InsertDataLelangResponse, error) {

	dataLelang := &entity.DataLelang{
		Kode:                 payload.Kode,
		KategoriObjek:        payload.KategoriObjek,
		TahunLelang:          payload.TahunLelang,
		Provinsi:             payload.Provinsi,
		Kota:                 payload.Kota,
		Kpknl:                payload.Kpknl,
		KategoriLokasiJanJun: payload.KategoriLokasiJanuariDanJuni,
		KategoriLokasiJulDes: payload.KategoriLokasiJuliDanDesember,
		Merek:                payload.Merek,
		Tipe:                 payload.Tipe,
		TahunPembuatan:       payload.TahunPembuatan,
		Warna:                payload.Warna,
		HargaLaku:            payload.HargaLaku,
	}

	if err := d.DataLelangRepo.InsertDataLelang(dataLelang); err != nil {
		return nil, fmt.Errorf("failed to insert data lelang: %w", err)
	}

	// Append to CSV
	csvPath := "assets/csv/data_lelang.csv"
	file, err := os.OpenFile(csvPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		writer := csv.NewWriter(file)
		defer file.Close()
		defer writer.Flush()
		record := []string{
			dataLelang.Kode,
			dataLelang.KategoriObjek,
			fmt.Sprintf("%d", dataLelang.TahunLelang),
			dataLelang.Provinsi,
			dataLelang.Kota,
			dataLelang.Kpknl,
			fmt.Sprintf("%d", dataLelang.KategoriLokasiJanJun),
			fmt.Sprintf("%d", dataLelang.KategoriLokasiJulDes),
			dataLelang.Merek,
			dataLelang.Tipe,
			fmt.Sprintf("%d", dataLelang.TahunPembuatan),
			dataLelang.Warna,
			fmt.Sprintf("%.2f", dataLelang.HargaLaku),
		}
		_ = writer.Write(record)
	}

	response := &dto.InsertDataLelangResponse{
		Status:  "success",
		Message: "Data lelang successfully inserted",
		Data: dto.DataLelang{
			Kode:                          dataLelang.Kode,
			KategoriObjek:                 dataLelang.KategoriObjek,
			TahunLelang:                   dataLelang.TahunLelang,
			Provinsi:                      dataLelang.Provinsi,
			Kota:                          dataLelang.Kota,
			Kpknl:                         dataLelang.Kpknl,
			KategoriLokasiJanuariDanJuni:  dataLelang.KategoriLokasiJanJun,
			KategoriLokasiJuliDanDesember: dataLelang.KategoriLokasiJulDes,
			Merek:                         dataLelang.Merek,
			Tipe:                          dataLelang.Tipe,
			TahunPembuatan:                dataLelang.TahunPembuatan,
			Warna:                         dataLelang.Warna,
			HargaLaku:                     dataLelang.HargaLaku,
			CreatedAt:                     dataLelang.CreatedAt,
			UpdatedAt:                     dataLelang.UpdatedAt,
		},
	}

	return response, nil

}
