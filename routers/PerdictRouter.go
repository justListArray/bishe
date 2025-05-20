package routers

import (
	"footballsys/controllers/match"

	"github.com/gin-gonic/gin"
)

func PredictionRouterInit(r *gin.Engine) {
	PredictionRouters := r.Group("/pre")
	//PredictionRouters.Use(models.CheckToken)
	{
		PredictionRouters.GET("/predict", match.PerdictController{}.Prediction1) //测试
		PredictionRouters.POST("/predict", match.PerdictController{}.Prediction)
	}
	//r.Use(controllers.MiddleController{}.AuthMiddleware())

}
