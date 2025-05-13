package controllers

import (
	"footballsys/controllers/base"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MiddleController struct {
	base.BaseController
}

// func (mid MiddleController) AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		session := sessions.Default(c)
// 		userRole := session.Get("user_role")

// 		if userRole == nil {
// 			// 用户未登录，重定向到登录页面
// 			c.Redirect(http.StatusFound, "https://150.158.2.43:8080/index/index")
// 			return
// 		}

// 		// 设置用户角色到上下文中，供后续中间件或处理器使用
// 		c.Set("user_role", userRole)
// 		c.Next()
// 	}
// }

func (mid MiddleController) AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.MustGet("user_role").(string)

		if userRole != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}
