package usecase

import (
	"context"
	"video_annotator/constants"
	"video_annotator/models"
	"video_annotator/store"
)

type annotationUsecase struct {
	VideoStore      store.VideoStore
	AnnotationStore store.AnnotationStore
}

func NewAnnotationUsecase(store store.Store) AnnotationUsecase {
	return annotationUsecase{VideoStore: store.VideoStore, AnnotationStore: store.AnnotationStore}
}

func (a annotationUsecase) CreateAnnotation(ctx context.Context, annotation *models.Annotation) (err *models.CustomErr) {
	video, err := a.VideoStore.GetVideoByID(ctx, annotation.VideoID, true)
	if err != nil {
		return err
	}

	err = a.validateAnnotationDuration(ctx, annotation, &video)
	if err != nil {
		return err
	}

	//if create and update validation deviates much, separate validation functions
	if annotation.Type == "" {
		return &models.CustomErr{Message: constants.AnnotationTypeEmptyErr,
			StatusCode: constants.HttpStatusBadRequest}
	}

	if err = a.AnnotationStore.CreateAnnotation(ctx, annotation); err != nil {
		return err
	}
	return nil
}

func (a annotationUsecase) UpdateAnnotation(ctx context.Context,
	videoID string, annotationUpdate *models.Annotation) (fetchedAnnotation models.Annotation, err *models.CustomErr) {

	video, err := a.VideoStore.GetVideoByID(ctx, videoID, true)
	if err != nil {
		return
	}

	fetchedAnnotation, err = a.AnnotationStore.GetAnnotation(ctx, annotationUpdate.ID)
	if err != nil {
		return
	}

	if annotationUpdate.Type != "" {
		fetchedAnnotation.Type = annotationUpdate.Type
	}

	if annotationUpdate.AdditionalNotes != "" {
		fetchedAnnotation.AdditionalNotes = annotationUpdate.AdditionalNotes
	}

	if annotationUpdate.StartTimeSec > 0 && annotationUpdate.EndTimeSec > 0 {
		err = a.validateAnnotationDuration(ctx, annotationUpdate, &video)
		if err != nil {
			return
		}
		fetchedAnnotation.StartTimeSec = annotationUpdate.StartTimeSec
		fetchedAnnotation.EndTimeSec = annotationUpdate.EndTimeSec
	}

	if err = a.AnnotationStore.UpdateAnnotation(ctx, &fetchedAnnotation); err != nil {
		return
	}

	return fetchedAnnotation, nil

}

func (a annotationUsecase) validateAnnotationDuration(_ context.Context, annotation *models.Annotation,
	video *models.Video) (err *models.CustomErr) {

	startTime := annotation.StartTimeSec
	endTime := annotation.EndTimeSec

	err = &models.CustomErr{StatusCode: constants.HttpStatusBadRequest}
	if startTime < 0 {
		err.Message = constants.AnnotationStartTimePositiveErr
		return err
	}

	if !(startTime < endTime) {
		//if start time is forced to be > 0, end time will be greater than 0 due to this condition
		err.Message = constants.AnnotationEndTimeGreaterToStartTimeErr
		return err
	}

	if endTime > video.DurationSec {
		err.Message = constants.AnnotationDurationExceedsVideoErr
		return err
	}

	for _, existingAnnotation := range video.Annotations {
		if existingAnnotation.StartTimeSec == startTime &&
			existingAnnotation.EndTimeSec == endTime &&
			existingAnnotation.Type == annotation.Type &&
			existingAnnotation.ID != annotation.ID {
			err.Message = constants.AnnotationExistsWithSameDurationErr
			err.StatusCode = constants.HttpResourceExists
			return err
		}
	}

	return nil
}

func (a annotationUsecase) DeleteAnnotation(ctx context.Context, videoID, annotationID string) (err *models.CustomErr) {
	_, err = a.VideoStore.GetVideoByID(ctx, videoID, false)
	if err != nil {
		return
	}

	fetchedAnnotation, err := a.AnnotationStore.GetAnnotation(ctx, annotationID)
	if err != nil {
		return
	}

	if fetchedAnnotation.VideoID != videoID {
		return &models.CustomErr{Message: constants.AnnotationDoesNotExistForVideoErr,
			StatusCode: constants.HttpStatusBadRequest}
	}

	if err = a.AnnotationStore.DeleteAnnotation(ctx, fetchedAnnotation); err != nil {
		return err
	}

	return nil
}
