package migrations

// import "database/sql"

// type insertUsersTable struct{}

// func (m *insertUsersTable) SkipProd() bool {
// 	return false
// }

// func getInsertUsersTable() migration {
// 	return &insertUsersTable{}
// }

// func (m *insertUsersTable) Name() string {
// 	return "insert_users"
// }

// func (m *insertUsersTable) Up(conn *sql.Tx) error {

// 	_, err := conn.Exec(`
// 	INSERT INTO users (
// 		nama_satker,
// 		kode_satker,
// 		email,
// 		nama,
// 		role,
// 		password
// 	) VALUES
// 	('kementerian keuangan republik indonesia', '015', 'panitia@example.com', 'John Doe', 'panitia', 'password'),
// 	('kementerian keuangan republik indonesia', '015', 'admin@example.com', 'Jane Doe', 'admin', 'password')`)

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

// func (m *insertUsersTable) Down(conn *sql.Tx) error {
// 	_, err := conn.Exec(`
// 	DELETE FROM users WHERE email IN ('panitia@example.com', 'admin@example.com')
// 	`)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
