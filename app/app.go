package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"video_annotator/handlers"
	"video_annotator/models"
	"video_annotator/store"
	"video_annotator/usecase"
)

func Start() {
	router := mux.NewRouter()

	dsn := "host=localhost user=postgres password=postgres dbname=video_annotation port=5432 sslmode=disable"

	db := store.ConnectPostgresDB(dsn)
	if err := db.AutoMigrate(&models.Video{}, &models.Annotation{}); err != nil {
		log.Fatal(fmt.Sprintf("error in DB migration %q", err.Error()))
	}

	repoStore := store.Store{
		VideoStore:      store.NewVideoStore(db),
		AnnotationStore: store.NewAnnotationStore(db),
	}

	videoUsecase := usecase.NewVideoUsecase(repoStore)
	annotationUsecase := usecase.NewAnnotationUsecase(repoStore)

	h := handlers.NewHandler(videoUsecase, annotationUsecase)

	router.HandleFunc("/videos", WithProtectedAuth(h.CreateVideo)).Methods("POST")
	router.HandleFunc("/videos/{videoID}", WithProtectedAuth(h.DeleteVideo)).Methods("DELETE")
	router.HandleFunc("/videos/{videoID}/annotations", WithProtectedAuth(h.GetVideo)).Methods("GET")

	router.HandleFunc("/videos/{videoID}/annotations", WithProtectedAuth(
		h.CreateAnnotation)).Methods("POST")
	router.HandleFunc("/videos/{videoID}/annotations/{annotationID}", WithProtectedAuth(
		h.UpdateAnnotation)).Methods("PATCH")
	router.HandleFunc("/videos/{videoID}/annotations/{annotationID}", WithProtectedAuth(
		h.DeleteAnnotation)).Methods("DELETE")

	log.Print("Server is starting up!!")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
