package songs

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
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

func GetSong(id uuid.UUID) (*models.Song, error) {
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

func AddSong(newSong models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO songs (id, artist, file_name, published_date, title) VALUES (?, ?, ?, ?, ?);", newSong.Id.String(), newSong.Artist, newSong.File_name, newSong.Published_date, newSong.Title)
	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return nil
}

func ModifySong(songId uuid.UUID, newSongData models.SongRequest) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	if newSongData.Artist != nil {
		_, err = tx.Exec("UPDATE songs SET artist=? WHERE id=?;", &newSongData.Artist, songId.String())
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if newSongData.File_name != nil {
		_, err = tx.Exec("UPDATE songs SET file_name=? WHERE id=?;", &newSongData.File_name, songId.String())
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if newSongData.Title != nil {
		_, err = tx.Exec("UPDATE songs SET title=? WHERE id=?;", &newSongData.Title, songId.String())
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	row := tx.QueryRow("SELECT * FROM songs WHERE id=?", songId.String())
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

func DeleteSong(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM songs WHERE id=?;", id.String())
	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return err
}
