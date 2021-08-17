package model

import "time"

type DiaryreportPost struct {
	Id       int64     `gorm:"primary_key" json:"id"`
	Contents string    `json:"contents"`
	PubDate  time.Time `json:"pub_date"`
}
