package bootstrap

import (
	"app/config"
	"fmt"
)

func Start() {
	port := config.GetInt("app.port", 8080)
	r.Run(fmt.Sprintf(":%v", port))
}
