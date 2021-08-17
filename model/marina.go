package model

type MarinaMarinaOmedetou struct {
	Id       int64  `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Sentence string `json:"sentence"`
}

type MarinaMarinaOmedetouCount struct {
	Count int64 `json:"count"`
}
