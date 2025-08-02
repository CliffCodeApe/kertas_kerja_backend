package migrations

import "database/sql"

type createLelangKendaraanTable struct{}

func (m *createLelangKendaraanTable) SkipProd() bool {
	return false
}

func getCreateLelangKendaraanTable() migration {
	return &createLelangKendaraanTable{}
}

func (m *createLelangKendaraanTable) Name() string {
	return "create-lelang-kendaraan"
}

func (m *createLelangKendaraanTable) Up(conn *sql.Tx) error {

	_, err := conn.Exec(`
	CREATE TABLE lelang_kendaraan (
		id                	serial PRIMARY KEY,
		nama_objek			varchar(100),
		lokasi_objek		varchar(100),

		created_at			TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at			TIMESTAMP NOT NULL DEFAULT NOW()
	)`)

	if err != nil {
		return err
	}

	return err

}

func (m *createLelangKendaraanTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`
	DROP TABLE IF EXISTS lelang_kendaraan
	`)

	if err != nil {
		return err
	}

	return nil
}
