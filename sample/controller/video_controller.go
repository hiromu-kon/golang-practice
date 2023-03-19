package controller

import (
	"sample/entity"
	"sample/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	videoService service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {
	return &videoController{videoService: service}
}

func (controller *videoController) FindAll() []entity.Video {
	return controller.videoService.FindAll()
}

func (controller *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	controller.videoService.Save(video)
	return video
}
