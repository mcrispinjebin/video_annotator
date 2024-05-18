package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"video_annotator/models"
	"video_annotator/store"
)

type videoUsecase struct {
	repoStore store.VideoStore
}

func NewVideoUsecase(store store.Store) VideoUsecase {
	return videoUsecase{repoStore: store.VideoStore}
}

func (v videoUsecase) CreateVideo(ctx context.Context, video *models.Video) (err error) {
	//check source

	videos, err := v.repoStore.GetAllVideos(ctx)
	if err != nil {
		return err
	}

	for _, existingVideo := range videos {
		if existingVideo.Url == video.Url {
			err = errors.New(fmt.Sprintf("video with Url %s already exists", video.Url))
			return err
		}
	}

	video.ID = uuid.New().String()
	err = v.repoStore.CreateNewVideo(ctx, video)
	if err != nil {
		return err
	}

	return
}

func (v videoUsecase) GetVideo(ctx context.Context, videoID string) (video models.Video, err error) {
	video, err = v.repoStore.GetVideoByID(ctx, videoID)
	if err != nil {
		return
	}

	return video, nil
}

func (v videoUsecase) DeleteVideo(ctx context.Context, videoID string) (err error) {
	video, err := v.repoStore.GetVideoByID(ctx, videoID)
	if err != nil {
		err = errors.New(fmt.Sprintf("Video with ID %s is not found to be deleted", videoID))
		return
	}

	if err = v.repoStore.DeleteVideo(ctx, video); err != nil {
		return err
	}

	return nil
}
