package app

import (
	"fmt"
	"log"
	"video_annotator/handlers"
	"video_annotator/models"
	"video_annotator/store"
	"video_annotator/usecase"
)

func Start() {

	//add env and defer panic exception codes

	dsn := "host=postgres user=postgres password=postgres dbname=video_annotation port=5432 sslmode=disable"

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
	SetupRoutes(h)
}
