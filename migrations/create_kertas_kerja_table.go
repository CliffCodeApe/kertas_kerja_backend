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
		id                  	serial PRIMARY KEY,
		user_id					bigint NOT NULL,
		nama_objek				varchar(100),
		NUP						varchar(50) UNIQUE,
		kode_satker				varchar(50),
		hasil_nilai_taksiran	float8,
		pdf_path				text,
		excel_path				text,
		is_verified	boolean 	NOT NULL DEFAULT false,
		created_at				TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at				TIMESTAMP NOT NULL DEFAULT NOW(),

		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
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
