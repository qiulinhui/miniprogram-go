package middlewares

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := os.MkdirAll("./logs/"+time.Now().Format("2006/1"), 0644)
		if err != nil {
			panic(err)
		}
		f, err := os.OpenFile("./logs/"+time.Now().Format("2006/1/2")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(f)
		raw, _ := c.GetRawData()
		log.WithFields(log.Fields{
			"ip":     c.ClientIP(),
			"method": c.Request.Method,
			"path":   c.FullPath(),
			"req":    string(raw),
			"query":  c.Request.URL.RawQuery,
			"header": c.Request.Header,
		}).Info()
		c.Next()
	}
}
