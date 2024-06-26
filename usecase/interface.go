package usecase

import (
	"context"
	"video_annotator/models"
)

type VideoUsecase interface {
	CreateVideo(ctx context.Context, video *models.Video) (err *models.CustomErr)
	GetVideo(ctx context.Context, videoID string) (video models.Video, err *models.CustomErr)
	DeleteVideo(ctx context.Context, videoID string) (err *models.CustomErr)
}

type AnnotationUsecase interface {
	CreateAnnotation(ctx context.Context, annotation *models.Annotation) (err *models.CustomErr)
	UpdateAnnotation(ctx context.Context, videoID string,
		annotationUpdate *models.Annotation) (fetchedAnnotation models.Annotation, err *models.CustomErr)
	DeleteAnnotation(ctx context.Context, videoID, annotationID string) (err *models.CustomErr)
}

type Usecase struct {
	VideoUsecase      VideoUsecase
	AnnotationUsecase AnnotationUsecase
}
