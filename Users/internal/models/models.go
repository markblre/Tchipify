package models

import (
	"github.com/gofrs/uuid"
)
//normalement fini
type Collection struct {
	Id      *uuid.UUID `json:"id"`
	Name string     `json:"name"`
	Username string     `json:"username"`
	DateInscription string     `json:"date_inscription"`
}


