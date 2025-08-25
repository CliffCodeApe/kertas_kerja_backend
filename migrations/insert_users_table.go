package migrations

import "database/sql"

type insertUsersTable struct{}

func (m *insertUsersTable) SkipProd() bool {
	return false
}

func getInsertUsersTable() migration {
	return &insertUsersTable{}
}

func (m *insertUsersTable) Name() string {
	return "insert_users"
}

func (m *insertUsersTable) Up(conn *sql.Tx) error {

	_, err := conn.Exec(`
	INSERT INTO users (
		nama_satker,
		kode_kl,
		email,
		role,
		password,
		is_verified
	) VALUES
	('John Doe', '016', 'panitia@example.com', 'satker', '$2a$12$iQlTqu3UViJztxiofh4z0eVUpsOG6rSxUy21CVacWEVlwzoEVcbWm', 'tervalidasi'),
	('Jane Doe', '017', 'admin@example.com', 'admin', '$2a$12$iQlTqu3UViJztxiofh4z0eVUpsOG6rSxUy21CVacWEVlwzoEVcbWm', 'tervalidasi'),
	('James Doe', '069', 'superadmin@example.com', 'superadmin', '$2a$12$iQlTqu3UViJztxiofh4z0eVUpsOG6rSxUy21CVacWEVlwzoEVcbWm', 'tervalidasi')
	`)

	if err != nil {
		return err
	}

	return nil

}

func (m *insertUsersTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`
	DELETE FROM users WHERE email IN ('panitia@example.com', 'admin@example.com')
	`)

	if err != nil {
		return err
	}

	return nil
}
