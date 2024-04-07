package main

import (
	"database/sql"
	"log"
	"time"
)

type Book struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"publishedDate"`
}

// CreateBook inserts a new book into the database
func CreateBook(db *sql.DB, book Book) error {
	_, err := db.Exec("INSERT INTO books (title, author, published_date) VALUES (?, ?, ?)", book.Title, book.Author, book.PublishedDate)
	return err
}

// GetBooks retrieves all books from the database
func GetBooks(db *sql.DB) ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author, published_date FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.PublishedDate)
		if err != nil {
			log.Println(err)
			continue
		}
		books = append(books, b)
	}

	return books, nil
}

// GetBook retrieves a specific book from the database using its ID
func GetBook(db *sql.DB, bookID int) (*Book, error) {
	row := db.QueryRow("SELECT id, title, author, published_date FROM books WHERE id = ?", bookID)

	var b Book
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.PublishedDate)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// UpdateBook updates details of a book in the database
func UpdateBook(db *sql.DB, book Book) error {
	_, err := db.Exec("UPDATE books SET title = ?, author = ?, published_date = ? WHERE id = ?", book.Title, book.Author, book.PublishedDate, book.ID)
	return err
}

// DeleteBook removes a book from the database using its ID
func DeleteBook(db *sql.DB, bookID int) error {
	_, err := db.Exec("DELETE FROM books WHERE id = ?", bookID)
	return err
}
