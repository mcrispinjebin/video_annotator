package usecase

import (
	"context"
	"fmt"
	"testing"
	"video_annotator/constants"
	"video_annotator/mocks"
	"video_annotator/models"

	"github.com/stretchr/testify/assert"
)

func TestVideoUsecase_CreateVideo(t *testing.T) {
	tests := []struct {
		name       string
		ID         string
		Url        string
		Message    string
		StatusCode int
	}{
		{"same url error",
			"2",
			"youtube.com",
			fmt.Sprintf(constants.VideoWithSameUrlExistsErr, "youtube.com"),
			constants.HttpResourceExists,
		},
		{"happy flow",
			"3",
			"google.com",
			"",
			0,
		},
	}

	for _, test := range tests {
		videoStore := mocks.MockVideoStore{
			Videos: map[string]models.Video{"1": {ID: "1", Url: "youtube.com"}},
		}
		videoUsecaseTest := NewVideoUsecase(&videoStore)
		t.Run(test.name, func(t *testing.T) {
			cErr := videoUsecaseTest.CreateVideo(context.Background(), &models.Video{Url: test.Url, ID: test.ID})
			if test.Message != "" {
				assert.Equal(t, test.Message, cErr.Message)
				assert.Equal(t, test.StatusCode, cErr.StatusCode)
			} else {
				_, exists := videoStore.Videos[test.ID]
				assert.Equal(t, true, exists)
			}
		})
	}
}

func TestVideoUsecase_GetVideo(t *testing.T) {
	tests := []struct {
		name       string
		ID         string
		Message    string
		StatusCode int
	}{
		{"resource does not exists",
			"2",
			constants.VideoResourceNotFound,
			constants.HttpResourceNotFound,
		},
		{"resource exists",
			"1",
			"",
			0,
		},
	}

	for _, test := range tests {
		url := "youtube.com"
		videoStore := mocks.MockVideoStore{
			Videos: map[string]models.Video{"1": {ID: "1", Url: url}},
		}
		videoUsecaseTest := NewVideoUsecase(&videoStore)
		t.Run(test.name, func(t *testing.T) {
			video, cErr := videoUsecaseTest.GetVideo(context.Background(), test.ID)
			if test.Message != "" {
				assert.Equal(t, test.Message, cErr.Message)
				assert.Equal(t, test.StatusCode, cErr.StatusCode)
			} else {
				assert.Equal(t, url, video.Url)
			}
		})
	}
}

func TestVideoUsecase_DeleteVideo(t *testing.T) {
	tests := []struct {
		name       string
		ID         string
		Message    string
		StatusCode int
	}{
		{"resource does not exists",
			"2",
			constants.VideoResourceNotFound,
			constants.HttpResourceNotFound,
		},
		{"resource exists",
			"1",
			"",
			0,
		},
	}

	for _, test := range tests {
		url := "youtube.com"
		videoStore := mocks.MockVideoStore{
			Videos: map[string]models.Video{"1": {ID: "1", Url: url}},
		}
		videoUsecaseTest := NewVideoUsecase(&videoStore)
		t.Run(test.name, func(t *testing.T) {
			cErr := videoUsecaseTest.DeleteVideo(context.Background(), test.ID)
			if test.Message != "" {
				assert.Equal(t, test.Message, cErr.Message)
				assert.Equal(t, test.StatusCode, cErr.StatusCode)
			} else {
				assert.Equal(t, 0, len(videoStore.Videos))
			}
		})
	}
}
