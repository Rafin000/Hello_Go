// package blog_tag

// import (
//     "github.com/gin-gonic/gin"
// )

// func RegisterRoutes(r *gin.Engine, handler *BlogTagHandler) {
//     blogGroup := r.Group("/blog-tags")
//     {
//         blogGroup.GET("/:blog_id/tag", handler.GetTags)
//         blogGroup.POST("/:blog_id/tag", handler.AddTag)
//         blogGroup.DELETE("/:blog_id/tag", handler.DeleteTag)
//     }
// }
