package models

import (
	"github.com/gofrs/uuid"
)

type Rating struct {
	Id      *uuid.UUID `json:"id"`
	Content string     `json:"content"`
}
