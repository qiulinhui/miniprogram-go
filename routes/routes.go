package routes

import (
	"app/web/controllers"
	"app/web/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	r.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "notfound",
		})
	})

	var (
		novelController = controllers.NovelController
		userController  = controllers.UserController
	)

	jwt := middleware.JwtMiddleware()
	authz := middleware.AuthzMiddleware()
	r.POST("/login", jwt.LoginHandler) // 授权登录
	auth := r.Group("/auth")
	auth.Use(jwt.MiddlewareFunc())
	auth.GET("/refresh_token", jwt.RefreshHandler) // 刷新令牌
	auth.GET("/logout", jwt.LogoutHandler)         // 退出登录

	user := r.Group("/user")
	user.Use(jwt.MiddlewareFunc())
	{
		user.GET("/hello", userController.Hello)
	}

	admin := r.Group("/admin")
	admin.Use(authz)

	novel := r.Group("/novel")
	{
		novel.GET("/:id", novelController.Get)
		novel.PATCH("/:id", novelController.Update)
		novel.DELETE("/:id", novelController.Delete)
		novel.POST("", novelController.Create)
	}
}
