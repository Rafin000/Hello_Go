package database

import (
    "fmt"
    "hello_go/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error

    host := os.Getenv("DB_HOST")
    if host == "" {
        host = "localhost"
    }

    user := os.Getenv("DB_USER")
    if user == "" {
        user = "postgres"
    }

    password := os.Getenv("DB_PASSWORD")
    if password == "" {
        password = "postgres"
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "blogdb"
    }

    port := os.Getenv("DB_PORT")
    if port == "" {
        port = "5433"
    }

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbName, port,
    )

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    
    err = DB.AutoMigrate(
        &models.Blog{},
        &models.User{},
        &models.SocialMediaLink{},
        &models.Testimonial{},
        &models.Skill{},
        &models.Experience{},
        &models.Education{},
    )
    if err != nil {
        log.Fatalf("Failed to migrate models: %v", err)
    }

    log.Println("Database connection established")
}
