package routes

import (
	"bookstore/app/controllers"
	"bookstore/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "notfound",
		})
	})
	r.POST("/login", lib.JWT().LoginHandler)
	// e, err := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// r.Use(authz.NewAuthorizer(e))
	auth := r.Group("/auth")
	auth.Use(lib.JWT().MiddlewareFunc())
	{
		auth.GET("/refresh_token", lib.JWT().RefreshHandler)
	}
	auth.GET("/hello", controllers.HelloHandler).Use(lib.JWT().MiddlewareFunc())

	var bookController *controllers.BookController
	r.GET("/book/:id", bookController.Get)
	r.PATCH("/book", bookController.Update)
	r.DELETE("/book/:id", bookController.Delete)
	r.POST("/books/create", bookController.Create)
}
