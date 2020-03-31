package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	C *gin.Context
}

//通过gin.Context转为Request
func New(c *gin.Context) *Request {
	return &Request{C: c}
}

//获取GET参数
func (r *Request) Gets() map[string]string {
	values := r.C.Request.URL.Query()
	temp := map[string]string{}
	for k, v := range values {
		temp[k] = v[0]
	}
	return temp
}
func (r *Request) Get(k string) string {
	return r.C.Query(k)
}
func (r *Request) DefaultGet(k, v string) string {
	return r.C.DefaultQuery(k, v)
}

//获取POST参数
func (r *Request) Posts() map[string]string {
	err := r.C.Request.ParseForm()
	if err != nil {
		log.Fatalf("%v", err)
	}
	values := r.C.Request.PostForm
	temp := map[string]string{}
	for k, v := range values {
		temp[k] = v[0]
	}
	return temp
}
func (r *Request) Post(k string) string {
	return r.C.PostForm(k)
}
func (r *Request) DefaultPost(k, v string) string {
	return r.C.DefaultPostForm(k, v)
}

//获取指定Headers
func (r *Request) Headers() map[string]string {
	temp := map[string]string{}
	for k, v := range r.C.Request.Header {
		temp[k] = v[0]
	}
	return temp
}
func (r *Request) Header(k string) string {
	return r.C.GetHeader(k)
}

//错误返回请求
func (r *Request) Error(message interface{}) bool {

	r.C.JSON(http.StatusOK, map[string]interface{}{
		"data":       "",
		"message":    message,
		"statusCode": http.StatusBadRequest,
	})
	return false
}

//成功返回请求
func (r *Request) Success(data interface{}) bool {

	r.C.JSON(http.StatusOK, map[string]interface{}{
		"data":       data,
		"message":    "",
		"statusCode": http.StatusOK,
	})
	return true
}

// 仅验证字段是否缺少
func (r *Request) Validate(list []string) error {
	if strings.ToUpper(r.C.Request.Method) == "POST" {
		for i := 0; i < len(list); i++ {
			if _, b := r.C.GetPostForm(list[i]); b == false {
				return errors.New("缺少字段：" + list[i])
			}
		}
		return nil
	}
	for i := 0; i < len(list); i++ {
		if _, b := r.C.GetQuery(list[i]); b == false {
			return errors.New("缺少字段：" + list[i])
		}
	}
	return nil
}

//------------ 常见参数封装 ------------

// 获取id参数
func (r *Request) Id() string {
	return r.DefaultGet("id", "0")
}

//获取请求中的分页信息
func (r *Request) Page() int {
	temp := r.DefaultGet("page", "1")
	page, err := strconv.Atoi(temp)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return page - 1
}

//获取请求中的分页信息
func (r *Request) PerPage() int {
	temp := r.DefaultGet("per_page", "10")
	perPage, err := strconv.Atoi(temp)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return perPage
}

//--------------- end ----------------------
