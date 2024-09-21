package models

import (
	"gorm.io/gorm"
)

// Author represents the authors table
type Author struct {
    AuthorID  uint   `gorm:"primaryKey"`
    FirstName string `gorm:"not null"`
    LastName  string `gorm:"not null"`
    Bio       string
    Books     []Book `gorm:"foreignKey:AuthorID"`
}

// Category represents the categories table
type Category struct {
    CategoryID uint   `gorm:"primaryKey"`
    Name       string `gorm:"not null"`
    Description string
    Books      []Book `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
}

// Book represents the books table
type Book struct {
    BookID          uint    `gorm:"primaryKey"`
    Title           string  `gorm:"not null"`
    ISBN            string  `gorm:"unique"`
    AuthorID        uint
    Author          Author
    Publisher       string
    PublicationYear int
    CategoryID      uint
    Category        Category
    Description     string
}

// BeforeCreate is a GORM hook that runs before a new record is inserted
func (a *Author) BeforeCreate(tx *gorm.DB) (err error) {
    // Implement any logic you need before creating an author
    return
}