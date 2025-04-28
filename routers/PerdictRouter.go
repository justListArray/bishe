package routers

import (
	"footballsys/controllers/match"

	"github.com/gin-gonic/gin"
)

func PredictionRouterInit(r *gin.Engine) {
	r.GET("/perdict", match.PerdictController{}.Prediction) //测试
}
