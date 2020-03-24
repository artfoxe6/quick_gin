package Migrate

import (
	"quick_gin/config/db"
	"quick_gin/model/ArticleModel"
	"quick_gin/model/UserModel"
)

func Run() {
	db.Instance().AutoMigrate(
		&UserModel.User{},
		&ArticleModel.Article{},
	)
}
