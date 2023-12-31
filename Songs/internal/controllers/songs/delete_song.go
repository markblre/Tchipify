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
// @Tags         songs
// @Summary      Delete a song.
// @Description  Delete a song.
// @Param        id           	path      string  true  "Song UUID formatted ID"
// @Success      204            "No Content"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songID, _ := ctx.Value("songID").(uuid.UUID)

	err := songs.DeleteSong(songID)
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
