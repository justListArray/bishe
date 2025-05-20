package club

import (
	"fmt"
	"footballsys/controllers/base"
	"footballsys/controllers/match"
	"footballsys/models"
	"log"

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

func (club ClubController) Index(c *gin.Context) { //添加球员完成
	c.HTML(http.StatusOK, "member/index.html", nil)
}

func (club ClubController) AddMember(c *gin.Context) { //添加球员完成
	c.HTML(http.StatusOK, "member/addmember.html", nil)
}

func (club ClubController) SaveMember(c *gin.Context) { //添加球员完成
	var player models.Member
	c.JSON(http.StatusOK, c)
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Received player data: %+v", player)
	if err := models.DB.Create(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add Players"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player added successfully"})
}

func (club ClubController) DeleteMember1(c *gin.Context) {
	c.HTML(http.StatusOK, "member/deletemember.html", nil)
}

// player
func (club ClubController) DeleteMember(c *gin.Context) { //删除球员完成
	var request struct {
		PlayerId   int    `json:"player_id"`
		PlayerName string `json:"name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var player models.Member
	if request.PlayerId != 0 {
		if err := models.DB.First(&player, request.PlayerId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Player not found by ID"})
			return
		}
	} else if request.PlayerName != "" {
		if err := models.DB.Where("name = ?", request.PlayerName).First(&player).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Player not found by Name"})
			return
		}
	}

	if request.PlayerId == 0 && request.PlayerName == "" {
		if err := models.DB.Where("1=1").Delete(&models.Member{}).Error; err != nil { //永真
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all Players"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "All Players deleted successfully"})
		return
	}
	if request.PlayerId != 0 {
		if err := models.DB.Where("id = ?", strconv.Itoa(request.PlayerId)).Delete(&models.Member{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Player by id"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully by id"})
		return
	}
	if request.PlayerName != "" {
		if err := models.DB.Where("name = ?", request.PlayerName).Delete(&models.Member{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Player by name"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully by name"})
		return
	}
}

func (club ClubController) SearchMember1(c *gin.Context) {
	c.HTML(http.StatusOK, "member/searchmember.html", nil)
}

// player
func (club ClubController) SearchMember(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}
	var Player models.Member
	if err := c.ShouldBindJSON(&request); err != nil { // 使用 ShouldBindQuery 绑定查询参数
		log.Printf("Error binding query data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Printf("Received JSON data: %+v", request) // 打印接收到的 JSON 数据
	if request.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "没找到Player"})
		return
	}
	if err := models.DB.Where("name =?", request.Name).First(&Player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search Player by name"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Player": Player})
}

func (club ClubController) AddTeamPlayer1(c *gin.Context) {
	c.HTML(http.StatusOK, "member/addTeam.html", nil)
}

// 重点
func (club ClubController) AddTeamPlayer(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}
	err := c.ShouldBindJSON(&request) //球队名字（英文）
	log.Printf("球队的名字%v", request.Name)
	if err != nil {
		log.Printf("Error binding JSON data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	teamid := match.PerdictController.QueryTeamID(match.PerdictController(club), request.Name, c)
	season := 2023
	url := "https://v3.football.api-sports.io/players?season=" + strconv.Itoa(season) + "&team=" + strconv.Itoa(teamid)
	response := club.BaseController.PreOperation(c, url)
	// c.JSON(http.StatusOK, gin.H{"message": response})
	// 提取球员信息并保存到数据库
	players := response["response"].([]interface{})
	for _, player := range players {
		playerMap := player.(map[string]interface{})
		playerData := playerMap["player"].(map[string]interface{})
		playerId, _ := playerData["id"].(float64)
		var existingMember models.Member
		if err := models.DB.Where("player_id = ?", playerId).First(&existingMember).Error; err == nil {
			continue // 如果球员已存在，则跳过
		}
		// c.JSON(http.StatusOK, gin.H{
		// 	"ID": playerId,
		// })
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
			PlayerId: int(playerId),
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

// func (club ClubController) SearchAllMember1(c *gin.Context) {
// 	c.HTML(http.StatusOK, "member/searchAllMember.html", nil)
// }

func (club *ClubController) SearchAllMember(c *gin.Context) {
	var players []models.Member
	if err := models.DB.Find(&players).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法查询球员信息"})
		return
	}

	// 将查询结果转换为 map[int]models.Member
	playersMap := make(map[int]models.Member)
	for _, player := range players {
		playersMap[player.Id] = player
	}

	// 渲染模板并传递数据
	c.HTML(http.StatusOK, "member/searchAllMember.html", gin.H{
		"Members": playersMap,
	})
}
