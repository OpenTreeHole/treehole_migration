package models

type Hole struct {
	ID      int `json:"id"`
	Mapping string
}

type AnonynameMapping struct {
	HoleID    int    `json:"hole_id" gorm:"primarykey"`
	UserID    int    `json:"user_id" gorm:"primarykey"`
	Anonyname string `json:"anonyname" gorm:"index;size:32"`
}
