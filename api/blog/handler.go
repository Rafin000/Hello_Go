package blog

import (
	"hello_go/models"
	"hello_go/interfaces"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlogHandler struct {
	service interfaces.BlogService
}

func NewBlogHandler(service interfaces.BlogService) *BlogHandler {
	return &BlogHandler{service: service}
}

// CreateBlog handles creating a new blog post
func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog data"})
		return
	}

	if err := h.service.CreateBlog(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog created successfully", "blog": blog})
}

// GetBlog fetches a blog by ID
func (h *BlogHandler) GetBlog(c *gin.Context) {
	id := c.Param("blog_id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	blog, err := h.service.GetBlog(parsedID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog})
}

// GetBlogs fetches all blogs
func (h *BlogHandler) GetBlogs(c *gin.Context) {
	blogs, err := h.service.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

// UpdateBlog updates a specific blog post
func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	id := c.Param("blog_id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog data"})
		return
	}

	blog.ID = parsedID
	if err := h.service.UpdateBlog(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully", "blog": blog})
}

// DeleteBlog deletes a specific blog post
func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	id := c.Param("blog_id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.DeleteBlog(parsedID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

// LikeBlog increments the like count of a blog
func (h *BlogHandler) LikeBlog(c *gin.Context) {
	id := c.Param("blog_id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	likes, err := h.service.LikeBlog(parsedID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog liked", "likes": likes})
}
