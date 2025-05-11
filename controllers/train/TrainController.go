package train

import (
	"footballsys/controllers/base"
	"footballsys/controllers/player"
	"footballsys/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainController struct {
	base.BaseController
}

func (train TrainController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "train/index.html", nil)
}

func (train TrainController) AddTraibInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "train/addTrain.html", nil)

}

func (train TrainController) TrainAdd(c *gin.Context) { //添加训练信息
	var train1 models.Train

	if err := c.ShouldBindJSON(&train1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	train1.UserId, _ = player.PlayerController.Trans(player.PlayerController(train), train1.Name)

	log.Printf("前端发送的信息:%v", train1)
	// dateLayout := "2006-01-02" // *定义日期格式*
	// date, err := time.Parse(dateLayout, a.Date.String())
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
	// 	return
	// }
	// a.Date = date
	if err := models.DB.Create(&train1).Error; err != nil { //注册
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fill in the training information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Complete the training information successfully:"})
}

func (train TrainController) TrainSearch1(c *gin.Context) { //查询训练信息
	c.HTML(http.StatusOK, "train/searchTrain.html", nil)
}
func (train TrainController) TrainSearch(c *gin.Context) { //查询训练信息
	var request struct {
		PlayerId   int    `json:"id"`
		PlayerName string `json:"name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("训练的id:%d,训练的名字：%v", request.PlayerId, request.PlayerName)
	// name := c.Query("name")
	// id := c.Query("user_id")

	if request.PlayerId == 0 && request.PlayerName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "没找到member"})
		return
	}
	var trainCount int64
	if request.PlayerId != 0 {
		models.DB.Model(&models.Train{}).Where("user_id = ?", request.PlayerId).Count(&trainCount)
		c.JSON(http.StatusOK, gin.H{"message": gin.H{"player_id": request.PlayerId, "train_count": trainCount}})
		return
	}
	if request.PlayerName != "" {
		models.DB.Model(&models.Train{}).Where("name = ?", request.PlayerName).Count(&trainCount)
		if trainCount == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "他没有训练"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": gin.H{"player_name": request.PlayerName, "train_count": trainCount}})

		return
	}

}
