package controllers

import (
	"bookstore/model"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get("id")
	c.JSON(http.StatusOK, gin.H{
		"userID":   claims["id"],
		"userName": user.(*model.User).Nickname,
		"text":     "Hello World.",
	})
}
