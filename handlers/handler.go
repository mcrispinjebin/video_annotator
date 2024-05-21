package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"video_annotator/constants"
	"video_annotator/models"
	"video_annotator/usecase"
	"video_annotator/utils"
)

type handler struct {
	VideoUsecase      usecase.VideoUsecase
	AnnotationUsecase usecase.AnnotationUsecase
}

func NewHandler(videoUsecase usecase.VideoUsecase, annotationUsecase usecase.AnnotationUsecase) Handler {
	return &handler{VideoUsecase: videoUsecase, AnnotationUsecase: annotationUsecase}
}

type Handler interface {
	CreateVideo(w http.ResponseWriter, r *http.Request)
	GetVideo(w http.ResponseWriter, r *http.Request)
	DeleteVideo(w http.ResponseWriter, r *http.Request)

	CreateAnnotation(w http.ResponseWriter, r *http.Request)
	UpdateAnnotation(w http.ResponseWriter, r *http.Request)
	DeleteAnnotation(w http.ResponseWriter, r *http.Request)
}

func (h *handler) CreateVideo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	video := &models.Video{}
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.VideoCreateJSONDecodeErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(video)
	if err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.VideoCreateRequestValidationErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	if cErr := h.VideoUsecase.CreateVideo(ctx, video); cErr != nil {
		utils.ErrorResponse(w, cErr)
		return
	}

	utils.ReturnResponse(w, constants.HttpStatusOK, video)
	return
}

func (h *handler) GetVideo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.VideoGetIDParamErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	video, cErr := h.VideoUsecase.GetVideo(ctx, videoID)
	if cErr != nil {
		utils.ErrorResponse(w, cErr)
		return
	}

	utils.ReturnResponse(w, constants.HttpStatusOK, video)
	return
}

func (h *handler) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.VideoDeleteIDParamErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	cErr := h.VideoUsecase.DeleteVideo(ctx, videoID)
	if cErr != nil {
		utils.ErrorResponse(w, cErr)
		return
	}

	utils.ReturnResponse(w, constants.HttpStatusNoContent, nil)
	return
}

func (h *handler) CreateAnnotation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.AnnotationCreateURLParamErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	annotation := &models.Annotation{}
	if err = json.NewDecoder(r.Body).Decode(&annotation); err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.AnnotationCreateJSONDecodeErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	annotation.VideoID = videoID
	if cErr := h.AnnotationUsecase.CreateAnnotation(ctx, annotation); cErr != nil {
		utils.ErrorResponse(w, cErr)
		return
	}

	utils.ReturnResponse(w, constants.HttpStatusOK, annotation)
	return
}

func (h *handler) UpdateAnnotation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.AnnotationUpdateVideoIDParamErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	annotationID, err := utils.GetURLParam(r, "annotationID")
	if err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.AnnotationUpdateAnnotationIDParamErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	annotationUpdate := &models.Annotation{}
	if err := json.NewDecoder(r.Body).Decode(&annotationUpdate); err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.AnnotationUpdateDecodeErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	annotationUpdate.ID = annotationID
	updatedAnnotation, cErr := h.AnnotationUsecase.UpdateAnnotation(ctx, videoID, annotationUpdate)
	if cErr != nil {
		log.Printf("%s - %q", cErr.Message, cErr.Err)
		utils.ErrorResponse(w, cErr)
		return
	}

	utils.ReturnResponse(w, constants.HttpStatusOK, updatedAnnotation)
	return
}

func (h *handler) DeleteAnnotation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.AnnotationDeleteVideoIDParamErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	annotationID, err := utils.GetURLParam(r, "annotationID")
	if err != nil {
		cErr := &models.CustomErr{Err: err, Message: constants.AnnotationDeleteAnnotationIDParamErr,
			StatusCode: constants.HttpStatusBadRequest}
		utils.ErrorResponse(w, cErr)
		return
	}

	cErr := h.AnnotationUsecase.DeleteAnnotation(ctx, videoID, annotationID)
	if cErr != nil {
		utils.ErrorResponse(w, cErr)
		return
	}

	utils.ReturnResponse(w, constants.HttpStatusNoContent, nil)
	return
}
