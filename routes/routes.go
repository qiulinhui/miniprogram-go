package routes

import (
	"app/controllers"
	"app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	r.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "notfound",
		})
	})
	type controller struct {
		Book *controllers.BookController
		User *controllers.UserController
	}
	var c controller
	jwt := middlewares.JwtMiddleware()
	authz := middlewares.AuthzMiddleware()
	r.POST("/login", jwt.LoginHandler) // 授权登录
	auth := r.Group("/auth")
	auth.Use(jwt.MiddlewareFunc())
	auth.GET("/refresh_token", jwt.RefreshHandler) // 刷新令牌
	auth.GET("/logout", jwt.LogoutHandler)         // 退出登录

	user := r.Group("/user")
	user.Use(jwt.MiddlewareFunc())
	{
		user.GET("/hello", c.User.Hello)
	}

	admin := r.Group("/admin")
	admin.Use(authz)

	book := r.Group("/book")
	{
		book.GET("/book/:id", c.Book.Get)
		book.PATCH("/book", c.Book.Update)
		book.DELETE("/book/:id", c.Book.Delete)
		book.POST("/books/create", c.Book.Create)
	}
}
