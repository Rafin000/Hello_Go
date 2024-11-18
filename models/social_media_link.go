package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)
type SocialMediaLink struct {
    ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
    Platform string    `json:"platform" gorm:"size:255;not null"`
    URL      string    `json:"url" gorm:"size:255;not null"`
}

func (socialMediaLink *SocialMediaLink) BeforeCreate(tx *gorm.DB) (err error) {
    socialMediaLink.ID = uuid.New()
    return
}
