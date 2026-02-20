package service

import (
	"errors"

	"github.com/devalejandroaf/crud/internal/model"
	"github.com/devalejandroaf/crud/internal/store"
)

type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) GetAllBooks() ([]*model.Book, error) {
	books, err := s.store.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *Service) GetBookById(id int) (*model.Book, error) {
	return s.store.GetByID(id)
}

func (s *Service) CreateBook(book model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("Title can not be empty")
	}
	return s.store.Create(&book)
}

func (s *Service) UpdateBook(id int, book model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("Title can not be empty")
	}
	return s.store.Update(id, &book)
}

func (s *Service) RemoveBook(id int) error {
	return s.store.Delete(id)
}
