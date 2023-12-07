package songs

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"time"
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Artist, &data.File_name, &data.Published_date, &data.Title)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Artist, &song.File_name, &song.Published_date, &song.Title)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func PostSong(id uuid.UUID, artist string, file_name string, published_date time.Time, title string) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("INSERT INTO songs (id, artist, file_name, published_date, title) VALUES (?, ?, ?, ?, ?);", id.String(), artist, file_name, published_date, title)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	row := tx.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	var song models.Song
	err = row.Scan(&song.Id, &song.Artist, &song.File_name, &song.Published_date, &song.Title)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	helpers.CloseDB(db)

	return &song, err
}

func PutSong(id uuid.UUID, artist string, file_name string, title string) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("UPDATE songs SET artist=?, file_name=?, title=? WHERE id=?;", artist, file_name, title, id.String())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	row := tx.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	var song models.Song
	err = row.Scan(&song.Id, &song.Artist, &song.File_name, &song.Published_date, &song.Title)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	helpers.CloseDB(db)

	return &song, err
}
