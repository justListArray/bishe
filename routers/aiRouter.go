package routers

import (
	ai "footballsys/controllers/AI"

	"github.com/gin-gonic/gin"
)

func AIRouterInit(r *gin.Engine) {
	r.GET("/ai", ai.UseKimi1) //登录
	r.POST("/ai", ai.UseKimi) //登录
}
