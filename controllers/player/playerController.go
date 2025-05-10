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
func NilOperation(num interface{}) (ret int) {
	if num == nil {
		return 0
	} else {
		return int(num.(float64))
	}
}

func (player PlayerController) AddPlayerInfo1(c *gin.Context) {

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

	for _, playerData := range responseData {
		playerArray := playerData.(map[string]interface{})
		Player.Id = int(playerArray["player"].(map[string]interface{})["id"].(float64))
		Player.Name = playerArray["player"].(map[string]interface{})["name"].(string)
		playerinfo := playerArray["statistics"].([]interface{})
		//int(
		goals := playerinfo[0].(map[string]interface{})["goals"].(map[string]interface{})["total"] //.(float64))
		Player.Goal = NilOperation(goals)
		passes := playerinfo[0].(map[string]interface{})["passes"].(map[string]interface{})["total"] //xiu
		Player.Pass = NilOperation(passes)
		tackles := playerinfo[0].(map[string]interface{})["tackles"].(map[string]interface{})["total"]
		Player.Tackle = NilOperation(tackles)
		fouls := playerinfo[0].(map[string]interface{})["fouls"].(map[string]interface{})["total"]
		Player.Foul = NilOperation(fouls)
		yellowcards := playerinfo[0].(map[string]interface{})["cards"].(map[string]interface{})["yellow"]
		Player.Yellowcard = NilOperation(yellowcards)
		redcard := playerinfo[0].(map[string]interface{})["cards"].(map[string]interface{})["red"]
		Player.Redcard = NilOperation(redcard)
	}
	models.DB.Create(&Player)

	// if err := c.ShouldBindJSON(&Player); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"Player":       Player,
		"responseData": responseData,
	})

}

func (player1 PlayerController) AnalysePlayer(c *gin.Context) {
	var players []models.Player
	if err := models.DB.Find(&players).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch player data from database"})
		return
	}
	var totalGoals, totalPasses, totalTackles, totalFouls, totalYellowCards, totalRedCards int
	var maxGoals, maxPasses, maxTackles, maxFouls, maxYellowCards, maxRedCards int
	var bestGoalScorer, bestPasser, bestTackler, bestFouler, mostYellowCards, mostRedCards models.Player

	for _, player := range players {
		totalGoals += player.Goal
		totalPasses += player.Pass
		totalTackles += player.Tackle
		totalFouls += player.Foul
		totalYellowCards += player.Yellowcard
		totalRedCards += player.Redcard

		if player.Goal > maxGoals {
			maxGoals = player.Goal
			bestGoalScorer = player
		}
		if player.Pass > maxPasses {
			maxPasses = player.Pass
			bestPasser = player
		}
		if player.Tackle > maxTackles {
			maxTackles = player.Tackle
			bestTackler = player
		}
		if player.Foul > maxFouls {
			maxFouls = player.Foul
			bestFouler = player
		}
		if player.Yellowcard > maxYellowCards {
			maxYellowCards = player.Yellowcard
			mostYellowCards = player
		}
		if player.Redcard > maxRedCards {
			maxRedCards = player.Redcard
			mostRedCards = player
		}
	}

	fmt.Println("\nData Analysis:")
	fmt.Printf("Total Goals: %d\n", totalGoals)
	fmt.Printf("Total Passes: %d\n", totalPasses)
	fmt.Printf("Total Tackles: %d\n", totalTackles)
	fmt.Printf("Total Fouls: %d\n", totalFouls)
	fmt.Printf("Total Yellow Cards: %d\n", totalYellowCards)
	fmt.Printf("Total Red Cards: %d\n", totalRedCards)

	fmt.Printf("\nBest Goal Scorer: %s with %d goals\n", bestGoalScorer.Name, bestGoalScorer.Goal)
	fmt.Printf("Best Passer: %s with %d passes\n", bestPasser.Name, bestPasser.Pass)
	fmt.Printf("Best Tackler: %s with %d tackles\n", bestTackler.Name, bestTackler.Tackle)
	fmt.Printf("Most Fouls: %s with %d fouls\n", bestFouler.Name, bestFouler.Foul)
	fmt.Printf("Most Yellow Cards: %s with %d yellow cards\n", mostYellowCards.Name, mostYellowCards.Yellowcard)
	fmt.Printf("Most Red Cards: %s with %d red cards\n", mostRedCards.Name, mostRedCards.Redcard)
}
