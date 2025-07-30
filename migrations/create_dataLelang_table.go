package migrations

import "database/sql"

type createDataLelangTable struct{}

func (m *createDataLelangTable) SkipProd() bool {
	return false
}

func getCreateDataLelangTable() migration {
	return &createDataLelangTable{}
}

func (m *createDataLelangTable) Name() string {
	return "create-dataLelang"
}

func (m *createDataLelangTable) Up(conn *sql.Tx) error {

	_, err := conn.Exec(`
	CREATE TABLE data_lelang (
		kode                 varchar(10) PRIMARY KEY,
		kategori_objek       varchar(20),
		tahun_lelang         int,
		provinsi             varchar(50),
		kota                 varchar(50),
		kpknl                varchar(50),
		kategori_lokasi_jan_jun  smallint,
		kategori_lokasi_jul_des  smallint,
		merek                varchar(50),
		tipe                 varchar(100),
		tahun_pembuatan      int,
		warna                varchar(50),
		harga_laku           numeric(15,0),
		updated_at           TIMESTAMP NOT NULL DEFAULT NOW(),
		created_at           TIMESTAMP NOT NULL DEFAULT NOW()
	)`)

	if err != nil {
		return err
	}

	return err

}

func (m *createDataLelangTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`
	DROP TABLE IF EXISTS data_lelang
	`)

	if err != nil {
		return err
	}

	return nil
}
