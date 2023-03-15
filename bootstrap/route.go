package bootstrap

import (
	"app/routes"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRoutes() {
	r = gin.Default()
	routes.Register(r)
}
