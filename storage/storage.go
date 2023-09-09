package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
	return s.booksData, nil
}
