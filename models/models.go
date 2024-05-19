package models

import "time"

type Video struct {
	ID          string       `gorm:"primaryKey" json:"id"`
	Title       string       `json:"title" validate:"required"`
	Description string       `json:"description" validate:"required"`
	DurationSec int          `json:"durationSeconds" validate:"required,gt=0"`
	Url         string       `json:"url" validate:"required,url"`
	Source      string       `json:"source" validate:"required"`
	CreatedAt   time.Time    `json:"createdAt"`
	Active      bool         `gorm:"default:true" json:"-"`
	Annotations []Annotation `gorm:"foreignKey:VideoID" json:"annotations"`
}

type Annotation struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	VideoID         string    `json:"VideoID"`
	StartTimeSec    int       `json:"StartTimeSeconds"`
	EndTimeSec      int       `json:"EndTimeSeconds"`
	Type            string    `json:"Type"`
	AdditionalNotes string    `json:"AdditionalNotes"`
	Active          bool      `gorm:"default:true" json:"-"`
	CreatedAt       time.Time `json:"CreatedAt"`
	UpdatedAt       time.Time `json:"UpdatedAt"`

	Video Video `gorm:"foreignKey:VideoID;references:ID" json:"-"`
}
