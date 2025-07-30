package migrations

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type insertDataLelangTable struct{}

func (m *insertDataLelangTable) SkipProd() bool {
	return false
}

func getInsertDataLelangTable() migration {
	return &insertDataLelangTable{}
}

func (m *insertDataLelangTable) Name() string {
	return "insert-dataLelang"
}

func (m *insertDataLelangTable) Up(conn *sql.Tx) error {

	file, err := os.Open("assets/csv/data_lelang.csv")
	if err != nil {
		return fmt.Errorf("open csv file: %w", err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	_, err = r.Read() // Skip header
	if err != nil {
		return fmt.Errorf("read csv header: %w", err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read csv record: %w", err)
		}
		if len(record) < 13 {
			return fmt.Errorf("csv record too short: %v", record)
		}
		values := make([]any, 13)
		for i := 0; i < 13; i++ {
			val := strings.TrimSpace(record[i])
			// Kolom tahun_pembuatan (misal index ke-10, cek urutan sesuai tabel)
			if i == 10 || i == 12 { // tahun_lelang atau harga_laku
				if val == "" {
					values[i] = nil // akan menjadi NULL di DB
				} else {
					values[i] = val
				}
			} else {
				values[i] = val
			}
		}

		_, err = conn.Exec(`
        INSERT INTO data_lelang (
            kode, kategori_objek, tahun_lelang, provinsi, kota, kpknl,
            kategori_lokasi_jan_jun, kategori_lokasi_jul_des,
            merek, tipe, tahun_pembuatan, warna, harga_laku
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
    `, values...)
		if err != nil {
			return fmt.Errorf("insert record into data_lelang (values: %v): %w", values, err)
		}
	}

	return nil

}

func (m *insertDataLelangTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`
	DELETE FROM data_lelang
	`)

	if err != nil {
		return err
	}

	return nil
}
