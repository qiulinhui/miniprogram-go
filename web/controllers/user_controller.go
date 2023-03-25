package controllers

import (
	"app/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (*userController) Hello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get("id")

	c.JSON(200, gin.H{
		"userID": claims["id"],
		"OpenId": user.(*models.User).Openid,
		"text":   "hello world",
	})
}
