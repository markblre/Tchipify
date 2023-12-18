package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Rating struct {
	Id         uuid.UUID `json:"id"`
	Comment    string    `json:"comment"`
	Rating     int       `json:"rating"`
	RatingDate time.Time `json:"rating_date"`
	SongID     uuid.UUID `json:"song_id"`
	UserID     uuid.UUID `json:"user_id"`
}

type RatingRequest struct {
	Comment *string `json:"comment"`
	Rating  *int    `json:"rating"`
	UserID  *string `json:"user_id"`
}
