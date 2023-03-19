package service

import "sample/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func NewVideoService() VideoService {
	// return &videoService{}
	service := &videoService{
		videos: []entity.Video{
			{
				Title:       "My First Video",
				Description: "This is a video description",
				URL:         "https://example.com/video",
			},
		},
	}

	return service
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
