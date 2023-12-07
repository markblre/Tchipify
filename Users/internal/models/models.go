package models

import (
	"github.com/gofrs/uuid"
	"time"
)
//normalement fini
type Collection struct {
	Id      *uuid.UUID `json:"id"`
	Name string     `json:"name"`
	Username string     `json:"username"`
	DateInscription time.Time `json:"date_inscription"`
}


