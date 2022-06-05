package main

import (
	"fmt"
	"log"
	"pustaka_api/book"
	"pustaka_api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Set to development mode
	gin.SetMode(gin.DebugMode)

	// Connecting to database
	dsn := "root:@tcp(127.0.0.1:3306)/go_pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error")
	}
	fmt.Println("Database connection established")

	// Migrating the table from struct
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("api/v1")
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/books", bookHandler.FindAllBookHandler)
	v1.GET("/book/:id", bookHandler.FindBookHandler)
	v1.POST("/book", bookHandler.PostBookHandler)
	v1.PUT("/book/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/book/:id", bookHandler.DeleteBookHandler)

	router.Run()
}
