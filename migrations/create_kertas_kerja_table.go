package migrations

import "database/sql"

type createKertasKerjaTable struct{}

func (m *createKertasKerjaTable) SkipProd() bool {
	return false
}

func getCreateKertasKerjaTable() migration {
	return &createKertasKerjaTable{}
}

func (m *createKertasKerjaTable) Name() string {
	return "create-kertas-kerja"
}

func (m *createKertasKerjaTable) Up(conn *sql.Tx) error {

	_, err := conn.Exec(`
	CREATE TABLE kertas_kerja (
		id                  serial PRIMARY KEY,
		nama_objek			varchar(100),
		created_at			TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at			TIMESTAMP NOT NULL DEFAULT NOW()
	)`)

	if err != nil {
		return err
	}

	return err

}

func (m *createKertasKerjaTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`
	DROP TABLE IF EXISTS kertas_kerja
	`)

	if err != nil {
		return err
	}

	return nil
}
