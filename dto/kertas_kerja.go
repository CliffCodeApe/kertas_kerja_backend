package dto

type KertasKerjaRequest struct {
	NamaObjek           string `json:"nama_objek"`
	LokasiObjek         string `json:"lokasi_objek"`
	NUP                 string `json:"nup"`
	KategoriLokasi      string `json:"kategori_lokasi"`
	MerekKendaraan      string `json:"merek_kendaraan"`
	TipeKendaraan       string `json:"tipe_kendaraan"`
	NomorPolisi         string `json:"nomor_polisi"`
	DokumenKepemilikan  string `json:"dokumen_kepemilikan"`
	PemilikDokumen      string `json:"pemilik_dokumen"`
	JenisKendaraan      string `json:"jenis_kendaraan"`
	MasaBerlaku         string `json:"masa_berlaku"`
	PenggunaanKendaraan string `json:"penggunaan_kendaraan"`
	Keterangan          string `json:"keterangan"`
	Warna               string `json:"warna"`
	TahunPembuatan      int    `json:"tahun_pembuatan"`
	BahanBakar          string `json:"bahan_bakar"`
	KondisiKendaraan    string `json:"kondisi_kendaraan"`
	TahunLelang         int    `json:"tahun_lelang"`
	Provinsi            string `json:"provinsi"`
}

type DataPembanding struct {
	KodeLelang     string
	Merek          string
	Tipe           string
	TahunPembuatan int
	TahunTransaksi int
	Lokasi         string
	KategoriLokasi int
	HargaLelang    float64
}

type DataPenyesuaian struct {
	Tipe           string
	Merek          string
	Waktu          string
	Lokasi         string
	TahunPembuatan int
	Total          float64
	NilaiTaksiran  float64
}

type KertasKerjaData struct {
	InputLelang    KertasKerjaRequest `json:"input_lelang"`
	DataPembanding []DataPembanding   `json:"data_pembanding"`
}

type KertasKerjaResponse struct {
	Status  string          `json:"status_code"`
	Message string          `json:"message"`
	Data    KertasKerjaData `json:"data"`
}

type DataPembandingResponse struct {
	Status  string         `json:"status_code"`
	Message string         `json:"message"`
	Data    DataPembanding `json:"data"`
}
