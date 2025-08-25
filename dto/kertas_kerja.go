package dto

type IdentitasKendaraan struct {
	NamaObjek           string   `json:"nama_objek"`
	LokasiObjek         string   `json:"lokasi_objek"`
	NUP                 string   `json:"nup"`
	KategoriLokasi      string   `json:"kategori_lokasi"`
	MerekKendaraan      string   `json:"merek_kendaraan"`
	TipeKendaraan       string   `json:"tipe_kendaraan"`
	NomorPolisi         string   `json:"nomor_polisi"`
	DokumenKepemilikan  []string `json:"dokumen_kepemilikan"`
	PemilikDokumen      string   `json:"pemilik_dokumen"`
	JenisKendaraan      string   `json:"jenis_kendaraan"`
	MasaBerlaku         string   `json:"masa_berlaku"`
	PenggunaanKendaraan string   `json:"penggunaan_kendaraan"`
	Keterangan          string   `json:"keterangan"`
	Warna               string   `json:"warna"`
	TahunPembuatan      int      `json:"tahun_pembuatan"`
	BahanBakar          string   `json:"bahan_bakar"`
	KondisiKendaraan    string   `json:"kondisi_kendaraan"`
	Provinsi            string   `json:"provinsi"`
}

type KertasKerjaResponse struct {
	Status  string          `json:"status_code"`
	Message string          `json:"message"`
	Data    KertasKerjaData `json:"data"`
}

type KertasKerjaData struct {
	InputLelang    IdentitasKendaraan `json:"input_lelang"`
	DataPembanding []DataPembanding   `json:"data_pembanding"`
}

type IsiKertasKerjaRequest struct {
	InputLelang     IdentitasKendaraan `json:"input_lelang"`
	DataPembanding  []DataPembanding   `json:"data_pembanding"`
	DataPenyesuaian []DataPenyesuaian  `json:"data_penyesuaian"`
	HasilTaksiran   HasilTaksiran      `json:"hasil_taksiran"`
	NUP             string             `json:"nup"`
	FaktorKondisi   float64            `json:"faktor_kondisi"`
}

type DataPembandingResponse struct {
	Status  string         `json:"status_code"`
	Message string         `json:"message"`
	Data    DataPembanding `json:"data"`
}

type DataPembanding struct {
	KodeLelang     string  `json:"kode_lelang"`
	Merek          string  `json:"merek"`
	Tipe           string  `json:"tipe"`
	TahunPembuatan int     `json:"tahun_pembuatan"`
	Kpknl          string  `json:"kpknl"`
	Lokasi         string  `json:"lokasi"`
	KategoriLokasi int     `json:"kategori_lokasi"`
	HargaLelang    float64 `json:"harga_lelang"`
	TahunTransaksi int     `json:"tahun_transaksi"`
}

type DataPenyesuaian struct {
	DataHasilLelang string  `json:"data_hasil_lelang"`
	Tipe            float64 `json:"tipe"`
	Merek           float64 `json:"merek"`
	Waktu           float64 `json:"waktu"`
	Lokasi          float64 `json:"lokasi"`
	TahunPembuatan  float64 `json:"tahun_pembuatan"`
	Total           float64 `json:"total"`
	NilaiTaksiran   float64 `json:"nilai_taksiran"`
}

type HasilTaksiran struct {
	TotalNilaiTaksiran       float64 `json:"total_nilai_taksiran"`
	RataRataNilaiTaksiran    float64 `json:"rata_rata_nilai_taksiran"`
	TaksiranNilaiLimitLelang float64 `json:"taksiran_nilai_limit_lelang"`
	Pembulatan               float64 `json:"pembulatan"`
}

type GetAllRiwayatKertasKerjaResponse struct {
	Status  string                   `json:"status_code"`
	Message string                   `json:"message"`
	Data    []RiwayatKertasKerjaData `json:"data"`
}

type RiwayatKertasKerjaRequest struct {
	UserID    uint64 `json:"user_id"`
	NamaObjek string `json:"nama_objek"`
	ExcelPath string `json:"excel_path"`
}

type RiwayatKertasKerjaResponse struct {
	Status  string                 `json:"status_code"`
	Message string                 `json:"message"`
	Data    RiwayatKertasKerjaData `json:"data"`
}

type DeleteRiwayatKertasKerjaResponse struct {
	Status  string `json:"status_code"`
	Message string `json:"message"`
}

type RiwayatKertasKerjaData struct {
	ID                 uint64  `json:"id"`
	UserID             uint64  `json:"user_id,omitempty"`
	KodeSatker         string  `json:"kode_satker"`
	NUP                string  `json:"nup"`
	HasilNilaiTaksiran float64 `json:"hasil_nilai_taksiran"`
	NamaObjek          string  `json:"nama_objek"`
	PdfPath            string  `json:"pdf_path"`
	ExcelPath          string  `json:"excel_path"`
	IsVerified         bool    `json:"is_verified"`
	KodeKL             string  `json:"kode_kl"`
}

type ValidasiKertasKerjaResponse struct {
	Status  string `json:"status_code"`
	Message string `json:"message"`
}
