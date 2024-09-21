package repository

import (
    "github.com/Msad668/library-management/database"
    "github.com/Msad668/library-management/models"
)

// BookRepository defines methods for book data access
type BookRepository struct{}

// NewBookRepository creates a new BookRepository
func NewBookRepository() *BookRepository {
    return &BookRepository{}
}

// Create adds a new book to the database
func (r *BookRepository) Create(book *models.Book) error {
    return database.DB.Create(book).Error
}

// FindAll retrieves all books from the database
func (r *BookRepository) FindAll() ([]models.Book, error) {
    var books []models.Book
    // Preload associated Author and Category
    err := database.DB.Preload("Author").Preload("Category").Find(&books).Error
    return books, err
}

// FindByID retrieves a single book by its ID
func (r *BookRepository) FindByID(id uint) (*models.Book, error) {
    var book models.Book
    // Preload associated Author and Category
    err := database.DB.Preload("Author").Preload("Category").First(&book, id).Error
    return &book, err
}

// Update modifies an existing book's details
func (r *BookRepository) Update(book *models.Book) error {
    return database.DB.Save(book).Error
}

// Delete removes a book from the database
func (r *BookRepository) Delete(id uint) error {
    return database.DB.Delete(&models.Book{}, id).Error
}