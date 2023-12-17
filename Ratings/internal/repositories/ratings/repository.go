package ratings

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllRatingsBySongId(songId uuid.UUID) ([]models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings WHERE song_id=?", songId.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	ratings := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.Id, &data.Comment, &data.Rating, &data.Rating_date, &data.Song_id, &data.User_id)
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
	err = row.Scan(&rating.Id, &rating.Comment, &rating.Rating, &rating.Rating_date, &rating.Song_id, &rating.User_id)
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

	_, err = db.Exec("INSERT INTO ratings (id, comment, rating, rating_date, song_id, user_id) VALUES (?, ?, ?, ?, ?, ?);", newRating.Id.String(), newRating.Comment, newRating.Rating, newRating.Rating_date, newRating.Song_id.String(), newRating.User_id.String())
	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return nil
}
