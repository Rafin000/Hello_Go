// package blog_tag

// import (
// 	"gorm.io/gorm"
// 	"hello_go/interfaces"
// )

// type Blog struct {
// 	ID    string   `gorm:"primaryKey"`
// 	Title string   `gorm:"size:255"`
// 	Tags  []string `gorm:"type:text[]"`
// }

// type blogTagRepository struct {
// 	db *gorm.DB
// }

// func NewBlogTagRepository(db *gorm.DB) interfaces.BlogTagRepository {
// 	return &blogTagRepository{db: db}
// }

// func (r *blogTagRepository) GetBlogByID(blogID string) (*Blog, error) {
// 	var blog Blog
// 	if err := r.db.Where("id = ?", blogID).First(&blog).Error; err != nil {
// 		return nil, err
// 	}
// 	return &blog, nil
// }

// func (r *blogTagRepository) SaveBlog(blog *Blog) error {
// 	if err := r.db.Save(blog).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
