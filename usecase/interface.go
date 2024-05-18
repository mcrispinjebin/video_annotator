package usecase

import (
	"context"
	"video_annotator/models"
)

type VideoUsecase interface {
	CreateVideo(ctx context.Context, video *models.Video) (err error)
	GetVideo(ctx context.Context, videoID string) (video models.Video, err error)
	DeleteVideo(ctx context.Context, videoID string) (err error)
}

type AnnotationUsecase interface {
	CreateAnnotation(ctx context.Context, annotation *models.Annotation) (err error)
	UpdateAnnotation(ctx context.Context, id string, annotations *models.Annotation) (err error)
	DeleteAnnotation(ctx context.Context, id string) (err error)
}

type Usecase struct {
	VideoUsecase      VideoUsecase
	AnnotationUsecase AnnotationUsecase
}
