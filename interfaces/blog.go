package interfaces

import (
    "hello_go/models"
    "github.com/google/uuid"
)

type BlogRepository interface {
    Create(blog *models.Blog) error
    GetByID(id uuid.UUID) (*models.Blog, error)
    GetAll() ([]models.Blog, error)
    Update(blog *models.Blog) error
    Delete(id uuid.UUID) error
}

type BlogService interface {
    CreateBlog(blog *models.Blog) error
    GetBlog(id uuid.UUID) (*models.Blog, error)
    GetAllBlogs() ([]models.Blog, error)
    UpdateBlog(blog *models.Blog) error
    DeleteBlog(id uuid.UUID) error
    LikeBlog(id uuid.UUID) (int, error)
} 