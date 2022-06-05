package book

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Price       string `json:"price" binding:"required,number"`
	Rating      string `json:"rating" binding:"required,number"`
}
