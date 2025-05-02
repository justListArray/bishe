package player

import (
	"fmt"
	"footballsys/controllers/base"
	"footballsys/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	base.BaseController
}

func (player PlayerController) trans(c *gin.Context, name string) (id int) {

	url := "https://v3.football.api-sports.io/players/profiles?search=" + name
	response := base.BaseController.PreOperation(player.BaseController, c, url)
	responseData := response["response"].([]interface{})
	if len(response) == 0 {
		fmt.Println("No fixtures found for the specified team and season.")
		return
	}
	for _, PlayerData := range responseData {
		playerArray := PlayerData.(map[string]interface{})
		playerID := int(playerArray["player"].(map[string]interface{})["id"].(float64))
		id = int(playerID)
	}
	return id
}

func (player PlayerController) AddPlayerInfo(c *gin.Context) {
	var Player models.Player
	season := "2023"
	name := c.Query("name")
	id := strconv.Itoa(player.trans(c, name))
	url := "https://v3.football.api-sports.io/players?id=" + id + "&season=" + season
	response := base.BaseController.PreOperation(player.BaseController, c, url)

	responseData := response["response"].([]interface{})
	if len(response) == 0 {
		fmt.Println("No fixtures found for the specified team and season.")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Player":       Player,
		"responseData": responseData,
	})

	// if err := c.ShouldBindJSON(&Player); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

}
