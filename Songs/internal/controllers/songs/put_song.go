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
// @Tags         put new song
// @Summary      Modify a song.
// @Description  Modify a song.
// @Success      200            {array}  models.Song
// @Failure      422             "missing fields"
// @Failure      422             "Cannot parse id"
// @Failure      500             "Something went wrong"
// @Router       /song/{id} [put]
func PutSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(uuid.UUID)

	var newSong models.Song
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		logrus.Errorf("Data decoding error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	song, err := songs.PutSong(songId, newSong.Artist, newSong.File_name, newSong.Title)
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

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
	return
}
