package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
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

func (a annotationStore) CreateAnnotation(_ context.Context, annotation *models.Annotation) (err error) {
	annotation.ID = uuid.New().String()
	result := a.DB.Create(annotation)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (a annotationStore) UpdateAnnotation(_ context.Context, annotation *models.Annotation) (err error) {
	result := a.DB.Updates(annotation)
	if result.Error != nil {
		err = result.Error
		return
	}
	return nil
}

func (a annotationStore) DeleteAnnotation(_ context.Context, annotation models.Annotation) (err error) {
	annotation.Active = false
	result := a.DB.Save(&annotation)
	if result.Error != nil {
		err = result.Error
		return
	}

	return nil
}

func (a annotationStore) GetAnnotation(_ context.Context, annotationID string) (annotation models.Annotation, err error) {
	result := a.DB.First(&annotation, "id = ?  AND active = ?", annotationID, true)
	if result.Error != nil {
		err = result.Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("no annotation is matched with given ID %q", err)
			return
		}
		return
	}
	return annotation, nil
}
