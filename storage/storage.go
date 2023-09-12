package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var BookNotFound = fmt.Errorf("book not found")

type Book struct {
	ID        string `json:"id"`
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published string `json:"published"`
}

type Storage struct {
	sqlClient *sqlx.DB
}

func NewStorage(sqlClient *sqlx.DB) *Storage {
	return &Storage{sqlClient: sqlClient}
}

func (s *Storage) GetBooks(query string, page, limit int) ([]Book, error) {
	// get books from database
	books := []Book{}
	sqlQuery := `SELECT * FROM books WHERE title LIKE ? LIMIT ? OFFSET ?`
	err := s.sqlClient.Select(&books, sqlQuery, "%"+query+"%", limit, (page - 1))
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *Storage) CreateBook(book Book) (*Book, error) {
	// insert book to database
	sqlQuery := `INSERT INTO books (isbn, title, author, published) VALUES (?, ?, ?, ?)`
	res, err := s.sqlClient.Exec(sqlQuery, book.ISBN, book.Title, book.Author, book.Published)
	if err != nil {
		return nil, err
	}

	// get last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	// set id to book
	book.ID = fmt.Sprintf("%d", id)

	return &book, nil
}

func (s *Storage) GetBookByID(id string) (*Book, error) {
	// TODO: implement this
	return nil, nil
}

func (s *Storage) UpdateBook(id string, book Book) (*Book, error) {
	// TODO: implement this
	return nil, nil
}

func (s *Storage) DeleteBook(id string) error {
	// TODO: implement this
	return nil
}
