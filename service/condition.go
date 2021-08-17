package service

import (
	"portfolio/model"
	"time"

	"github.com/jinzhu/gorm"
)

type MaririntonCondition struct{}

func CreateCondition(user string, mental, physical float64, datetime time.Time) (string, error) {
	db, err := DBConn()
	if err != nil {
		return "", err
	}
	defer db.Close()

	d := datetime.Format("2006-01-02")

	var condition model.Condition
	err = db.Table("maririnton_condition").Where("username = ?", user).Where("DATE_FORMAT(created_at, '%Y-%m-%d') = date(?)", d).Find(&condition).Error

	// record not found error の時の対応
	// これが来るのが正常な処理、ほんとはエラーによって分岐したい
	if err != nil {
		db.Exec("INSERT INTO maririnton_condition(username, mental, physical, created_at) VALUES (?, ?, ?, ?);", user, mental, physical, datetime)
		return "created", nil
	} else {
		id := condition.Id
		db.Exec("UPDATE maririnton_condition SET mental = ?, physical = ?, created_at = ? WHERE id = ?;", mental, physical, datetime, id)
		return "updated", nil
	}
}

func (m MaririntonCondition) GetCondition(username string) (res []model.Condition, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	var query *gorm.DB
	if username == "" {
		query = db.Table("maririnton_condition").Find(&res)
	} else {
		// query = db.Exec("select * from maririnton_condition where username='U01TVTSN1TQ';")
		query = db.Table("maririnton_condition").Where("username = ?", username).Find(&res)
	}
	if err = query.Error; err != nil {
		return
	}
	return
}
