package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Campaign struct {
	gorm.Model

	SenderID        uint   `gorm:"not null"`
	Title           string `gorm:"not null"`
	FormSender      string `gorm:"not null"`
	Subject         string `gorm:"not null"`
	EmailPreheader  string `gorm:"not null"`
	Content         string `gorm:"not null"`
	HTMLTitle       string `gorm:"not null"`
	HTMLDescription string `gorm:"not null"`
}

type CampaignService interface {
	Create(*Campaign) error

	ByID(uint) (*Campaign, error)

	Update(*Campaign) error

	Delete(uint) error
}

type CampaignGORM struct {
	*gorm.DB
}

func NewCampaignGORM(db *gorm.DB) *CampaignGORM {
	return &CampaignGORM{db}
}

func (s CampaignGORM) Create(c *Campaign) error {
	r := s.DB.Create(c)
	return r.Error
}

func (s CampaignGORM) byQuery(q *gorm.DB) (*Campaign, error) {
	e := Campaign{}

	err := q.First(&e).Error
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (s CampaignGORM) ByID(id uint) (*Campaign, error) {
	return s.byQuery(s.DB.Where("id = ?", id))
}

func (s CampaignGORM) Update(c *Campaign) error {
	r := s.DB.Save(c)
	return r.Error
}

func (s CampaignGORM) Delete(id uint) error {
	e := Campaign{
		Model: gorm.Model{ID: id},
	}
	r := s.DB.Delete(e)
	return r.Error
}

func (s CampaignGORM) DestructiveReset() {
	s.DropTableIfExists(&Campaign{})
	s.AutoMigrate()
}

func (s CampaignGORM) AutoMigrate() {
	s.DB.AutoMigrate(&Campaign{})
}
