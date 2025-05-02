package routers

import (
	"footballsys/controllers/club"
	"footballsys/models"

	"github.com/gin-gonic/gin"
)

func ClubRouterInit(r *gin.Engine) {
	ClubRouters := r.Group("/club")
	models.DB.AutoMigrate(&models.Member{})
	{
		//ClubRouters.GET("/login", club.clubController{}.Login)
		ClubRouters.GET("/delete", club.ClubController{}.DeleteMember)
		ClubRouters.GET("/add", club.ClubController{}.AddMember)
		ClubRouters.GET("/search", club.ClubController{}.SearchMember)
		ClubRouters.GET("/addTeam", club.ClubController{}.AddTeamPlayer)
		//ClubRouters.GET("/analyse", club.ClubController{}.AnalysePlayer)
	}
}
