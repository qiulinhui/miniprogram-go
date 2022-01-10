package controllers

import (
	"bookstore/model"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (*UserController) Hello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get("id")

	c.JSON(200, gin.H{
		"userID": claims["id"],
		"OpenId": user.(*model.User).Openid,
		"text":   "hello world",
	})
}
