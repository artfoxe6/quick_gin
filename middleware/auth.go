package middleware

import (
	"github.com/gin-gonic/gin"
	"quick_gin/util/request"
	"quick_gin/util/token"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := request.New(c)

		t := r.Header("Authorization", "")
		if t == "" {
			r.Error("认证失败")
			r.Abort()
		}
		err := token.VerifyJwtToken(t)
		if err != nil {
			r.Error("认证失败")
			r.Abort()
		}
		//fmt.Println("middleware before")
		r.Next()
		//fmt.Println("middleware after")
	}
}
