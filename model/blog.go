package model

import "time"

type Posts struct {
	Category interface{} `json:"category"`
	Current  int         `json:"current"`
	First    int         `json:"first"`
	Last     int         `json:"last"`
	Next     int         `json:"next"`
	PageSize int         `json:"page_size"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID               int         `json:"id"`
		Title            string      `json:"title"`
		Contents         string      `json:"contents"`
		ContentsImageURL string      `json:"contents_image_url"`
		Category         string      `json:"category"`
		IsOpen           bool        `json:"is_open"`
		PubDate          time.Time   `json:"pub_date"`
		Comment          interface{} `json:"comment"`
	} `json:"results"`
	Total int `json:"total"`
}

type BlogappPost struct {
	Id               int64            `gorm:"primary_key" json:"id"`
	Title            string           `json:"title"`
	Contents         string           `json:"contents"`
	ContentsImageUrl string           `json:"contents_image_url"`
	Name             string           `json:"category"`
	Open             bool             `json:"is_open"`
	PubDate          time.Time        `json:"pub_date"`
	Comment          []BlogappComment `json:"comment"`
}

type Rss struct {
	Id       int64     `gorm:"primary_key" json:"id"`
	Title    string    `json:"title"`
	Contents string    `json:"contents"`
	PubDate  time.Time `json:"pub_date"`
}

type BlogappCategory struct {
	Id   int64  `gorm:"pk autoincr int(64)" form:"id" json:"id"`
	Name string `gorm:"varchar(100)" json:"name" form:"name"`
}

type BlogappComment struct {
	Contents string    `json:"contents"`
	PubDate  time.Time `json:"pub_date"`
	PostIdId int64     `json:"post_id"`
	Name     string    `json:"name"`
}

type PostData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

type Post struct {
	Title    string    `json:"title"`
	Contents string    `json:"contents"`
	Name     string    `json:"category"`
	Open     bool      `json:"is_open"`
	PubDate  time.Time `json:"pub_date"`
}

type Category struct {
	Name string `json:"name" form:"name"`
}
