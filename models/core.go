package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"footballsys/utils"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var FormTemplate *template.Template
var FormTemplateContent string

func init() {
	dsn := "bishe:123456@tcp(127.0.0.1:3306)/fbsys?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	err = DB.AutoMigrate()
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	data, err := ioutil.ReadFile("./templates/test/test.html")
	if err != nil {
		log.Fatalf("Error reading template file: %v", err)
	}
	FormTemplateContent = string(data)
	if err != nil {
		log.Fatalf("Error parsing template file: %v", err)
	}

}

func CheckToken(c *gin.Context) {
	fmt.Println("进入拦截器")
	// 编写过滤的逻辑，token校验等
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.Redirect(302, "http://150.158.2.43:8080/index/login")
		return
	}
	_, err := utils.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	//  c.Next():放行执行后续动作  c.Abort():不执行后续动作
	c.Next()
}
