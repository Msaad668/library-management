package handlers

import (
	"net/http"

	"github.com/Msad668/library-management/models"
	"github.com/Msad668/library-management/repository"
	"github.com/gin-gonic/gin"
)

// AuthorHandler handles CRUD operations for authors
type AuthorHandler struct {
	Repo *repository.AuthorRepository
}

// NewAuthorHandler creates a new AuthorHandler
func NewAuthorHandler(repo *repository.AuthorRepository) *AuthorHandler {
	return &AuthorHandler{Repo: repo}
}

// CreateAuthor creates a new author
func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.Create(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, author)
}

// GetAuthors returns all authors
func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	authors, err := h.Repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authors)
}