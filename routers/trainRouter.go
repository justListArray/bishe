package routers

import (
	"footballsys/controllers/train"
	"footballsys/models"

	"github.com/gin-gonic/gin"
)

func TrainRouterInit(r *gin.Engine) {
	TrainRouters := r.Group("/train")
	models.DB.AutoMigrate(&models.Train{})
	{
		//ClubRouters.GET("/login", club.clubController{}.Login)
		TrainRouters.GET("/index", train.TrainController{}.Index)
		TrainRouters.GET("/add", train.TrainController{}.AddTraibInfo)
		TrainRouters.POST("/add", train.TrainController{}.TrainAdd)
		TrainRouters.GET("/search", train.TrainController{}.TrainSearch1)
		TrainRouters.POST("/search", train.TrainController{}.TrainSearch)
	}
}
