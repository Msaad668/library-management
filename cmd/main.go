package main

import (
	"log"

	"github.com/Msad668/library-management/database"
	"github.com/Msad668/library-management/handlers"
	"github.com/Msad668/library-management/models"
	"github.com/Msad668/library-management/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	database.Connect()

	// Auto-migrate models to create/update tables
	err := database.DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Category{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	authorRepo := repository.NewAuthorRepository()
	// bookRepo := repository.NewBookRepository()
	// categoryRepo := repository.NewCategoryRepository()

	// Initialize handlers
	authorHandler := handlers.NewAuthorHandler(authorRepo)
	// bookHandler := handlers.NewBookHandler(bookRepo)
	// categoryHandler := handlers.NewCategoryHandler(categoryRepo)

	// Create a new Gin router
	r := gin.Default()

	// Author routes
	r.POST("/authors", authorHandler.CreateAuthor)
	r.GET("/authors", authorHandler.GetAuthors)

	// Book routes
	// r.POST("/books", bookHandler.CreateBook)
	// r.GET("/books", bookHandler.GetBooks)
	// r.GET("/books/:id", bookHandler.GetBookByID)

	// // Category routes
	// r.POST("/categories", categoryHandler.CreateCategory)
	// r.GET("/categories", categoryHandler.GetCategories)
	// r.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	// Start the server
	r.Run(":8080") // Run the server on port 8080
}