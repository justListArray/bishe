package train

import (
	"footballsys/controllers/base"
	"footballsys/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainController struct {
	base.BaseController
}

func (train TrainController) TrainAdd(c *gin.Context) { //添加训练信息
	var a models.Train

	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// dateLayout := "2006-01-02" // *定义日期格式*
	// date, err := time.Parse(dateLayout, a.Date.String())
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
	// 	return
	// }
	// a.Date = date
	if err := models.DB.Create(&a).Error; err != nil { //注册
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fill in the training information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Complete the training information successfully:"})
}

func (train TrainController) TrainSearch(c *gin.Context) { //查询训练信息
	name := c.Query("name")
	id := c.Query("user_id")
	var a models.Train
	if id == "" && name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "没找到member"})
		return
	}
	if id != "" {
		if err := models.DB.Where("user_id=?", id).First(&a).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search member by id"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": a})
		return
	}
	if name != "" {
		if err := models.DB.Where("name=?", name).First(&a).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search member by name"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": a})
		return
	}

}
