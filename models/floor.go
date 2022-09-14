package models

import "time"

type Floor struct {
	ID       int    `json:"id"`
	Like     int    `json:"like"`
	LikeData string `json:"like_data"`
	History  string `json:"history"`
}

type FloorLike struct {
	FloorID  int  `json:"floor_id" gorm:"primarykey"`
	UserID   int  `json:"user_id" gorm:"primarykey"`
	LikeData int8 `json:"like_data"`
}

type FloorHistory struct {
	BaseModel
	Content string `json:"content"`
	Reason  string `json:"reason"`
	FloorID int    `json:"floor_id"`
	UserID  int    `json:"user_id"` // The one who modified the floor
}

type FloorHistoryOld struct {
	Content     string    `json:"content"`
	AlteredBy   int       `json:"altered_by"`
	AlteredTime time.Time `json:"altered_time"`
}
