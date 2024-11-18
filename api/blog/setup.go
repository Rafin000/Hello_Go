package blog

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBlogModule(r *gin.Engine, db *gorm.DB) {
	repo := NewBlogRepository(db)
	service := NewBlogService(repo)
	handler := NewBlogHandler(service)
	
	RegisterRoutes(r, handler)
} 