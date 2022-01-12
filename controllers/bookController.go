package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
}

var bookModel models.Book

func (*BookController) Create(c *gin.Context) {
	name := c.Query("name")
	author := c.Query("author")
	bookModel.Name = name
	bookModel.Author = author
	err := bookModel.Insert()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"data": bookModel,
	})
}

func (*BookController) Update(c *gin.Context) {

}

func (*BookController) Delete(c *gin.Context) {

}

func (ctr *BookController) Get(c *gin.Context) {
	var p struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
	}
	bookModel.ID = p.ID
	bookModel.SelectByID()
	c.JSON(http.StatusOK, gin.H{
		"data": bookModel,
	})
}
