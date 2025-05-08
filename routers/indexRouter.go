package routers

import (
	"footballsys/controllers/index"
	"footballsys/models"

	"github.com/gin-gonic/gin"
)

func IndexRouterInit(r *gin.Engine) {
	IndexRouters := r.Group("/index")
	models.DB.AutoMigrate(&models.User{})
	{
		IndexRouters.GET("/login", index.IndexController{}.Index1) //登录
		IndexRouters.POST("login", index.IndexController{}.Login)
		IndexRouters.GET("/signin", index.IndexController{}.Signin) //注册
		IndexRouters.GET("/test", index.IndexController{}.Test)     //测试
	}
}
