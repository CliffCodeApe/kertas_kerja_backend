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
		password
	) VALUES
	('John Doe', '015', 'panitia@example.com', 'panitia', '$2a$12$iQlTqu3UViJztxiofh4z0eVUpsOG6rSxUy21CVacWEVlwzoEVcbWm'),
	('Jane Doe', '015', 'admin@example.com', 'admin', '$2a$12$iQlTqu3UViJztxiofh4z0eVUpsOG6rSxUy21CVacWEVlwzoEVcbWm')`)

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
