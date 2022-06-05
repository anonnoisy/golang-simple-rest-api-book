package book

import "strconv"

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := strconv.ParseFloat(bookRequest.Price, 64)
	rating, _ := strconv.ParseUint(bookRequest.Rating, 10, 64)

	book := Book{
		Title:       bookRequest.Title,
		Description: string(bookRequest.Description),
		Price:       price,
		Rating:      uint(rating),
	}

	return s.repository.Create(book)
}
