package model

import "time"

type Analytics struct {
	Id        int64     `gorm:"primary_key" json:"id"`
	Domain    string    `json:"title"`
	UA        string    `json:"ua"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}

type Master struct {
	Id   int64  `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
