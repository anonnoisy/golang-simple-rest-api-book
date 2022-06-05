package book

import "strconv"

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(bookRequest BookRequest, ID int) (Book, error)
	Delete(ID int) error
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
	price, _ := bookRequest.Price.Float64()
	rating, _ := strconv.ParseUint(bookRequest.Rating, 10, 64)

	book := Book{
		Title:       bookRequest.Title,
		Description: string(bookRequest.Description),
		Price:       price,
		Rating:      uint(rating),
	}

	return s.repository.Create(book)
}

func (s *service) Update(bookRequest BookRequest, ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return Book{}, err
	}

	price, _ := bookRequest.Price.Float64()
	rating, _ := strconv.ParseUint(bookRequest.Rating, 10, 64)

	book.Title = bookRequest.Title
	book.Description = string(bookRequest.Description)
	book.Price = price
	book.Rating = uint(rating)

	return s.repository.Update(book)
}

func (s *service) Delete(ID int) error {
	return s.repository.Delete(ID)
}
