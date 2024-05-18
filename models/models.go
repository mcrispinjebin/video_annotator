package models

import "time"

type Video struct {
	ID          string       `gorm:"primaryKey" json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	DurationSec int          `json:"durationSeconds"`
	Url         string       `json:"url"`
	Source      string       `json:"source"`
	CreatedAt   time.Time    `json:"createdAt"`
	Annotations []Annotation `gorm:"foreignKey:VideoID" json:"annotations"`
}

type Annotation struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	VideoID         string    `json:"VideoID"`
	StartTimeSec    int       `json:"StartTimeSeconds"`
	EndTimeSec      int       `json:"EndTimeSeconds"`
	Type            string    `json:"Type"`
	AdditionalNotes string    `json:"AdditionalNotes"`
	Active          bool      `gorm:"default:true" json:"Active"`
	CreatedAt       time.Time `json:"CreatedAt"`
	UpdatedAt       time.Time `json:"UpdatedAt"`

	Video Video `gorm:"foreignKey:VideoID;references:ID"`
}
