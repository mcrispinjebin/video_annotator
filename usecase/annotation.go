package usecase

import (
	"context"
	"fmt"
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

func (a annotationUsecase) CreateAnnotation(ctx context.Context, annotation *models.Annotation) (err error) {
	video, err := a.VideoStore.GetVideoByID(ctx, annotation.VideoID, true)
	if err != nil {
		return err
	}

	err = a.validateAnnotationDuration(ctx, annotation, &video)
	if err != nil {
		return err
	}

	if err = a.AnnotationStore.CreateAnnotation(ctx, annotation); err != nil {
		return err
	}
	return nil
}

func (a annotationUsecase) UpdateAnnotation(ctx context.Context,
	videoID string, annotationUpdate *models.Annotation) (fetchedAnnotation models.Annotation, err error) {

	video, err := a.VideoStore.GetVideoByID(ctx, videoID, true)
	if err != nil {
		return
	}

	fetchedAnnotation, err = a.AnnotationStore.GetAnnotation(ctx, annotationUpdate.ID)
	if err != nil {
		return
	}

	if fetchedAnnotation.ID == "" {
		err = fmt.Errorf("annotation ID does not exists for the video")
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
	video *models.Video) (err error) {

	startTime := annotation.StartTimeSec
	endTime := annotation.EndTimeSec

	if startTime < 0 || !(startTime < endTime) {
		//if start time is forced to be > 0, end time will be greater than 0 due to second condition
		err = fmt.Errorf("duration of annotation is invalid")
		return err
	}

	if endTime > video.DurationSec {
		err = fmt.Errorf("annotation duration exceeds video duration")
		return err
	}

	if annotation.Type == "" {
		err = fmt.Errorf("validation failed for annotation type")
		return err
	}

	for _, existingAnnotation := range video.Annotations {
		if existingAnnotation.StartTimeSec == startTime &&
			existingAnnotation.EndTimeSec == endTime &&
			existingAnnotation.Type == annotation.Type &&
			existingAnnotation.ID != annotation.ID {
			err = fmt.Errorf("another annotation with same type exists with the same duration")
			return err
		}
	}

	return nil
}

func (a annotationUsecase) DeleteAnnotation(ctx context.Context, videoID, annotationID string) (err error) {
	_, err = a.VideoStore.GetVideoByID(ctx, videoID, false)
	if err != nil {
		err = fmt.Errorf("video does not exist for given ID %q", err.Error())
		return
	}

	fetchedAnnotation, err := a.AnnotationStore.GetAnnotation(ctx, annotationID)
	if err != nil {
		return
	}

	if fetchedAnnotation.VideoID != videoID {
		err = fmt.Errorf("annotation does not exist for the given video ID %q", err.Error())
		return
	}

	if err = a.AnnotationStore.DeleteAnnotation(ctx, fetchedAnnotation); err != nil {
		return err
	}

	return nil
}
