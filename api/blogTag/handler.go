// package blog_tag

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

// type BlogTagHandler struct {
// 	service BlogTagService
// }

// func NewBlogTagHandler(service BlogTagService) *BlogTagHandler {
// 	return &BlogTagHandler{service: service}
// }

// func (h *BlogTagHandler) GetTags(c *gin.Context) {
// 	blogID := c.Param("blog_id")

// 	tags, err := h.service.GetTags(blogID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Blog not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"tags": tags})
// }

// func (h *BlogTagHandler) AddTag(c *gin.Context) {
// 	blogID := c.Param("blog_id")

// 	var requestBody struct {
// 		Tag string `json:"tag"`
// 	}
// 	if err := c.ShouldBindJSON(&requestBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	if requestBody.Tag == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag is required"})
// 		return
// 	}

// 	err := h.service.AddTag(blogID, requestBody.Tag)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Successfully added tag"})
// }

// func (h *BlogTagHandler) DeleteTag(c *gin.Context) {
// 	blogID := c.Param("blog_id")

// 	var requestBody struct {
// 		Tag string `json:"tag"`
// 	}
// 	if err := c.ShouldBindJSON(&requestBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	if requestBody.Tag == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag is required"})
// 		return
// 	}

// 	err := h.service.DeleteTag(blogID, requestBody.Tag)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted tag"})
// }
