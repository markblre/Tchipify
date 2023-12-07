package collections

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/gofrs/uuid"
	"middleware/example/internal/services/collections"
	"net/http"
	"middleware/example/internal/models"
)

// GetCollections
// @Tags         collections
// @Summary      Get collections.
// @Description  Get collections.
// @Success      200            {array}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /collections [get]

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // context de la requête (on met l'id dans l'url)
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID) // uuid -> type de données, CollectionId est un nom que l'on crée

	err := collections.DeleteUserById(collectionId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())// la doc de logrus est expliquée dans le tp
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
