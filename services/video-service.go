package services

import (
	"github.com/mic3ael/pragmaticreviews/entities"
)

type VideoService interface {
	Save(entities.Video) entities.Video
	FindAll() []entities.Video
}

type videoService struct {
	videos []entities.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entities.Video) entities.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entities.Video {
	if service.videos == nil {
		return make([]entities.Video, 0)
	}
	return service.videos
}
