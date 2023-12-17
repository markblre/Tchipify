package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Rating struct {
	Id          uuid.UUID `json:"id"`
	Comment     string    `json:"comment"`
	Rating      int       `json:"rating"`
	Rating_date time.Time `json:"rating_date"`
	Song_id     uuid.UUID `json:"song_id"`
	User_id     uuid.UUID `json:"user_id"`
}

type RatingRequest struct {
	Comment *string `json:"comment"`
	Rating  *int    `json:"rating"`
	User_id *string `json:"user_id"`
}
