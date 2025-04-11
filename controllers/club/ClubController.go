package club

import (
	"footballsys/controllers/base"
	"footballsys/models"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

type ClubController struct {
	base.BaseController
}

// var db *gorm.DB

func (club ClubController) AddMember(c *gin.Context) { //添加球员完成
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add member"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Member added successfully", "member": member})
}

func (club ClubController) DeleteMember(c *gin.Context) { //删除球员完成
	id := c.Query("id")
	name := c.Query("name")
	if id == "" && name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "没找到member"})
	}
	if id != "" {
		if err := models.DB.Where("id = ?", id).Delete(&models.Member{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete member by id"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully by id"})
		return
	}
	if name != "" {
		if err := models.DB.Where("name = ?", name).Delete(&models.Member{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete member by name"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully by name"})
		return
	}
}

func (club ClubController) SearchMember(c *gin.Context) {
	var member models.Member
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "没找到member"})
		return
	}
	if err := models.DB.Where("name = ?", name).First(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search member by name"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"member": member})
}

//func(club ClubController)
