package main

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

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Video Annotator service is here"))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler)
	//db := initPostgres()
	////initDB()
	//err := db.AutoMigrate(&Video{})

	dsn := "host=localhost user=postgres password=postgres dbname=video_annotation port=5432 sslmode=disable"

	db := store.ConnectPostgresDB(dsn)
	fmt.Printf("%#v", db)
	db.AutoMigrate(&models.Video{}, &models.Annotation{})
	//db.Migrator().CreateConstraint(&Video{}, "Annotation")
	//db.Migrator().CreateConstraint(&Video{}, "fk_Videos_annotations")

	store := store.Store{
		VideoStore:      store.NewVideoStore(db),
		AnnotationStore: store.NewAnnotationStore(db),
	}

	videoUsecase := usecase.NewVideoUsecase(store)
	annotationUsecase := usecase.NewAnnotationUsecase(store)

	h := handlers.NewHandler(videoUsecase, annotationUsecase)

	router.HandleFunc("/videos", h.CreateVideo).Methods("POST")
	router.HandleFunc("/videos/{videoID}", h.DeleteVideo).Methods("DELETE")
	router.HandleFunc("/videos/{videoID}/annotations", h.GetVideo).Methods("GET")

	router.HandleFunc("/videos/{videoID}/annotations",
		h.CreateAnnotation).Methods("POST")
	router.HandleFunc("/videos/{videoID}/annotations/{annotationID}",
		h.UpdateAnnotation).Methods("PATCH")
	router.HandleFunc("/videos/{videoID}/annotations/{annotationID}",
		h.DeleteAnnotation).Methods("DELETE")

	log.Print("Server is starting up!!")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
