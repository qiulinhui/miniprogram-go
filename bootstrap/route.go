package bootstrap

import (
	"bookstore/routes"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func SetUpRoute() {
	r = gin.Default()
	routes.Register(r)
}
