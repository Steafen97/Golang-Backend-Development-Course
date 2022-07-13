package crud

import (
	"Quiz-3/models"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func Middleware(r *http.Request) string {
	uname, pwd, ok := r.BasicAuth()
	if !ok {
		return "Username atau Password tidak boleh kosong"
	}

	if uname == "admin" && pwd == "password" {
		return ""
	} else if uname == "editor" && pwd == "secret" {
		return ""
	} else if uname == "trainer" && pwd == "rahasia" {
		return ""
	} else {
		return "Username atau Password tidak sesuai"
	}
}

func InputValidation(imageURL string, year int) map[string]string {
	invalids := make(map[string]string)
	_, err := url.ParseRequestURI(imageURL)
	if err != nil {
		invalids["Invalid URL"] = "Please input a valid image_url with format http://example.com"
	}

	if year < 1980 || year > 2021 {
		invalids["Invalid Year"] = "Please input release_year between 1980 and 2021"
	}

	return invalids
}

func ThicknessConvertion(page int) string {
	if page <= 100 {
		return "Tipis"
	} else if page <= 200 {
		return "Sedang"
	} else {
		return "Tebal"
	}
}

func BooksFilter(books []models.Book, r *http.Request) []models.Book {
	title := r.URL.Query().Get("title")
	minYear := r.URL.Query().Get("minYear")
	maxYear := r.URL.Query().Get("maxYear")
	minPage := r.URL.Query().Get("minPage")
	maxPage := r.URL.Query().Get("maxPage")
	sortByTitle := r.URL.Query().Get("sortByTitle")

	if title != "" {
		books = filterByTitle(books, title)
	}

	if minYear != "" {
		intYear, _ := strconv.Atoi(minYear)
		books = filterByMinYear(books, intYear)
	}

	if maxYear != "" {
		intYear, _ := strconv.Atoi(maxPage)
		books = filterByMaxYear(books, intYear)
	}

	if minPage != "" {
		intPage, _ := strconv.Atoi(minPage)
		books = filterByMinPage(books, intPage)
	}

	if maxPage != "" {
		intPage, _ := strconv.Atoi(maxPage)
		books = filterByMaxPage(books, intPage)
	}

	if sortByTitle == "desc" {
		books = sortByTitleDesc(books)
	} else if sortByTitle == "asc" {
		books = sortByTitleAsc(books)
	}

	return books
}

func filterByTitle(books []models.Book, title string) (filtered []models.Book) {
	for _, data := range books {
		if strings.Contains(strings.ToLower(data.Title), strings.ToLower(title)) {
			filtered = append(filtered, data)
		}
	}

	return
}

func filterByMinYear(books []models.Book, minYear int) (filtered []models.Book) {
	for _, data := range books {
		if data.ReleaseYear >= minYear {
			filtered = append(filtered, data)
		}
	}

	return
}

func filterByMaxYear(books []models.Book, maxYear int) (filtered []models.Book) {
	for _, data := range books {
		if data.ReleaseYear <= maxYear {
			filtered = append(filtered, data)
		}
	}

	return
}

func filterByMinPage(books []models.Book, minPage int) (filtered []models.Book) {
	for _, data := range books {
		if data.TotalPage >= minPage {
			filtered = append(filtered, data)
		}
	}

	return
}

func filterByMaxPage(books []models.Book, maxPage int) (filtered []models.Book) {
	for _, data := range books {
		if data.TotalPage <= maxPage {
			filtered = append(filtered, data)
		}
	}

	return
}

func sortByTitleDesc(books []models.Book) []models.Book {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})

	return books
}

func sortByTitleAsc(books []models.Book) []models.Book {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Title > books[j].Title
	})

	return books
}