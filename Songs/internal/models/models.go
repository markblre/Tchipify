package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Song struct {
	Id             *uuid.UUID `json:"id"`
	Artist         string     `json:"artist"`
	File_name      string     `json:"file_name"`
	Published_date time.Time  `json:"published_date"`
	Title          string     `json:"title"`
}

type SongRequest struct {
	Artist    string `json:"artist"`
	File_name string `json:"file_name"`
	Title     string `json:"title"`
}
