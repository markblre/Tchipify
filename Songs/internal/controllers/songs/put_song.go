package songs

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/songs"
	"net/http"
)

// PutSong
// @Tags         songs
// @Summary      Modify a song.
// @Description  Modify a song.
// @Param        id           	path      	string  true   "Song UUID formatted ID"
// @Param        artist         header      string  false  "Artist of the song"
// @Param        file_name      header      string  false  "Song file name"
// @Param        title          header      string  false  "Title of the song"
// @Success      200            {object}  models.Song
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [put]
func PutSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(uuid.UUID)

	var newSongData models.SongRequest
	err := json.NewDecoder(r.Body).Decode(&newSongData)
	if err != nil {
		logrus.Errorf("Data decoding error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	song, err := songs.PutSong(songId, newSongData)
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

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
	return
}
