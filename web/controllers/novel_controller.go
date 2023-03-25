package controllers

import (
	"app/app"
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
		novelService: services.NewNovelService(repositories.NewNovelRepository(app.DB)),
	}
}

type validate struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Author string `form:"author" json:"author" binding:"required"`
}

func (c *novelController) Get(ctx *gin.Context) {

	var p struct {
		ID uint64 `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
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
