package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	Id             *uuid.UUID `json:"id"`
	Artist         string     `json:"artist"`
	File_name      string     `json:"file_name"`
	Published_date string     `json:"published_date"`
	Title          string     `json:"title"`
}
