package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zemags/go_workshop_3/controller"
	"github.com/zemags/go_workshop_3/middlewares"
	"github.com/zemags/go_workshop_3/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8000")
}
