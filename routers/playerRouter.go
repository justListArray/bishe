package routers

import (
	"footballsys/controllers/player"

	"github.com/gin-gonic/gin"
)

func PlayerRouterInit(r *gin.Engine) {
	r.GET("/play", player.PlayerController{}.AddPlayerInfo) //测试
}
