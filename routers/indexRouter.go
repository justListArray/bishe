package routers

import (
	"footballsys/controllers/index"

	"github.com/gin-gonic/gin"
)

func IndexRouterInit(r *gin.Engine) {
	IndexRouters := r.Group("/index")
	{
		IndexRouters.GET("/login", index.IndexController{}.Index1) //登录
		IndexRouters.POST("login", index.IndexController{}.Login)
		IndexRouters.GET("/signin", index.IndexController{}.Signin) //注册
		IndexRouters.GET("/test", index.IndexController{}.Success)  //测试
	}
}
