package mata_kuliah

import (
	"Tugas-15/config"
	"Tugas-15/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "mata_kuliah"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.MataKuliah, error) {
	var matkuls []models.MataKuliah
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
		var matkul models.MataKuliah
		var createdAt, updatedAt string

		if err = rowQuery.Scan(
			&matkul.ID,
			&matkul.Nama,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		matkul.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		matkul.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		matkuls = append(matkuls, matkul)
	}

	return matkuls, nil
}

func Insert(ctx context.Context, matkul models.MataKuliah) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"INSERT INTO %v (nama, created_at, updated_at) values('%s', NOW(), NOW())",
		table,
		matkul.Nama,
	)
	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func Update(ctx context.Context, matkul models.MataKuliah, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"UPDATE %v set nama = '%s', updated_at = NOW() WHERE id = %s",
		table,
		matkul.Nama,
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
		"DELETE FROM %v WHERE id = %s",
		table,
		id,
	)
	s, err := db.ExecContext(ctx, queryText)
	if err != nil && err != sql.ErrNoRows {
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
