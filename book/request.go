package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description"`
	Price       json.Number `json:"price" binding:"required,numeric"`
	Rating      string      `json:"rating" binding:"required,number"`
}
