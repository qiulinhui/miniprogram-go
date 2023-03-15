package controllers

import (
	"app/models"
	"app/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

var NovelController = newNovelController()

type novelController struct {
}

func newNovelController() *novelController {
	return &novelController{}
}

var novelRepo = repositories.NovelRepository

func (*novelController) Create(ctx *gin.Context) {
	name := ctx.Query("name")
	author := ctx.Query("author")

	ok, novel := novelRepo.Create(&models.Novel{
		Name:   name,
		Author: author,
	})

	if ok {
		ctx.JSON(http.StatusBadRequest, "新增失败")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": novel,
	})
}

func (*novelController) Update(c *gin.Context) {

}

func (*novelController) Delete(c *gin.Context) {

}

func (c *novelController) Get(ctx *gin.Context) {
	var p struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
	}
	novel := novelRepo.Find(int64(p.ID))

	ctx.JSON(http.StatusOK, gin.H{
		"data": novel,
	})
}
