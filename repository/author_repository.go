package repository

import (
    "github.com/Msad668/library-management/database"
    "github.com/Msad668/library-management/models"
)

// AuthorRepository defines methods for author data access
type AuthorRepository struct{}

// NewAuthorRepository creates a new AuthorRepository
func NewAuthorRepository() *AuthorRepository {
    return &AuthorRepository{}
}

// Create adds a new author to the database
func (r *AuthorRepository) Create(author *models.Author) error {
    return database.DB.Create(author).Error
}

// FindAll retrieves all authors
func (r *AuthorRepository) FindAll() ([]models.Author, error) {
    var authors []models.Author
    err := database.DB.Find(&authors).Error
    return authors, err
}