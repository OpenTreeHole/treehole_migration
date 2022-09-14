package models

import "time"

type Map = map[string]interface{}

type BaseModel struct {
	ID        int       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"time_created"`
	UpdatedAt time.Time `json:"time_updated"`
}
