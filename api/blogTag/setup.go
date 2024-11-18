// package blog_tag

// import (
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// func SetupBlogTagModule(r *gin.Engine, db *gorm.DB) {
// 	repo := NewBlogTagRepository(db)
// 	service := NewBlogTagService(repo)
// 	handler := NewBlogTagHandler(service)

// 	RegisterRoutes(r, handler)
// }
