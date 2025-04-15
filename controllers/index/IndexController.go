package index

import (
	"footballsys/controllers/base"
	"footballsys/models"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

type IndexController struct {
	base.BaseController
}

// var db *gorm.DB

func (index IndexController) Success(c *gin.Context) {
	c.String(200, "成功")
}

// func (index IndexController) fail(c *gin.Context) {
// 	c.String(200, "失败")
// }

func (index IndexController) Index1(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", nil)
}
func (index IndexController) Login(c *gin.Context) {
	// 定义一个结构体用于接收请求中的用户名和密码

	var Login1 models.User
	// 将请求中的 JSON 数据绑定到 loginReq 结构体中

	if err := c.ShouldBindJSON(&Login1); err != nil {
		c.HTML(http.StatusBadRequest, "admin/index.html", gin.H{"error": "Invalid request data"})
		return
	}

	// 查询数据库，检查用户名和密码是否匹配
	var user models.User
	if result := models.DB.Where("username = ? AND password = ?", Login1.Username, Login1.Password).First(&user); result.Error != nil {
		// 如果查询失败，返回错误信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 如果查询成功，返回登录成功信息
	c.JSON(http.StatusOK, user)
}

func (index IndexController) Signin(c *gin.Context) { //ok
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} //它的作用是将 HTTP 请求中的 JSON 数据绑定到 user 结构体中，并在绑定失败时返回错误信息。
	if err := models.DB.Create(&user).Error; err != nil { //注册
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully:"})
}
