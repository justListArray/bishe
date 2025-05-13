package index

import (
	"footballsys/controllers/base"
	"footballsys/models"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

type IndexController struct {
	base.BaseController
}

// var db *gorm.DB

//	func (index IndexController) fail(c *gin.Context) {
//		c.String(200, "失败")
//	}
func (index IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", nil)
}

func (index IndexController) Login1(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login.html", nil)
}

func (index IndexController) Login(c *gin.Context) {
	// 定义一个结构体用于接收请求中的用户名和密码

	var Login1 models.User
	// 将请求中的 JSON 数据绑定到 loginReq 结构体中

	if err := c.ShouldBindJSON(&Login1); err != nil {
		c.HTML(http.StatusBadRequest, "admin/login.html", gin.H{"error": "Invalid request data"})
		return
	}

	// 查询数据库，检查用户名和密码是否匹配
	var user models.User
	if result := models.DB.Where("username = ? AND password = ?", Login1.Username, Login1.Password).First(&user); result.Error != nil {
		// 如果查询失败，返回错误信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	if user.Identity != "admin" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "role": "admin"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "role": "user"})
	}
	session := sessions.Default(c)
	session.Set("user_role", user.Identity)
	session.Save()
	// 如果查询成功，返回登录成功信息
	c.JSON(http.StatusOK, user)
}

func (index IndexController) Signin1(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/signin.html", nil)
}

func (index IndexController) Signin(c *gin.Context) { //ok
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} //它的作用是将 HTTP 请求中的 JSON 数据绑定到 user 结构体中，并在绑定失败时返回错误信息。
	log.Printf("注册的username:%v,注册的identity:%v", user.Username, user.Identity)
	usernameRegex := `^[a-zA-Z0-9_]{6,20}$`
	passwordRegex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d!@#$%^&*()]{8,20}$`
	matched, _ := regexp.MatchString(usernameRegex, user.Username)
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}
	matched, _ = regexp.MatchString(passwordRegex, user.Password)
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}
	if err := models.DB.Create(&user).Error; err != nil { //注册
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}
	session := sessions.Default(c)
	session.Set("user_role", user.Identity)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": user})
}
