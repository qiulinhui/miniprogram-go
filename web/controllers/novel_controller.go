package controllers

import (
	"app/models"
	"app/repositories"
	"app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type novelController struct {
	novelService services.NovelService
}

func NewNovelController() *novelController {
	return &novelController{
		novelService: services.NewNovelService(repositories.NewNovelRepository()),
	}
}

func (c *novelController) Get(ctx *gin.Context) {

	var p struct {
		ID uint64 `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	novel := c.novelService.Get(p.ID)
	if novel == nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": novel,
	})
}

func (c *novelController) Add(ctx *gin.Context) {
	var validate struct {
		Name   string `form:"name" json:"name" binding:"required"`
		Author string `form:"author" json:"author" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	ok := c.novelService.Add(&models.Novel{
		Name:   validate.Name,
		Author: validate.Author,
	})

	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "写入成功",
		})
	}

}
