package main

import (
    "hello_go/config/database"
    "hello_go/api/blog"
    // "hello_go/api/blogTag"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    database.ConnectDatabase()

    blog.SetupBlogModule(r, database.DB)
    // blog_tag.SetupBlogTagModule(r, database.DB)

    r.GET("/alive", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "alive"})
    })

    log.Println("Server starting on :8081")
    r.Run(":8083")
}
