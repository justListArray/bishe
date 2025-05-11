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
		IndexRouters.GET("/login", index.IndexController{}.Login1) //登录
		IndexRouters.POST("login", index.IndexController{}.Login)
		IndexRouters.GET("/signin", index.IndexController{}.Signin1) //注册
		IndexRouters.POST("/signin", index.IndexController{}.Signin)
		IndexRouters.GET("/index", index.IndexController{}.Index) //测试

	}
}
