package store

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"video_annotator/constants"
	"video_annotator/models"
)

type videoStore struct {
	DB *gorm.DB
}

func NewVideoStore(db *gorm.DB) VideoStore {
	return &videoStore{DB: db}
}

func (v videoStore) CreateNewVideo(_ context.Context, video *models.Video) (err *models.CustomErr) {
	video.ID = uuid.New().String()
	result := v.DB.Create(video)
	if result.Error != nil {
		return &models.CustomErr{Err: result.Error, Message: constants.VideoCreateErr}
	}
	return nil
}

func (v videoStore) GetVideoByID(_ context.Context, videoID string, includeAnnotations bool) (video models.Video, err *models.CustomErr) {
	tx := v.DB
	if includeAnnotations {
		tx = v.DB.Preload("Annotations", "active = ?", true)
	}

	result := tx.First(&video, "id = ? AND active = ?", videoID, true)
	if result.Error != nil {
		err = &models.CustomErr{Err: result.Error}
		if errors.Is(err.Err, gorm.ErrRecordNotFound) {
			err.Message = constants.VideoResourceNotFound
			err.StatusCode = constants.HttpResourceNotFound
			return
		}
		err.Message = constants.VideoGetFetchErr
		return
	}
	return video, nil
}

func (v videoStore) DeleteVideo(_ context.Context, video models.Video) (cErr *models.CustomErr) {
	//set all annotations mapped to a video to inactive
	if err := v.DB.Model(&models.Annotation{}).
		Where("video_id = ?", video.ID).
		Update("active", false).Error; err != nil {
		return &models.CustomErr{Err: err, Message: constants.VideoDeleteErr}
	}

	video.Active = false
	result := v.DB.Save(&video)
	if result.Error != nil {
		return &models.CustomErr{Err: result.Error, Message: constants.VideoDeleteErr}
	}

	return nil
}

func (v videoStore) GetAllVideos(_ context.Context) (videos []models.Video, err *models.CustomErr) {
	result := v.DB.Find(&videos, "active = ?", true)
	if result.Error != nil {
		return videos, &models.CustomErr{Err: result.Error, Message: constants.VideoResourceNotFound}
	}
	return videos, nil
}
