package store

import (
	"context"
	"gorm.io/gorm"
	"video_annotator/models"
)

type videoStore struct {
	DB *gorm.DB
}

func NewVideoStore(db *gorm.DB) VideoStore {
	return &videoStore{DB: db}
}

func (v videoStore) CreateNewVideo(ctx context.Context, video *models.Video) (err error) {
	result := v.DB.Create(video)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (v videoStore) GetVideoByID(ctx context.Context, videoID string) (video models.Video, err error) {
	result := v.DB.First(&video, "id = ?", videoID)
	if result.Error != nil {
		err = result.Error
		return
	}
	return video, nil
}

func (v videoStore) DeleteVideo(ctx context.Context, video models.Video) (err error) {
	result := v.DB.Delete(video)
	if result.Error != nil {
		err = result.Error
		return
	}

	return nil
}

func (v videoStore) GetAllVideos(ctx context.Context) (videos []models.Video, err error) {
	//vw := make([]models.Video, 0)
	result := v.DB.Find(&videos)
	if result.Error != nil {
		err = result.Error
		return
	}
	return videos, nil
}
