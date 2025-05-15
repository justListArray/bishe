package player

import (
	"errors"
	"fmt"
	"footballsys/controllers/base"
	"footballsys/models"
	"footballsys/plotfoot"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlayerController struct {
	base.BaseController
}

func (player PlayerController) Trans(name string) (int, error) {

	var existingPlayer models.Member
	if err := models.DB.Where("name = ?", name).First(&existingPlayer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Player with name %s not found in the database", name)
			return 0, fmt.Errorf("player not found")
		}
		log.Printf("Error querying player: %v", err)
		return 0, err
	}
	return existingPlayer.PlayerId, nil

}
func NilOperation(num interface{}) (ret int) {
	if num == nil {
		return 0
	} else {
		return int(num.(float64))
	}
}
func (player PlayerController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "player/index.html", nil)

}
func (player PlayerController) AddPlayerInfo1(c *gin.Context) {
	c.HTML(http.StatusOK, "player/addPlayerInfo.html", nil)

}

func (player PlayerController) AddPlayerInfo(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding query data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(request.Name)
	var Player models.Player
	season := "2023"
	id, err := player.Trans(request.Name) //找不到id
	if err != nil {
		log.Printf("找不到id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id1 := strconv.Itoa(id)
	if id == 0 {
		fmt.Println("没有id")
	}
	log.Printf("%s的id:%d", request.Name, id)
	url := "https://v3.football.api-sports.io/players?id=" + id1 + "&season=" + season
	response := base.BaseController.PreOperation(player.BaseController, c, url)

	responseData := response["response"].([]interface{})
	if len(response) == 0 {
		fmt.Println("No fixtures found for the specified team and season.")
		return
	}
	log.Println(response)

	for _, playerData := range responseData {
		playerArray := playerData.(map[string]interface{})
		Player.Id = id
		log.Println(Player.Id)
		Player.Name = playerArray["player"].(map[string]interface{})["name"].(string)
		log.Println(Player.Name)
		playerinfo := playerArray["statistics"].([]interface{})
		//int(
		goals := playerinfo[0].(map[string]interface{})["goals"].(map[string]interface{})["total"] //.(float64))
		Player.Goal = NilOperation(goals)
		log.Println(Player.Goal)
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
	var existingPlayer models.Player
	if err := models.DB.Where("id = ?", Player.Id).First(&existingPlayer).Error; err == nil {
		// Player exists, update the existing record
		models.DB.Model(&existingPlayer).Updates(Player)
		log.Printf("Updated existing player: %v", Player.Name)
	} else {
		// Player does not exist, create a new record
		models.DB.Create(&Player)
		log.Printf("Created new player: %v", Player.Name)
	}
	log.Println(Player)
	_ = plotfoot.PlotFoot(Player)
	imgPath := filepath.Join("/home/ubuntu/football", Player.Name+"_analynse.png")

	// if err := c.ShouldBindJSON(&Player); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"Player":       Player,
		"responseData": responseData,
		"img":          imgPath,
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
		// log.Printf("球员的信息:%+v", player)

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

	analysisData := struct {
		TotalGoals       int
		TotalPasses      int
		TotalTackles     int
		TotalFouls       int
		TotalYellowCards int
		TotalRedCards    int
		BestGoalScorer   models.Player
		BestPasser       models.Player
		BestTackler      models.Player
		BestFouler       models.Player
		MostYellowCards  models.Player
		MostRedCards     models.Player
	}{
		TotalGoals:       totalGoals,
		TotalPasses:      totalPasses,
		TotalTackles:     totalTackles,
		TotalFouls:       totalFouls,
		TotalYellowCards: totalYellowCards,
		TotalRedCards:    totalRedCards,
		BestGoalScorer:   bestGoalScorer,
		BestPasser:       bestPasser,
		BestTackler:      bestTackler,
		BestFouler:       bestFouler,
		MostYellowCards:  mostYellowCards,
		MostRedCards:     mostRedCards,
	}

	c.HTML(http.StatusOK, "player/analynse.html", analysisData)
}
