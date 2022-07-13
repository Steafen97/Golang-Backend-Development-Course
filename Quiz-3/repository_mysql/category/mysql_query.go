package category

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
	table          = "categories"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
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
		var category models.Category
		var createdAt, updatedAt string

		if err = rowQuery.Scan(
			&category.ID,
			&category.Name,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		category.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		category.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func Insert(ctx context.Context, category models.Category) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"INSERT INTO %v (name, created_at, updated_at) values('%v', NOW(), NOW())",
		table,
		category.Name,
	)
	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func Update(ctx context.Context, category models.Category, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"UPDATE %v set name = '%s', updated_at = NOW() where id = %s",
		table,
		category.Name,
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
		"DELETE FROM %v where id = %s",
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

func BooksByForeignKey(ctx context.Context, id string) ([]models.Book, error) {
	var books []models.Book
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"SELECT * FROM books WHERE category_id = %v Order BY id DESC",
		id,
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
