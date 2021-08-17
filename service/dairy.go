package service

import (
	"portfolio/model"
	"time"

	"github.com/biezhi/gorm-paginator/pagination"
)

type Dairy struct{}

func (Dairy) GetAllDairy(page int) (pn *pagination.Paginator, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	d := []model.DiaryreportPost{}
	q := db.Select("id, pub_date").Table("diaryreport_post").Order("pub_date DESC").Find(&d)
	pn = pagination.Paging(&pagination.Param{
		DB:      q,
		Page:    page,
		Limit:   7,
		ShowSQL: false,
	}, &d)

	return
}

func (Dairy) GetDetailDairy(id string) (d model.DiaryreportPost, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Select("id, pub_date, contents").Table("diaryreport_post").Where("id = ?", id).Find(&d).Error; err != nil {
		return
	}

	return
}

func (Dairy) CreateDailyreport(contents string, pubDate time.Time) (daily model.DiaryreportPost, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	daily = model.DiaryreportPost{Contents: contents, PubDate: pubDate}
	if err = db.Table("diaryreport_post").Create(&daily).Error; err != nil {
		return
	}
	return
}

func (Dairy) UpdateDailyreport(id, contents string) (err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	db.Exec("UPDATE diaryreport_post SET contents = ? WHERE id = ?", contents, id)
	return
}
