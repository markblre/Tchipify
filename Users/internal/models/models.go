package models

import (
	"github.com/gofrs/uuid"
)

type Collection struct {
	Id      *uuid.UUID `json:"id"`
	Name string     `json:"name"`
	Username string     `json:"username"`
	Date_inscription string     `json:"date_inscription"`
}


