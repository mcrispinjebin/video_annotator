package usecase

import (
	"context"
	"video_annotator/models"
	"video_annotator/store"
)

type annotationUsecase struct {
	repoStore store.AnnotationStore
}

func NewAnnotationUsecase(store store.Store) AnnotationUsecase {
	return annotationUsecase{repoStore: store.AnnotationStore}
}

func (a annotationUsecase) CreateAnnotation(ctx context.Context, annotation *models.Annotation) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a annotationUsecase) UpdateAnnotation(ctx context.Context, id string, annotations *models.Annotation) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a annotationUsecase) DeleteAnnotation(ctx context.Context, id string) (err error) {
	//TODO implement me
	panic("implement me")
}
