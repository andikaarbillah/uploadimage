package controller

import (
	"images/model"
	"images/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	is  service.ImageService
	ctx *gin.Context
}

func (ic ImageController) VerifikasiImage(ctx *gin.Context) {
	var request model.ImageRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ic.is.Create(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "succesfully",
		"data":    data,
	})
}

func(ic ImageController)VerifikasiDeleteImage(ctx *gin.Context){
	imageID := ctx.PostForm("id")
	err := ic.is.Delete(ctx, imageID)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "succesfully delete image",
	})
}

func NewUserController(is service.ImageService, ctx *gin.Context) ImageController {
	return ImageController{
		is:  is,
		ctx: ctx,
	}
}
