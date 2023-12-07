package songs

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/songs"
	"net/http"
)

// DeleteSong
// @Tags         delete song
// @Summary      Delete a song.
// @Description  Delete a song.
// @Success      204            "No Content"
// @Failure      404            "Song not found"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /song/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(uuid.UUID)

	err := songs.DeleteSong(songId)
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

	w.WriteHeader(http.StatusNoContent)
	return
}
