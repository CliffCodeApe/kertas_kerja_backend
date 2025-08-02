package migrations

import "database/sql"

type createUsersTable struct{}

func (m *createUsersTable) SkipProd() bool {
	return false
}

func getCreateUsersTable() migration {
	return &createUsersTable{}
}

func (m *createUsersTable) Name() string {
	return "create-users"
}

func (m *createUsersTable) Up(conn *sql.Tx) error {

	_, err := conn.Exec(`
	CREATE TABLE users (
		id                 SERIAL PRIMARY KEY,
		nama_satker        varchar(255) NOT NULL,
		kode_kl            varchar(255) NOT NULL,
		email              varchar(255) NOT NULL,
		role               varchar(10) NOT NULL,
		password           varchar(255) NOT NULL,
		updated_at         TIMESTAMP NOT NULL DEFAULT NOW(),
		created_at         TIMESTAMP NOT NULL DEFAULT NOW()
	)`)

	if err != nil {
		return err
	}

	return err

}

func (m *createUsersTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`
	DROP TABLE IF EXISTS users
	`)

	if err != nil {
		return err
	}

	return nil
}
