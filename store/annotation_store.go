package store

import (
	"context"
	"gorm.io/gorm"
	"video_annotator/models"
)

// create interface for DB and not pass gorm.Db
type annotationStore struct {
	DB *gorm.DB
}

func NewAnnotationStore(db *gorm.DB) AnnotationStore {
	return &annotationStore{DB: db}
}

func (a annotationStore) CreateAnnotationByID(ctx context.Context, annotation *models.Annotation) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a annotationStore) UpdateAnnotationByID(ctx context.Context, id string, annotations *models.Annotation) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a annotationStore) DeleteAnnotationByID(ctx context.Context, id string) (err error) {
	//TODO implement me
	panic("implement me")
}
