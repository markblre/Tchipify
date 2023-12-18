package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type 	User struct {
	Id      *uuid.UUID `json:"id"`
	Name string     `json:"name"`
	Username string     `json:"username"`
	DateInscription time.Time `json:"date_inscription"`
}


