package models

import (
    "github.com/google/uuid"
    "github.com/lib/pq"
    "gorm.io/gorm"
    "time"
)

type Blog struct {
    ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Title        string    `json:"title" gorm:"size:255;not null"`
    Content      string    `json:"content" gorm:"type:text;not null"`
    Author       string    `json:"author" gorm:"size:100;not null"`
    Summary      string    `json:"summary" gorm:"size:500"`
    ReadingTime  string    `json:"reading_time" gorm:"size:50"`
    ThumbnailURL string    `json:"thumbnail_url" gorm:"size:255"`
    Tags         pq.StringArray `json:"tags" gorm:"type:text[]"`
    Likes        int       `json:"likes" gorm:"default:0;not null"`
    CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (blog *Blog) BeforeCreate(tx *gorm.DB) (err error) {
    if blog.ID == uuid.Nil {
        blog.ID = uuid.New()
    }
    return
}
