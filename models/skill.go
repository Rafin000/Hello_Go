package models

type Skill struct {
    Name string `json:"name" gorm:"size:255;not null"`
    Icon string `json:"icon,omitempty" gorm:"size:255"`
}
