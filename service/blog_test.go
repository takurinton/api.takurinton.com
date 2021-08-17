package service

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

// もくもくもっく！！
func Mock() (*gorm.DB, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open("mysql", database)
	if err != nil {
		return nil, nil, err
	}
	return db, mock, nil
}

func TestGetAllPost(t *testing.T) {
	//
}

func TestGetDetailPost(t *testing.T) {
	//
}

func TestGetComment(t *testing.T) {
	//
}

func TestPostComment(t *testing.T) {
	//
}
