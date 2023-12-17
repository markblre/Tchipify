package ratings

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllTheRatingsForASongByItsID(songID uuid.UUID) ([]models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings WHERE song_id=?", songID.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	ratings := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.Id, &data.Comment, &data.Rating, &data.RatingDate, &data.SongID, &data.UserID)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return ratings, err
}

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM ratings WHERE id=?", id.String())
	helpers.CloseDB(db)

	var rating models.Rating
	err = row.Scan(&rating.Id, &rating.Comment, &rating.Rating, &rating.RatingDate, &rating.SongID, &rating.UserID)
	if err != nil {
		return nil, err
	}
	return &rating, err
}

func PostRating(newRating models.Rating) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO ratings (id, comment, rating, rating_date, song_id, user_id) VALUES (?, ?, ?, ?, ?, ?);", newRating.Id.String(), newRating.Comment, newRating.Rating, newRating.RatingDate, newRating.SongID.String(), newRating.UserID.String())
	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return nil
}
