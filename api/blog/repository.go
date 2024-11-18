package blog

import (
	"hello_go/models"
	"hello_go/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) interfaces.BlogRepository {
	return &blogRepository{db: db}
}

func (r *blogRepository) Create(blog *models.Blog) error {
	return r.db.Create(blog).Error
}

func (r *blogRepository) GetByID(id uuid.UUID) (*models.Blog, error) {
	var blog models.Blog
	if err := r.db.First(&blog, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *blogRepository) GetAll() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := r.db.Order("created_at desc").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *blogRepository) Update(blog *models.Blog) error {
	return r.db.Save(blog).Error
}

func (r *blogRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Blog{}, "id = ?", id).Error
}
