package api

import (
	"github.com/gin-gonic/gin"
	"quick_gin/route/pagecache"
	"quick_gin/service/article"
	"quick_gin/util/request"
)

func LoadArticleRoute(r *gin.Engine) {

	/**
	 * @api {post} /article/add 添加文章
	 * @apiGroup Article
	 *
	 * @apiParam {string} title 文章标题
	 * @apiParam {string} content 文章内容
	 * @apiParam {integer} uid 用户
	 *
	 * @apiSuccessExample Success-Response:
	 * {
	 *     "data": null,
	 *     "message": "",
	 *     "statusCode": 200
	 * }
	 */
	r.POST("/article/add", func(c *gin.Context) {
		article.Add(request.New(c))
	})

	/**
	 * @api {get} /article/detail 文章详情
	 * @apiGroup Article
	 *
	 * @apiParam {integer} id 文章id
	 *
	 * @apiSuccessExample Success-Response:
	 * {
	 *     "data": {
	 *         "id": 1,
	 *         "created_at": "2020-03-24T09:20:05+08:00",
	 *         "updated_at": "2020-03-24T09:20:05+08:00",
	 *         "deleted_at": null,
	 *         "uid": 2,
	 *         "title": "标题2",
	 *         "content": "内容2",
	 *         "favorite": 0
	 *     },
	 *     "message": "",
	 *     "statusCode": 200
	 * }
	 */
	r.GET("/article/detail", pagecache.Second(300, func(c *gin.Context) {
		article.Detail(request.New(c))
	}))

}
