package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Experience struct {
    ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
    Year       int       `json:"year" gorm:"not null"`
    Position   string    `json:"position" gorm:"size:255;not null"`
    Company    string    `json:"company" gorm:"size:255;not null"`
    WorkDetails []string `json:"work_details" gorm:"type:text[]"`
}

func (experience *Experience) BeforeCreate(tx *gorm.DB) (err error) {
    experience.ID = uuid.New()  
    return
}
