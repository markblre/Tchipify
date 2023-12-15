package ratings

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/ratings"
	"net/http"
)

func GetAllRatingsBySongId(songId uuid.UUID) ([]models.Rating, error) {
	var err error
	// calling repository
	ratings, err := repository.GetAllRatingsBySongId(songId)
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving ratings : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return ratings, nil
}

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	rating, err := repository.GetRatingById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "rating not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving ratings : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return rating, err
}
