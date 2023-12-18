package songs

import (
	"database/sql"
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

func GetSong(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSong(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
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

func AddSong(songRequest models.SongRequest) (*models.Song, error) {
	if songRequest.Artist == nil || songRequest.File_name == nil || songRequest.Title == nil {
		return nil, &models.CustomError{
			Message: "missing fields",
			Code:    http.StatusUnprocessableEntity,
		}
	}

	id, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("error creating uuid : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	newSong := models.Song{
		Id:             id,
		Artist:         *songRequest.Artist,
		File_name:      *songRequest.File_name,
		Published_date: time.Now(),
		Title:          *songRequest.Title,
	}

	err = repository.AddSong(newSong)
	if err != nil {
		logrus.Errorf("Error adding song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return &newSong, err
}

func ModifySong(songId uuid.UUID, newSongData models.SongRequest) (*models.Song, error) {
	song, err := repository.ModifySong(songId, newSongData)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("Error modifying and retrieving song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return song, err
}

func DeleteSong(id uuid.UUID) error {
	err := repository.DeleteSong(id)
	if err != nil {
		logrus.Errorf("Error deleting song : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return err
}
