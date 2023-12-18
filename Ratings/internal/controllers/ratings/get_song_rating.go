package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"
)

// GetSongRating
// @Tags         ratings
// @Summary      Get a song rating.
// @Description  Get a song rating.
// @Param        song_id        path      string  true  "Song UUID formatted ID"
// @Param        rating_id      path      string  true  "Rating UUID formatted ID"
// @Success      200            {object}  models.Rating
// @Failure      404            "Rating not found"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{song_id}/ratings/{rating_id} [get]
func GetSongRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songID, _ := ctx.Value("songID").(uuid.UUID)
	ratingID, _ := ctx.Value("ratingID").(uuid.UUID)

	rating, err := ratings.GetSongRating(songID, ratingID)
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
	body, _ := json.Marshal(rating)
	_, _ = w.Write(body)
	return
}
