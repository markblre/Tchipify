package songs

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/songs"
	"net/http"
)

// PostSong
// @Tags         post new song
// @Summary      Post a song.
// @Description  Post a song.
// @Success      200            {array}  models.Song
// @Failure      422             "missing fields"
// @Failure      500             "Something went wrong"
// @Router       /songs [post]
func PostSong(w http.ResponseWriter, r *http.Request) {
	var songRequest models.SongRequest
	err := json.NewDecoder(r.Body).Decode(&songRequest)
	if err != nil {
		logrus.Errorf("Data decoding error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	song, err := songs.PostSong(songRequest)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	songURL := "/songs/" + song.Id.String()
	w.Header().Set("Location", songURL)

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
	return
}
