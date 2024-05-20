package app

import (
	"fmt"
	"log"
	"os"
	"video_annotator/handlers"
	"video_annotator/models"
	"video_annotator/store"
	"video_annotator/usecase"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal(fmt.Sprintf("env variable %s does not exist", key))
	}
	return value
}

func Start() {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		getEnv("DB_USER"),
		getEnv("DB_PASSWORD"),
		getEnv("DB_NAME"),
		getEnv("DB_HOST"),
		getEnv("DB_PORT"),
	)

	defer func() {
		if r := recover(); r != nil {
			log.Print("panic occurred")
		}
	}()

	db := store.ConnectPostgresDB(dsn)
	if err := db.AutoMigrate(&models.Video{}, &models.Annotation{}); err != nil {
		log.Fatal(fmt.Sprintf("error in DB migration %q", err.Error()))
	}

	repoStore := store.Store{
		VideoStore:      store.NewVideoStore(db),
		AnnotationStore: store.NewAnnotationStore(db),
	}

	videoUsecase := usecase.NewVideoUsecase(repoStore.VideoStore)
	annotationUsecase := usecase.NewAnnotationUsecase(repoStore)

	h := handlers.NewHandler(videoUsecase, annotationUsecase)
	SetupRoutes(h)
}
