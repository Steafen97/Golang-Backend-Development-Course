package nilai

import (
	"Tugas-15/config"
	"Tugas-15/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "nilai"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.Nilai, error) {
	var nilais []models.Nilai
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"SELECT * FROM %v ORDER BY id DESC",
		table,
	)
	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai models.Nilai
		var createdAt, updatedAt string

		if err = rowQuery.Scan(
			&nilai.ID,
			&nilai.Nilai,
			&nilai.IndeksNilai,
			&createdAt,
			&updatedAt,
			&nilai.MahasiswaID,
			&nilai.MataKuliahID,
		); err != nil {
			return nil, err
		}

		nilai.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		nilais = append(nilais, nilai)
	}

	return nilais, nil
}

func Insert(ctx context.Context, nilai models.Nilai) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal()
	}

	queryText := fmt.Sprintf(
		"INSERT INTO %v (nilai, indeks_nilai, created_at, updated_at, mahasiswa_id, mata_kuliah_id) values(%v, '%s', NOW(), NOW(), %v, %v)",
		table,
		nilai.Nilai,
		nilai.IndeksNilai,
		nilai.MahasiswaID,
		nilai.MataKuliahID,
	)
	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func Update(ctx context.Context, nilai models.Nilai, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"UPDATE %v set nilai = %v, indeks_nilai = '%v', updated_at = NOW() WHERE id = %v",
		table,
		nilai.Nilai,
		nilai.IndeksNilai,
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

	queryText := fmt.Sprintf(
		"DELETE FROM %v WHERE id = %v",
		table,
		id,
	)
	s, err := db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	check, err := s.RowsAffected()
	if check == 0 {
		return errors.New("ID tidak ditemukan")
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
