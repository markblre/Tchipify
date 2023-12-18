package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"
)

// DeleteSongRating
// @Tags         ratings
// @Summary      Delete a song rating.
// @Description  Delete a song rating.
// @Param        song_id        path      string  true  "Song UUID formatted ID"
// @Param        rating_id      path      string  true  "Rating UUID formatted ID"
// @Success      204            "No Content"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{song_id}/ratings/{rating_id} [delete]
func DeleteSongRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songID, _ := ctx.Value("songID").(uuid.UUID)
	ratingID, _ := ctx.Value("ratingID").(uuid.UUID)

	err := ratings.DeleteSongRating(songID, ratingID)
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
