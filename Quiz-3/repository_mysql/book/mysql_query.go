package book

import (
	"Quiz-3/config"
	"Quiz-3/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "books"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"SELECT * FROM %v Order BY id DESC",
		table,
	)
	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		var createdAt, updatedAt string

		if err = rowQuery.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageURL,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&createdAt,
			&updatedAt,
			&book.CategoryID,
		); err != nil {
			return nil, err
		}

		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	return books, nil
}

func Insert(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"INSERT INTO %v (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) values('%v', '%v', '%v', %v, '%v', %v, '%v', NOW(), NOW(), %v)",
		table,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
	)
	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func Update(ctx context.Context, book models.Book, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"UPDATE %v set title = '%v', description = '%v', image_url = '%v', release_year = %v, price = '%v', total_page = %v, thickness = '%v', updated_at = NOW(), category_id = %v where id = %v",
		table,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
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
		"DELETE FROM %v where id = %v",
		table,
		id,
	)
	s, err := db.ExecContext(ctx, queryText)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	if check == 0 {
		return errors.New("ID's not found")
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}