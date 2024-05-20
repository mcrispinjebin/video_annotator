package mocks

import (
	"context"
	"time"
	"video_annotator/constants"
	"video_annotator/models"
)

type MockVideoStore struct {
	Videos map[string]models.Video
}

func (store *MockVideoStore) CreateNewVideo(_ context.Context, video *models.Video) (err *models.CustomErr) {

	if _, exists := store.Videos[video.ID]; exists {
		return &models.CustomErr{StatusCode: constants.HttpResourceExists, Message: "Video already exists"}
	}

	video.CreatedAt = time.Now()
	store.Videos[video.ID] = *video
	return nil
}

func (store *MockVideoStore) GetVideoByID(_ context.Context, videoID string, _ bool) (video models.Video, err *models.CustomErr) {
	video, exists := store.Videos[videoID]
	if !exists {
		return models.Video{}, &models.CustomErr{StatusCode: constants.HttpResourceNotFound,
			Message: constants.VideoResourceNotFound}
	}

	return video, nil
}

func (store *MockVideoStore) DeleteVideo(_ context.Context, video models.Video) (cErr *models.CustomErr) {
	if _, exists := store.Videos[video.ID]; !exists {
		return &models.CustomErr{StatusCode: constants.HttpResourceNotFound, Message: constants.VideoDeleteErr}
	}

	delete(store.Videos, video.ID)
	return nil
}

func (store *MockVideoStore) GetAllVideos(_ context.Context) (videos []models.Video, err *models.CustomErr) {
	for _, video := range store.Videos {
		videos = append(videos, video)
	}

	return videos, nil
}
