package index

import (
	"encoding/json"
	"fmt"
	"footballsys/controllers/base"
	"footballsys/models"
	"footballsys/utils"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/dchest/captcha"
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

func (index IndexController) Login(c *gin.Context) { //登录 // token 生成
	// 定义一个结构体用于接收请求中的用户名和密码

	body, _ := io.ReadAll(c.Request.Body)
	log.Printf("Raw request body: %s", body)
	var Login1 struct {
		username string
		password string
	}
	// 将请求中的 JSON 数据绑定到 loginReq 结构体中
	//log.Print(c)
	if err := json.Unmarshal(body, &Login1); err != nil {
		log.Printf("Invalid request data: %+v", err)
		log.Printf("Login1: %v", Login1)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// 查询数据库，检查用户名和密码是否匹配
	var user models.User
	if result := models.DB.Where("username = ? AND password = ?", Login1.username, Login1.password).Find(&user); result.Error != nil {
		// 如果查询失败，返回错误信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	// if user.Identity != "admin" {
	// 	//c.JSON(http.StatusOK, gin.H{"message": "Login successful", "role": "admin"})
	// } else {
	// 	//c.JSON(http.StatusOK, gin.H{"message": "Login successful", "role": "user"})
	// }
	session := sessions.Default(c)
	session.Set("user_role", user.Identity)
	session.Save()
	// 如果查询成功，返回登录成功信息
	token, err := utils.GenerateToken(uint(user.Id)) //生出token
	if err != nil {
		log.Printf("TOEKN 生成失败:%+v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func (index IndexController) Signin1(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/signin.html", nil)
}

func (index IndexController) Signin(c *gin.Context) { //ok 注册 // token 生成
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} //它的作用是将 HTTP 请求中的 JSON 数据绑定到 user 结构体中，并在绑定失败时返回错误信息。
	log.Printf("注册的username:%v,注册的identity:%v", user.Username, user.Identity)
	usernameRegex := `^[a-zA-Z0-9_]{6,20}$`
	passwordRegex := `^[a-zA-Z0-9_]{6,20}$`
	matched, _ := regexp.MatchString(usernameRegex, user.Username)
	if !matched {
		log.Printf("注册的username格式不对")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}
	matched, _ = regexp.MatchString(passwordRegex, user.Password)
	if !matched {
		log.Printf("注册的Password格式不对")
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
	token, err := utils.GenerateToken(uint(user.Id)) //生出token
	if err != nil {
		log.Printf("TOEKN 生成失败:%+v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
	// index.Yanzhengma(c)
}

func (index IndexController) DeleteUserSession(userID int) (err error) { // 执行注销逻辑
	err1 := models.DB.Exec(`delete from users where id =?`, userID).Error
	return err1
}

func (index IndexController) Logoff1(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/delete.html", nil)
}

// token 解析
func (index IndexController) Logoff(c *gin.Context) { // 注销  找到数据库的数据，执行删除操作，因为在登录后才能注销，所以不需要判断是否存在该数据
	tokenString := c.GetHeader("Authorization")
	log.Println(tokenString)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
		return
	}
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	userID := claims.UserID
	log.Println(claims)

	// 从数据库中删除用户
	err = index.DeleteUserSession(int(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	// 返回注销成功的响应
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})

}
func (index IndexController) UpdatePass1(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/update.html", nil)
}

func (index IndexController) UpdatePass(c *gin.Context) { // 更改密码  找到数据库的数据，执行修改操作，因为在登录后才能注销，所以不需要判断是否存在该数据
	tokenString := c.GetHeader("Authorization")
	var response struct {
		username string
		password string
	}
	err := c.ShouldBindJSON(&response)
	if err != nil {
		log.Fatalf("无法把response ShouldBindJSON")
	}
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
		return
	}
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	userID := claims.UserID

	// 从数据库中修改密码
	err = index.UpdateUserSession(int(userID), response.password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update Password"})
		return
	}

	// 返回注销成功的响应
	c.JSON(http.StatusOK, gin.H{"message": "Update Password successful"})
}

func (index IndexController) UpdateUserSession(userID int, pwd string) error {
	_ = models.DB.Exec(`update users set password=? where id =?`, pwd, userID)
	return nil
}

func (index IndexController) Yanzhengma1(c *gin.Context) { //验证码
	c.HTML(http.StatusOK, "test/test.html", nil)
}
func (index IndexController) Yanzhengma(c *gin.Context) { //验证码
	http.HandleFunc("/yanzhengma", showFormHandler)
	http.HandleFunc("/process", processFormHandler)
	http.Handle("/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))
	fmt.Println("Server is at localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}

// var FormTemplate *template.Templateform
// var err error
// FormTemplate, err = template.ParseFiles("test/test.html")
//
//	if err != nil {
//		log.Fatalf("Error parsing template file: %v", err)
//	}
var formTemplate = template.Must(template.New("example").Parse(models.FormTemplateContent))

func showFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/yanzhengma" {
		http.NotFound(w, r)
		return
	}
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	if err := formTemplate.Execute(w, &d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func processFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if !captcha.VerifyString(r.FormValue("captchaId"), r.FormValue("captchaSolution")) {
		io.WriteString(w, "Wrong captcha solution! No robots allowed!\n")
	} else {
		io.WriteString(w, "Great job, human! You solved the captcha.\n")
	}
	io.WriteString(w, "<br><a href='/'>Try another one</a>")
}
