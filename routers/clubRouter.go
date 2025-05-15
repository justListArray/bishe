package routers

import (
	"footballsys/controllers/club"

	"footballsys/models"

	"github.com/gin-gonic/gin"
)

func ClubRouterInit(r *gin.Engine) {
	ClubRouters := r.Group("/club")

	//r.Use(controllers.MiddleController{}.AuthMiddleware())
	models.DB.AutoMigrate(&models.Member{})
	{
		//ClubRouters.GET("/login", club.clubController{}.Login)
		ClubRouters.GET("/delete", club.ClubController{}.DeleteMember1)
		ClubRouters.POST("/delete", club.ClubController{}.DeleteMember)
		ClubRouters.GET("/add", club.ClubController{}.AddMember)
		ClubRouters.GET("/search", club.ClubController{}.SearchMember1)
		ClubRouters.POST("/search", club.ClubController{}.SearchMember)
		ClubRouters.GET("/addTeam", club.ClubController{}.AddTeamPlayer1)
		ClubRouters.POST("/addTeam", club.ClubController{}.AddTeamPlayer)
		ClubRouters.GET("/index", club.ClubController{}.Index)
		ClubRouters.POST("/save", club.ClubController{}.SaveMember)
		ClubRouters.GET("/all", (&club.ClubController{}).SearchAllMember)

		//ClubRouters.GET("/analyse", club.ClubController{}.AnalysePlayer)
	}
}
