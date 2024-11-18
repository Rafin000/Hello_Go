package blog

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *BlogHandler) {
    blogGroup := r.Group("/blog")
    {
        blogGroup.POST("/", handler.CreateBlog)
        blogGroup.GET("/", handler.GetBlogs)
        blogGroup.GET("/:blog_id", handler.GetBlog)
        blogGroup.PUT("/:blog_id", handler.UpdateBlog)
        blogGroup.DELETE("/:blog_id", handler.DeleteBlog)
        blogGroup.POST("/:blog_id/like", handler.LikeBlog)
    }
}
