package routers

import (
	"footballsys/controllers/train"

	"github.com/gin-gonic/gin"
)

func TrainRouterInit(r *gin.Engine) {
	TrainRouters := r.Group("/train")
	{
		//ClubRouters.GET("/login", club.clubController{}.Login)
		TrainRouters.GET("/add", train.TrainController{}.TrainAdd)
		TrainRouters.GET("/search", train.TrainController{}.TrainSearch)
	}
}
