package ArticleService

import (
	"quick_gin/model/ArticleModel"
	"quick_gin/util/lib"
	"quick_gin/util/request"
)

func Add(r *request.Request) {

	err := r.Validate([]string{"title", "content", "uid"})

	if err != nil {
		r.Error(err.Error())
		return
	}

	inputs := r.Posts()

	//articleModel := new(ArticleModel.Article)

	articleModel := ArticleModel.Article{
		Uid:      uint(lib.Int(inputs["uid"])),
		Title:    inputs["title"],
		Content:  inputs["content"],
		Favorite: 0,
	}

	err = articleModel.Add()

	if err != nil {
		r.Error(err.Error())
		return
	}
	r.Success(nil)
}
