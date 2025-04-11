package routers

import (
	"footballsys/controllers/club"

	"github.com/gin-gonic/gin"
)

func ClubRouterInit(r *gin.Engine) {
	ClubRouters := r.Group("/club")
	{
		//ClubRouters.GET("/login", club.clubController{}.Login)
		ClubRouters.GET("/delete", club.ClubController{}.DeleteMember)
		ClubRouters.GET("/add", club.ClubController{}.AddMember)
		ClubRouters.GET("/search", club.ClubController{}.SearchMember)
	}
}
