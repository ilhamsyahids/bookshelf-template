package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
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
	booksData []Book
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Load(filename string) error {
	rawBook, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read file due: %v", err)
	}

	err = json.Unmarshal(rawBook, &s.booksData)
	if err != nil {
		return fmt.Errorf("unable to init books data due: %v", err)
	}

	return nil
}

func (s *Storage) GetBooks(query string, page, limit int) ([]Book, error) {
	var books []Book
	for _, book := range s.booksData {
		lowerTitle := strings.ToLower(book.Title)
		lowerQuery := strings.ToLower(query)
		if strings.Contains(lowerTitle, lowerQuery) {
			books = append(books, book)
		}
	}

	// pagination
	startIndex := (page - 1) * limit
	endIndex := page * limit
	if startIndex >= len(books) {
		return []Book{}, nil
	}
	if endIndex > len(books) {
		endIndex = len(books)
	}
	books = books[(page-1)*limit : endIndex]

	return books, nil
}

func (s *Storage) CreateBook(book Book) (*Book, error) {
	book.ID = fmt.Sprintf("%d", len(s.booksData)+1)
	s.booksData = append(s.booksData, book)
	return &book, nil
}

func (s *Storage) GetBookByID(id string) (*Book, error) {
	for _, book := range s.booksData {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, BookNotFound
}

func (s *Storage) UpdateBook(id string, book Book) (*Book, error) {
	for i, b := range s.booksData {
		if b.ID == id {
			if book.ISBN != "" {
				s.booksData[i].ISBN = book.ISBN
			}
			if book.Title != "" {
				s.booksData[i].Title = book.Title
			}
			if book.Author != "" {
				s.booksData[i].Author = book.Author
			}
			if book.Published != "" {
				s.booksData[i].Published = book.Published
			}
			return &s.booksData[i], nil
		}
	}
	return nil, BookNotFound
}

func (s *Storage) DeleteBook(id string) error {
	for i, book := range s.booksData {
		if book.ID == id {
			s.booksData = append(s.booksData[:i], s.booksData[i+1:]...)
			return nil
		}
	}
	return BookNotFound
}
