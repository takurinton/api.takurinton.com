package service

import (
	"portfolio/model"
	"time"
)

type Portfolio struct{}

type Props struct {
	Intern []model.PortfolioInternship
	Skill  []model.PortfolioSkill
	Made   []model.PortfolioMadeproduction
	Mine   model.PortfolioMine
}

func (p *Portfolio) GetPortfolio() (props Props, err error) {
	db, err := DBConn()
	if err != nil {
		return
	}
	defer db.Close()

	intern := []model.PortfolioInternship{}
	skill := []model.PortfolioSkill{}
	made := []model.PortfolioMadeproduction{}
	mine := model.PortfolioMine{}

	if err = db.Table("portfolio_internship").Order("id desc").Find(&intern).Error; err != nil {
		return
	}
	if err = db.Table("portfolio_skill").Find(&skill).Error; err != nil {
		return
	}
	if err = db.Table("portfolio_madeproduction").Find(&made).Error; err != nil {
		return
	}
	if err = db.Table("portfolio_mine").Find(&mine).Error; err != nil {
		return
	}

	props = Props{
		Intern: intern,
		Skill:  skill,
		Made:   made,
		Mine:   mine,
	}

	return
}

func (p Portfolio) PostContact(name, mail, contents string) error {
	t := time.Now()
	db, err := DBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	c := model.BlogappContact{Name: name, Mail: mail, Contents: contents, PubDate: t}
	if err := db.Table("blogapp_contact").Create(&c).Error; err != nil {
		return err
	}

	return nil
}
