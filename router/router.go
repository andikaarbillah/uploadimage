package router

import (
	"images/controller"
	"images/repository"
	"images/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Router(router *gin.Engine, db *gorm.DB) {
	ImageRepo := repository.NewImageRepository(db)
	ImageService := service.NewImageRepository(ImageRepo)
	ImageController := controller.NewUserController(ImageService, ctx)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"SuccesFully": "coy",
		})
	})

	v1 := router.Group("/api/v1")

	{
		v1.POST("image/delete", ImageController.VerifikasiDeleteImage)
		v1.POST("/image", ImageController.VerifikasiImage)
	}
}
