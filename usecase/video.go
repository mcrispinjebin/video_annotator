package usecase

import (
	"context"
	"fmt"
	"video_annotator/constants"
	"video_annotator/models"
	"video_annotator/store"
)

type videoUsecase struct {
	videoStore store.VideoStore
}

func NewVideoUsecase(store store.Store) VideoUsecase {
	return videoUsecase{videoStore: store.VideoStore}
}

func (v videoUsecase) CreateVideo(ctx context.Context, video *models.Video) (err *models.CustomErr) {
	videos, err := v.videoStore.GetAllVideos(ctx)
	if err != nil {
		return err
	}

	for _, existingVideo := range videos {
		if existingVideo.Url == video.Url {
			err.Message = fmt.Sprintf(constants.VideoWithSameUrlExistsErr, video.Url)
			return err
		}
	}

	err = v.videoStore.CreateNewVideo(ctx, video)
	if err != nil {
		return err
	}

	return
}

func (v videoUsecase) GetVideo(ctx context.Context, videoID string) (video models.Video, err *models.CustomErr) {
	video, err = v.videoStore.GetVideoByID(ctx, videoID, true)
	if err != nil {
		return
	}

	return video, nil
}

func (v videoUsecase) DeleteVideo(ctx context.Context, videoID string) (err *models.CustomErr) {
	video, err := v.videoStore.GetVideoByID(ctx, videoID, true)
	if err != nil {
		return
	}

	if err = v.videoStore.DeleteVideo(ctx, video); err != nil {
		return err
	}

	return nil
}
