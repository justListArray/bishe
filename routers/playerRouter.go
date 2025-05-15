package routers

import (
	"footballsys/controllers/player"
	"footballsys/models"

	"github.com/gin-gonic/gin"
)

func PlayerRouterInit(r *gin.Engine) {
	//r.Use(controllers.MiddleController{}.AuthMiddleware())
	PlayerRouters := r.Group("/play")
	models.DB.AutoMigrate(&models.Player{})
	{
		PlayerRouters.GET("/index", player.PlayerController{}.Index)
		PlayerRouters.GET("/play", player.PlayerController{}.AddPlayerInfo1) //测试
		PlayerRouters.POST("/play", player.PlayerController{}.AddPlayerInfo) //测试
		// PlayerRouters.GET("/analynse", player.PlayerController{}.AnalysePlayer1)
		PlayerRouters.GET("/analynse", player.PlayerController{}.AnalysePlayer)

	}
}
