package service

import "portfolio/model"

type Marina struct{}

func (m Marina) GetOmedetou() (o []*model.MarinaMarinaOmedetou, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Table("marina_marinaomedetou").Select("id, name, sentence").Order("id desc").Find(&o).Error; err != nil {
		return
	}

	return
}

func (m Marina) GetCongrats() (o []*model.MarinaMarinaOmedetouCount, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Table("marina_marinaomedetoucount").Select("sum(count) as count").Find(&o).Error; err != nil {
		return
	}

	return
}

func (m Marina) PostOmedetou(name, sentence string) (err error) {
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Exec("INSERT INTO marina_marinaomedetou(name, sentence) VALUES (?, ?);", name, sentence)
	return nil
}

func (m Marina) PostCongrats(count int64) (err error) {
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Exec("INSERT INTO marina_marinaomedetoucount(count) VALUES (?);", count)
	return nil
}
