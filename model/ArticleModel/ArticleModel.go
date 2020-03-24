package ArticleModel

import (
	"quick_gin/config/db"
	"quick_gin/model"
)

type Article struct {
	model.Base
	Uid      uint   `json:"uid"`
	Title    string `gorm:"size:30" json:"title"`
	Content  string `json:"content" gorm:"type:longtext"`
	Favorite uint   `json:"favorite" gorm:"default:0"`
}

func (Article) TableName() string {
	return "article"
}

type Articles []Article

func (article *Article) Add() error {

	if err := db.Instance().Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (article *Article) Detail(id int) error {
	if err := db.Instance().First(article, id).Error; err != nil {
		return err
	}
	return nil
}
