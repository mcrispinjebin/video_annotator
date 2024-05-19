package store

import (
	"context"
	"video_annotator/models"
)

type VideoStore interface {
	CreateNewVideo(ctx context.Context, video *models.Video) (err error)
	GetVideoByID(ctx context.Context, videoID string, includeAnnotations bool) (video models.Video, err error)
	DeleteVideo(ctx context.Context, video models.Video) (err error)
	GetAllVideos(ctx context.Context) (videos []models.Video, err error)
}

type AnnotationStore interface {
	CreateAnnotation(ctx context.Context, annotation *models.Annotation) (err error)
	UpdateAnnotation(ctx context.Context, annotation *models.Annotation) (err error)
	DeleteAnnotation(ctx context.Context, annotation models.Annotation) (err error)
	GetAnnotation(ctx context.Context, annotationID string) (annotation models.Annotation, err error)
}

type Store struct {
	VideoStore      VideoStore
	AnnotationStore AnnotationStore
}
