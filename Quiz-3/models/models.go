package models

import "time"

type (
	Category struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Book struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		ImageURL    string    `json:"image_url"`
		ReleaseYear int       `json:"release_year"`
		Price       string    `json:"price"`
		TotalPage   int       `json:"total_page"`
		Thickness   string    `json:"thickness"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		CategoryID  int       `json:"category_id"`
	}
)