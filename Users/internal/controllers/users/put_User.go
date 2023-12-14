// post : on envoie dans le body de la requete du json 
//le post il crée un utilisateur
package users

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/users"
	"net/http"
	"io/ioutil"
)

// Put
// @Accept       json
// @Produce      json
// @Tags         Users
// @Summary      Put a User.
// @Description  Put a User.
// @Param        id           	 path      string  true  "Collection UUID formatted ID"
// @Param        body          		body   models.User    true  "User"
// @Success      200             {object}  models.User
// @Failure      404             "User not found"
// @Failure      500             "Something went wrong"
// @Failure      422             "missing fields"
// @Failure      409             "User already exists"
// @Router       /users/{id} [put]
func PutUser(w http.ResponseWriter, r *http.Request) {

	
	ctx := r.Context() // context de la requête (on met l'id dans l'url)
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID) // uuid -> type de données, CollectionId est un nom que l'on crée

	//il faut recuperer le format json dans le body de la requete 
	body, err := ioutil.ReadAll(r.Body)//-> lit le body et renvoie des données
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
	 // renvoie une chaine de caractère avec du json -> com pas au bon endroit
	 if err != nil {
        panic(err)
    }
	var t models.User
	err = json.Unmarshal(body, &t)// je veux une structure go -> a refaire 
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

	t.Id = &collectionId

	collections, err := users.PutAUser(t) // qui se retrouve dans repository/service
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
	w.WriteHeader(http.StatusOK)
	body, _ = json.Marshal(collections)// je le veux en json 
	_, _ = w.Write(body) // Bytes()
	return
}
