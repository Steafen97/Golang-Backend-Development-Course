package mahasiswa

import (
	"Tugas-14/config"
	"Tugas-14/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.Mahasiswa, error) {
	var mahasiswas []models.Mahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var mahasiswa models.Mahasiswa
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&mahasiswa.ID,
			&mahasiswa.Nama,
			&mahasiswa.MataKuliah,
			&mahasiswa.IndeksNilai,
			&mahasiswa.Nilai,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		mahasiswas = append(mahasiswas, mahasiswa)
	}

	return mahasiswas, nil
}

func Insert(ctx context.Context, mahasiswa models.Mahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama, mata_kuliah, indeks_nilai, nilai, created_at, updated_at) values('%v', '%v', '%v', %v, NOW(), NOW())",
		table,
		mahasiswa.Nama,
		mahasiswa.MataKuliah,
		mahasiswa.IndeksNilai,
		mahasiswa.Nilai,
	)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func Update(ctx context.Context, mahasiswa models.Mahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama = '%s', mata_kuliah = '%s', indeks_nilai = '%s', nilai = %d, updated_at = NOW() where id = %s",
		table,
		mahasiswa.Nama,
		mahasiswa.MataKuliah,
		mahasiswa.IndeksNilai,
		mahasiswa.Nilai,
		id,
	)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)
	s, err := db.ExecContext(ctx, queryText)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
