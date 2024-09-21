package repository

import (
    "github.com/Msad668/library-management/database"
    "github.com/Msad668/library-management/models"
)

// CategoryRepository defines methods for category data access
type CategoryRepository struct{}

// NewCategoryRepository creates a new CategoryRepository
func NewCategoryRepository() *CategoryRepository {
    return &CategoryRepository{}
}

// Create adds a new category to the database
func (r *CategoryRepository) Create(category *models.Category) error {
    return database.DB.Create(category).Error
}

// FindAll retrieves all categories from the database
func (r *CategoryRepository) FindAll() ([]models.Category, error) {
    var categories []models.Category
    err := database.DB.Find(&categories).Error
    return categories, err
}

// FindByID retrieves a single category by its ID
func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
    var category models.Category
    err := database.DB.First(&category, id).Error
    return &category, err
}

// Update modifies an existing category's details
func (r *CategoryRepository) Update(category *models.Category) error {
    return database.DB.Save(category).Error
}

// Delete removes a category from the database
func (r *CategoryRepository) Delete(id uint) error {
    return database.DB.Delete(&models.Category{}, id).Error
}