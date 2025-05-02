package club

import (
	"encoding/json"
	"fmt"
	"footballsys/controllers/base"
	"footballsys/controllers/match"
	"footballsys/models"
	"io"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

type ClubController struct {
	base.BaseController
}

// var db *gorm.DB
// player
func (club ClubController) AddMember(c *gin.Context) { //添加球员完成
	var player models.Member
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add Players"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Players added successfully", "member": player})
}

// player
func (club ClubController) DeleteMember(c *gin.Context) { //删除球员完成
	id := c.Query("id")
	name := c.Query("name")
	if id == "" && name == "" {
		if err := models.DB.Where("1=1").Delete(&models.Member{}).Error; err != nil { //永真
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all Players"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "All Players deleted successfully"})
		return
	}
	if id != "" {
		if err := models.DB.Where("id = ?", id).Delete(&models.Member{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Player by id"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully by id"})
		return
	}
	if name != "" {
		if err := models.DB.Where("name = ?", name).Delete(&models.Member{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Player by name"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully by name"})
		return
	}
}

// player
func (club ClubController) SearchMember(c *gin.Context) {
	var Player models.Member
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "没找到Player"})
		return
	}
	if err := models.DB.Where("name =?", name).First(&Player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search Player by name"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Player": Player})
}

// 重点
func (club ClubController) AddTeamPlayer(c *gin.Context) { //如果重复怎么办
	name := c.Query("name") //球队名字（英文）
	teamid := match.QueryTeamID(name, c)
	season := 2021
	url := "https://v3.football.api-sports.io/players?season=" + strconv.Itoa(season) + "&team=" + strconv.Itoa(teamid)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("x-rapidapi-key", models.ApiKey)
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	if errors, ok := response["errors"].(map[string]interface{}); ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"message": response})
	// 提取球员信息并保存到数据库
	players := response["response"].([]interface{})
	for _, player := range players {
		playerMap := player.(map[string]interface{})
		playerData := playerMap["player"].(map[string]interface{})
		name, _ := playerData["name"].(string)
		age, _ := playerData["age"].(float64)
		// position := ""
		// if playerData["position"] != nil {
		// 	position = playerData["position"].(string)
		// }
		// jerseyNumber := 0
		// if playerData["jersey_number"] != nil {
		// 	jerseyNumber = int(playerData["jersey_number"].(float64))
		// }
		member := models.Member{
			Username: fmt.Sprintf("%s", name),
			Identity: "Player",
			Name:     name,
			Age:      int(age),
			//Position:      position,
			//Jersey_number: jerseyNumber,
		}
		models.DB.Create(&member)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Players added successfully"})

}

// func (club ClubController) StatisticPlayer(c *gin.Context){

// }

// func (club ClubController) AnalysePlayer(c *gin.Context) {
// 	players := c.Query("players")

// }
