package service

import "portfolio/model"

type Get101Content struct{}

func (Get101Content) Get101Content() (p []model.PortfolioMineOneHandledPlusOneContent, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Select("id, title").Table("portfolio_mineonehandledplusonecontent").Find(&p).Error; err != nil {
		return
	}

	return
}
