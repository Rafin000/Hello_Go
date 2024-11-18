package interfaces

import (
    "hello_go/models"
)

type BlogTagService interface {
	GetTags(blogID string) ([]string, error)
	AddTag(blogID string, tag string) error
	DeleteTag(blogID string, tag string) error
}
 
type BlogTagRepository interface {
	GetBlogByID(blogID string) (*models.Blog, error)
	SaveBlog(blog *models.Blog) error
}