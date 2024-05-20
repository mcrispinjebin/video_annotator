package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"video_annotator/handlers"
)

func SetupRoutes(h handlers.Handler) {
	router := mux.NewRouter()

	serviceRouter := router.PathPrefix("/video-annotator").Subrouter()

	serviceRouter.HandleFunc("/videos", WithProtectedAuth(h.CreateVideo)).Methods("POST")
	serviceRouter.HandleFunc("/videos/{videoID}", WithProtectedAuth(h.DeleteVideo)).Methods("DELETE")
	serviceRouter.HandleFunc("/videos/{videoID}/annotations", WithProtectedAuth(h.GetVideo)).Methods("GET")

	serviceRouter.HandleFunc("/videos/{videoID}/annotations", WithProtectedAuth(
		h.CreateAnnotation)).Methods("POST")
	serviceRouter.HandleFunc("/videos/{videoID}/annotations/{annotationID}", WithProtectedAuth(
		h.UpdateAnnotation)).Methods("PATCH")
	serviceRouter.HandleFunc("/videos/{videoID}/annotations/{annotationID}", WithProtectedAuth(
		h.DeleteAnnotation)).Methods("DELETE")

	log.Print("Server is starting up!!")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
