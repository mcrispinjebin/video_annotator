package store

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"video_annotator/constants"
	"video_annotator/models"
)

// create interface for DB and not pass gorm.Db
type annotationStore struct {
	DB *gorm.DB
}

func NewAnnotationStore(db *gorm.DB) AnnotationStore {
	return &annotationStore{DB: db}
}

func (a annotationStore) CreateAnnotation(_ context.Context, annotation *models.Annotation) (err *models.CustomErr) {
	annotation.ID = uuid.New().String()
	result := a.DB.Create(annotation)
	if result.Error != nil {
		err.Err = result.Error
		err.Message = constants.AnnotationCreateErr
		return
	}
	return nil
}

func (a annotationStore) UpdateAnnotation(_ context.Context, annotation *models.Annotation) (err *models.CustomErr) {
	result := a.DB.Updates(annotation)
	if result.Error != nil {
		err.Err = result.Error
		err.Message = constants.AnnotationUpdateErr
		return
	}
	return nil
}

func (a annotationStore) DeleteAnnotation(_ context.Context, annotation models.Annotation) (err *models.CustomErr) {
	annotation.Active = false
	result := a.DB.Save(&annotation)
	if result.Error != nil {
		err.Err = result.Error
		err.Message = constants.AnnotationDeleteErr
		return
	}

	return nil
}

func (a annotationStore) GetAnnotation(_ context.Context, annotationID string) (annotation models.Annotation, err *models.CustomErr) {
	result := a.DB.First(&annotation, "id = ?  AND active = ?", annotationID, true)
	if result.Error != nil {
		err.Err = result.Error
		if errors.Is(err.Err, gorm.ErrRecordNotFound) {
			err.StatusCode = constants.HttpResourceNotFound
			err.Message = constants.AnnotationResourceNotFound
			return
		}
		err.Message = constants.AnnotationResourceNotFound
		return
	}
	return annotation, nil
}
