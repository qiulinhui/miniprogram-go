package bootstrap

import (
	"app/app"
	"fmt"
)

func InItGin() {
	port := app.Config.GetInt("app.port", 8080)
	r.Run(fmt.Sprintf(":%v", port))
}
