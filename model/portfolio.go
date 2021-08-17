package model

import "time"

type PortfolioMadeproduction struct {
	Id          int64  `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Explanation string `json:"explanation"`
}

type PortfolioSkill struct {
	Id   int64  `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

type PortfolioInternship struct {
	Id          int64  `gorm:"primary_key" json:"id"`
	CompanyName string `json:"company_name"`
	Overview    string `json:"overview"`
	Period      string `json:"period"`
}

type PortfolioMine struct {
	Content string `json:"content"`
}

// 大人の事情でここに置く、許して
type BlogappContact struct {
	Id       int64     `gorm:"primary_key" json:"id"`
	Name     string    `json:"name"`
	Mail     string    `json:"mail"`
	Contents string    `json:"contents"`
	PubDate  time.Time `json:"pub_date"`
}
