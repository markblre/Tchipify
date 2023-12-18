package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Song struct {
	Id            uuid.UUID `json:"id"`
	Artist        string    `json:"artist"`
	FileName      string    `json:"file_name"`
	PublishedDate time.Time `json:"published_date"`
	Title         string    `json:"title"`
}

type SongRequest struct {
	Artist   *string `json:"artist"`
	FileName *string `json:"file_name"`
	Title    *string `json:"title"`
}
