package article

import (
	"quick_gin/model/ArticleModel"
	"quick_gin/util/lib"
	"quick_gin/util/request"
)

func Add(r *request.Request) bool {
	err := r.Validate([]string{"title", "content", "uid"})
	if err != nil {
		return r.Error(err.Error())
	}
	inputs := r.Posts()
	articleModel := ArticleModel.Article{
		Uid:      uint(lib.Int(inputs["uid"])),
		Title:    inputs["title"],
		Content:  inputs["content"],
		Favorite: 0,
	}
	err = articleModel.Add()
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(nil)
}

func Detail(r *request.Request) bool {
	err := r.Validate([]string{"id"})
	if err != nil {
		return r.Error(err.Error())
	}
	article := new(ArticleModel.Article)
	id := lib.Int(r.Id())
	err = article.Detail(id)
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(article)
}
