package UserModel

import (
	"quick_gin/config/db"
	"quick_gin/model"
	"quick_gin/model/ArticleModel"
	"time"
)

type User struct {
	model.Base
	UserName    string                `json:"user_name" gorm:"size:30"`
	Age         uint8                 `json:"age" gorm:"type:tinyint"`
	Password    string                `json:"password" gotm:"size:100"`
	LastLoginAt *time.Time            `json:"last_login_at"`
	Articles    ArticleModel.Articles `gorm:"foreignkey:Uid" json:"articles"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) Add() error {
	if err := db.Instance().Create(user).Error; err != nil {
		return err
	}
	return nil
}

type Users []User

func (users *Users) List() error {
	if err := db.Instance().Find(users).Error; err != nil {
		return err
	}
	return nil
}
func (users *Users) ListWithArticles() error {
	if err := db.Instance().Preload("Articles").Find(users).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) Info(id int) error {
	if err := db.Instance().Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
