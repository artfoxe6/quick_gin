package user

import (
	"golang.org/x/crypto/bcrypt"
	"quick_gin/model/UserModel"
	"quick_gin/util/lib"
	"quick_gin/util/request"
	"quick_gin/util/token"
)

//添加用户
func Add(r *request.Request) bool {

	if err := r.Validate([]string{"user_name", "age", "password"}); err != nil {
		return r.Error(err.Error())
	}
	inputs := r.Posts()
	password, err := bcrypt.GenerateFromPassword([]byte(inputs["password"]), bcrypt.MinCost)
	if err != nil {
		return r.Error(err.Error())
	}
	user := UserModel.User{
		UserName: inputs["user_name"],
		Age:      uint8(lib.Int(inputs["age"])),
		Password: string(password),
	}
	err = user.Add()
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(nil)
}

//用户列表
func List(r *request.Request) bool {

	users := new(UserModel.Users)
	err := users.List()
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(users)

}

//用户列表以及用户发布的文章
func ListWithArticles(r *request.Request) bool {

	users := new(UserModel.Users)
	err := users.ListWithArticles()
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(users.ToArray())
}

//用户信息
func Info(r *request.Request) bool {
	if err := r.Validate([]string{"id"}); err != nil {
		return r.Error(err.Error())
	}
	user := new(UserModel.User)
	id := lib.Int(r.Id())
	err := user.Info(id)
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(user.Source())
}

//获取token，仅做测试用，正常情况下是通过登录接口返回
func CreateToken(r *request.Request) bool {
	if err := r.Validate([]string{"uid"}); err != nil {
		return r.Error(err.Error())
	}
	jwtToken, err := token.CreateJwtToken(map[string]interface{}{
		"uid": r.Get("uid"),
	})
	if err != nil {
		return r.Error(err.Error())
	}
	return r.Success(jwtToken)

}
