package models

import (
    "github.com/google/uuid"
	"gorm.io/gorm"
    "time"
)

type Testimonial struct {
    ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
    Name      string    `json:"name" gorm:"size:255"`
    Position  string    `json:"position" gorm:"size:255"`
    Message   string    `json:"message" gorm:"type:text"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (testimonial *Testimonial) BeforeCreate(tx *gorm.DB) (err error) {
    testimonial.ID = uuid.New()
    return
}
