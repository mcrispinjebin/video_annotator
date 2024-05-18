package store

import (
	"context"
	"video_annotator/models"
)

type VideoStore interface {
	CreateNewVideo(ctx context.Context, video *models.Video) (err error)
	GetVideoByID(ctx context.Context, videoID string) (video models.Video, err error)
	DeleteVideo(ctx context.Context, video models.Video) (err error)
	GetAllVideos(ctx context.Context) (videos []models.Video, err error)
}

type AnnotationStore interface {
	CreateAnnotationByID(ctx context.Context, annotation *models.Annotation) (err error)
	UpdateAnnotationByID(ctx context.Context, id string, annotations *models.Annotation) (err error)
	DeleteAnnotationByID(ctx context.Context, id string) (err error)
}

type Store struct {
	VideoStore      VideoStore
	AnnotationStore AnnotationStore
}
