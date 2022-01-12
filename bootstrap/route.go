package bootstrap

import (
	"app/routes"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func SetUpRoute() {
	r = gin.Default()
	routes.Register(r)
}
