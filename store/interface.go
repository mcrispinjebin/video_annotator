package store

import (
	"context"
	"video_annotator/models"
)

type VideoStore interface {
	CreateNewVideo(ctx context.Context, video *models.Video) (err *models.CustomErr)
	GetVideoByID(ctx context.Context, videoID string, includeAnnotations bool) (video models.Video, err *models.CustomErr)
	DeleteVideo(ctx context.Context, video models.Video) (cErr *models.CustomErr)
	GetAllVideos(ctx context.Context) (videos []models.Video, err *models.CustomErr)
}

type AnnotationStore interface {
	CreateAnnotation(ctx context.Context, annotation *models.Annotation) (err *models.CustomErr)
	UpdateAnnotation(ctx context.Context, annotation *models.Annotation) (err *models.CustomErr)
	DeleteAnnotation(ctx context.Context, annotation models.Annotation) (err *models.CustomErr)
	GetAnnotation(ctx context.Context, annotationID string) (annotation models.Annotation, err *models.CustomErr)
}

type Store struct {
	VideoStore      VideoStore
	AnnotationStore AnnotationStore
}
