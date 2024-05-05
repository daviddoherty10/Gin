package main

import (
	"github.com/daviddoherty10/Gin/controller"
	"github.com/daviddoherty10/Gin/service.go"
	"github.com/gin-gonic/gin"
)

var (
    videoService service.VideoService = service.New()
    videoController controller.VideoController = controller.New(videoService)
)

func main() {
    server := gin.Default(); 

    server.GET("/videos", func(ctx *gin.Context) {
        ctx.JSON(200,videoController.FindAll())
    })

    server.POST("/videos", func(ctx *gin.Context) {
        ctx.JSON(200,videoController.Save(ctx))
    })

    server.Run(":8080")
}
