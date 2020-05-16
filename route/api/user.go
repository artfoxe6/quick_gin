package api

import (
	"github.com/gin-gonic/gin"
	"quick_gin/middleware"
	"quick_gin/service/user"
	"quick_gin/util/request"
)

// 此函数需要在 route.go 调用
func LoadUserRoute(r *gin.Engine) {
	g := r.Group("/user")

	/**
	 * @api {post} /user/add 添加用户
	 * @apiGroup User
	 *
	 * @apiParam {string} username 用户名
	 * @apiParam {string} age 年龄
	 * @apiParam {integer} password 登录密码
	 *
	 * @apiSuccessExample Success-Response:
	 * {
	 *     "data": null,
	 *     "message": "",
	 *     "statusCode": 200
	 * }
	 */
	g.POST("/add", func(context *gin.Context) {
		user.Add(request.New(context))
	})

	/**
	 * @api {get} /user/list 用户列表
	 * @apiGroup User
	 *
	 * @apiSuccessExample Success-Response:
	 * {
	 *     "data": [
	 *         {
	 *             "id": 1,
	 *             "created_at": "2020-05-16T17:44:51+08:00",
	 *             "updated_at": "2020-05-16T17:44:51+08:00",
	 *             "deleted_at": null,
	 *             "user_name": "张三2",
	 *             "age": 19,
	 *             "password": "$2a$04$O13ojGm.yz0NPrRSuRpkhuEVXvR13S501SchwQUvHMAjbOEVC1U9e",
	 *             "last_login_at": null,
	 *             "articles": null
	 *         }
	 *     ],
	 *     "message": "",
	 *     "statusCode": 200
	 * }
	 */
	g.GET("/list", func(context *gin.Context) {
		user.List(request.New(context))
	})

	/**
	 * @api {get} /user/list_with_article 用户列表以及发表的文章
	 * @apiGroup User
	 *
	 * @apiSuccessExample Success-Response:
	 * {
	 *     "data": [
	 *         {
	 *             "age": 19,
	 *             "articles": [
	 *                 {
	 *                     "content": "内容2",
	 *                     "id": 1,
	 *                     "title": "标题2"
	 *                 }
	 *             ],
	 *             "id": 1,
	 *             "last_login_at": null,
	 *             "user_name": "张三2"
	 *         }
	 *     ],
	 *     "message": "",
	 *     "statusCode": 200
	 * }
	 */
	g.GET("/list_with_article", func(c *gin.Context) {
		user.ListWithArticles(request.New(c))
	})

	/**
	 * @api {get} /user/token 获取token
	 * @apiGroup User
	 *
	 * @apiSuccessExample Success-Response:
	 * {
	 *     "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIn0.NgowRjbanOQa3-B5-q6JjxCVjT9dmn1AeeercH1zGSU",
	 *     "message": "",
	 *     "statusCode": 200
	 * }
	 */
	g.GET("/token", func(context *gin.Context) {
		user.CreateToken(request.New(context))
	})

	//需要token认证才能请求的路由
	auth := r.Group("/user/info")
	auth.Use(middleware.Auth())
	{
		/**
		 * @api {get} /user/info 用户详情
		 * @apiGroup User
		 *
		 * @apiSuccessExample Success-Response:
		 * {
		 *     "data": {
		 *         "age": 19,
		 *         "last_login_at": null,
		 *         "user_name": "张三2"
		 *     },
		 *     "message": "",
		 *     "statusCode": 200
		 * }
		 */
		auth.GET("/", func(context *gin.Context) {
			user.Info(request.New(context))
		})
	}

}
