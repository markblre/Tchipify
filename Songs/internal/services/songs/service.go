package songs

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/songs"
	"net/http"
	"time"
)

func GetAllSongs() ([]models.Song, error) {
	var err error
	// calling repository
	songs, err := repository.GetAllSongs()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

func PostSong(artist string, file_name string, title string) (*models.Song, error) {
	id, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("error creating uuid : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	if artist == "" || file_name == "" || title == "" {
		return nil, &models.CustomError{
			Message: "missing fields",
			Code:    http.StatusBadRequest,
		}
	}

	song, err := repository.PostSong(id, artist, file_name, time.Now(), title)
	if err != nil {
		logrus.Errorf("Error adding and retrieving song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return song, err
}
