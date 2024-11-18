package blog

import (
    "hello_go/models"
    "hello_go/interfaces"
    "github.com/google/uuid"
)

type blogService struct {
    repository interfaces.BlogRepository
}

func NewBlogService(repository interfaces.BlogRepository) interfaces.BlogService {
    return &blogService{repository: repository}
}

func (s *blogService) CreateBlog(blog *models.Blog) error {
    return s.repository.Create(blog)
}

func (s *blogService) GetBlog(id uuid.UUID) (*models.Blog, error) {
    return s.repository.GetByID(id)
}

func (s *blogService) GetAllBlogs() ([]models.Blog, error) {
    return s.repository.GetAll()
}

func (s *blogService) UpdateBlog(blog *models.Blog) error {
    return s.repository.Update(blog)
}

func (s *blogService) DeleteBlog(id uuid.UUID) error {
    return s.repository.Delete(id)
}

func (s *blogService) LikeBlog(id uuid.UUID) (int, error) {
    blog, err := s.repository.GetByID(id)
    if err != nil {
        return 0, err
    }

    blog.Likes++
    if err := s.repository.Update(blog); err != nil {
        return 0, err
    }

    return blog.Likes, nil
}
