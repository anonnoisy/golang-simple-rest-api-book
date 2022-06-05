package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	"pustaka_api/book"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"api":     "Book API V1",
		"version": "1.0",
		"release": "dev",
	})
}

func (h *bookHandler) FindAllBookHandler(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusConflict, err)
	}

	ctx.JSON(http.StatusOK, books)
}

func (h *bookHandler) FindBookHandler(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	book, err := h.bookService.FindByID(int(id))

	if err != nil {
		ctx.JSON(http.StatusConflict, err)
	}

	ctx.JSON(http.StatusOK, book)
}

// func to handle book post route
func (h *bookHandler) PostBookHandler(ctx *gin.Context) {
	var bookRequest book.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		ctx.JSON(http.StatusConflict, err)
	}

	ctx.JSON(http.StatusCreated, book)
}
