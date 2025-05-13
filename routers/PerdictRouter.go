package routers

import (
	"footballsys/controllers/match"

	"github.com/gin-gonic/gin"
)

func PredictionRouterInit(r *gin.Engine) {
	//r.Use(controllers.MiddleController{}.AuthMiddleware())
	r.GET("/predict", match.PerdictController{}.Prediction1) //测试
	r.POST("/predict", match.PerdictController{}.Prediction)
}
