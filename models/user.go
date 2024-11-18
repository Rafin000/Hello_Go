package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    ID                 uuid.UUID          `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Username           string             `json:"username" gorm:"size:255;not null;unique"`
    FullName           string             `json:"full_name" gorm:"size:255;not null"`
    Designation        string             `json:"designation" gorm:"size:255;not null"`
    About              string             `json:"about"`
    CVLink             string             `json:"cv_link"`
    ProfilePictureLink string             `json:"profile_picture_link"`
    SocialMediaLinks   []SocialMediaLink  `json:"social_media_links" gorm:"foreignKey:UserID"`
    Testimonials       []Testimonial      `json:"testimonials" gorm:"foreignKey:UserID"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
    user.ID = uuid.New()
    return
}
