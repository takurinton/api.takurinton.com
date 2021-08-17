package model

import "time"

type User struct {
	Id          int64     `gorm:"primary_key" json:"id"`
	Username    string    `json:"username"`
	Password    []byte    `json:"password"`
	Email       string    `json:"email"`
	IsActive    bool      `json:"is_active"`
	IsStaff     bool      `json:"is_staff"`
	IsSuperuser bool      `json:"is_superuser"`
	DateJoined  time.Time `json:"date_joined"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
