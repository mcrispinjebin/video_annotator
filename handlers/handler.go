package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
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
		log.Printf("error in parsing json content for video creation %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(video)
	if err != nil {
		log.Printf("error in validating request %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.VideoUsecase.CreateVideo(ctx, video); err != nil {
		log.Printf("error occurred in creating video %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ReturnResponse(w, http.StatusOK, video)
	return
}

func (h *handler) GetVideo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		log.Printf("error occurred in creating video %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	video, err := h.VideoUsecase.GetVideo(ctx, videoID)
	if err != nil {
		log.Printf("error occurred in fetching video  %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ReturnResponse(w, http.StatusOK, video)
	return
}

func (h *handler) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		log.Printf("error in fetching video ID %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.VideoUsecase.DeleteVideo(ctx, videoID)
	if err != nil {
		log.Printf("error occurred in deleting video %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ReturnResponse(w, http.StatusNoContent, "")
	return
}

func (h *handler) CreateAnnotation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		log.Printf("error in fetching videoID %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	annotation := &models.Annotation{}
	if err := json.NewDecoder(r.Body).Decode(&annotation); err != nil {
		log.Printf("error in parsing json content for annotation creation %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	annotation.VideoID = videoID
	if err := h.AnnotationUsecase.CreateAnnotation(ctx, annotation); err != nil {
		log.Printf("error occurred in creating annotation %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ReturnResponse(w, http.StatusOK, annotation)
	return
}

func (h *handler) UpdateAnnotation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		log.Printf("error in fetching videoID %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	annotationID, err := utils.GetURLParam(r, "annotationID")
	if err != nil {
		log.Printf("error in fetching annotationID %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	annotationUpdate := &models.Annotation{}
	if err := json.NewDecoder(r.Body).Decode(&annotationUpdate); err != nil {
		log.Printf("error in parsing json content for annotation update %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	annotationUpdate.ID = annotationID
	updatedAnnotation, err := h.AnnotationUsecase.UpdateAnnotation(ctx, videoID, annotationUpdate)
	if err != nil {
		log.Printf("error occurred in updating annotation %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.ReturnResponse(w, http.StatusOK, updatedAnnotation)
	return
}

func (h *handler) DeleteAnnotation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	videoID, err := utils.GetURLParam(r, "videoID")
	if err != nil {
		log.Printf("error in fetching video ID %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	annotationID, err := utils.GetURLParam(r, "annotationID")
	if err != nil {
		log.Printf("error in fetching annotationID %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.AnnotationUsecase.DeleteAnnotation(ctx, videoID, annotationID)
	if err != nil {
		log.Printf("error occurred in deleting annotation %q", err.Error())
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ReturnResponse(w, http.StatusNoContent, "")
	return
}
