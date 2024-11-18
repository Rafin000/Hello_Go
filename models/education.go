package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Education struct {
    ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
    Year       int       `json:"year" gorm:"not null"`
    Degree     string    `json:"degree" gorm:"size:255;not null"`
    University string    `json:"university" gorm:"size:255;not null"`
    CGPA       float64   `json:"cgpa" gorm:"not null"`
}

func (education *Education) BeforeCreate(tx *gorm.DB) (err error) {
    education.ID = uuid.New()
    return
}
