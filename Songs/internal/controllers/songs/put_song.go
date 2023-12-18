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
// @Param        songRequest 	body  		models.SongRequest 	true	"song request"
// @Success      200            {object}  models.Song
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [put]
func PutSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songID, _ := ctx.Value("songID").(uuid.UUID)

	var newSongData models.SongRequest
	err := json.NewDecoder(r.Body).Decode(&newSongData)
	if err != nil {
		logrus.Errorf("Data decoding error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	song, err := songs.ModifySong(songID, newSongData)
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
