package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"video_annotator/models"
)

type videoStore struct {
	DB *gorm.DB
}

func NewVideoStore(db *gorm.DB) VideoStore {
	return &videoStore{DB: db}
}

func (v videoStore) CreateNewVideo(_ context.Context, video *models.Video) (err error) {
	video.ID = uuid.New().String()
	result := v.DB.Create(video)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (v videoStore) GetVideoByID(_ context.Context, videoID string, includeAnnotations bool) (video models.Video, err error) {
	tx := v.DB
	if includeAnnotations {
		tx = v.DB.Preload("Annotations")
	}

	result := tx.First(&video, "id = ? AND active = ?", videoID, true)
	if result.Error != nil {
		err = result.Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("no video is matched with given ID %v", err)
			return
		}
		return
	}
	return video, nil
}

func (v videoStore) DeleteVideo(_ context.Context, video models.Video) (err error) {
	if err = v.DB.Model(&models.Annotation{}).
		Where("video_id = ?", video.ID).
		Update("active", false).Error; err != nil {
		return err
	}

	video.Active = false
	result := v.DB.Save(&video)
	if result.Error != nil {
		err = result.Error
		return
	}

	return nil
}

func (v videoStore) GetAllVideos(_ context.Context) (videos []models.Video, err error) {
	result := v.DB.Find(&videos, "active = ?", true)
	if result.Error != nil {
		err = result.Error
		return
	}
	return videos, nil
}
