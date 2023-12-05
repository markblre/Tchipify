package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"net/http"
)

// GetCollection
// @Tags         collections
// @Summary      Get a collection.
// @Description  Get a collection.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200            {object}  models.Collection
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [get]
func GetCollection(w http.ResponseWriter, r *http.Request) { // http.Request -> la requete http
																// ResponseWriter -> contient la reponse 
	ctx := r.Context() // context de la requête (on met l'id dans l'url)
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID) // uuid -> type de données, CollectionId est un nom que l'on crée

	collection, err := collections.GetCollectionById(collectionId)
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

	w.WriteHeader(http.StatusOK) // -> met le statut de la requete 
	body, _ := json.Marshal(collection) // -> string en json
	_, _ = w.Write(body) // ca ecrit le body dans la réponse c'est à dire renvoie la collection 
	return
}
