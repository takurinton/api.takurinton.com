package service

import (
	"portfolio/model"
)

type Analytics struct{}

func (Analytics) AddAnalytics(analytics model.Analytics) (err error) {
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	master := model.Master{}
	if err = db.Table("domain").Select("id").Where("name = ?", analytics.Domain).Find(&master).Error; err != nil {
		master.Id = 6
	}

	db.Exec("INSERT INTO analytics(domain, path, ua) VALUES (?, ?, ?);", master.Id, analytics.Path, analytics.UA)

	return
}
