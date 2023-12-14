package users

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/users"
	"net/http"
)

// GetUser
// @Accept       json
// @Produce      json
// @Tags         Users
// @Summary      Get a User.
// @Description  Get a User.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200            {object}  models.User
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) { // http.Request -> la requete http
																// ResponseWriter -> contient la reponse 
	ctx := r.Context() // context de la requête (on met l'id dans l'url)
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID) // uuid -> type de données, CollectionId est un nom que l'on crée

	collection, err := users.GetUserById(collectionId)
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
