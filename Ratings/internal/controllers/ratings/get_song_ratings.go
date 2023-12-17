package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"
)

// GetSongRatings
// @Tags         ratings
// @Summary      Get song ratings.
// @Description  Get song ratings.
// @Param        song_id        path      	string  true   "Song UUID formatted ID"
// @Success      200            {array}  models.Rating
// @Failure      500             "Something went wrong"
// @Router       /songs/{song_id}/ratings [get]
func GetSongRatings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songID, _ := ctx.Value("songId").(uuid.UUID)

	// calling service
	ratings, err := ratings.GetAllTheRatingsForASongByItsID(songID)
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(ratings)
	_, _ = w.Write(body)
	return
}
