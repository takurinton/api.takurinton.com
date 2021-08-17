package model

import "time"

type PortfolioMineOneHandledPlusOneContent struct {
	Id       int64     `gorm:"primary_key" json:"id"`
	Title    string    `json:"title"`
	Contents time.Time `json:"contents"`
}
