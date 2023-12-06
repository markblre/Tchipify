// post : on envoie dans le body de la requete du json 
//le post il crée un utilisateur
package collections

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"net/http"
	"io/ioutil"
	"fmt"
)

// GetCollections
// @Tags         collections
// @Summary      Get collections.
// @Description  Get collections.
// @Success      200            {array}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /collections [get]

func PostUser(w http.ResponseWriter, r *http.Request) {
	//il faut recuperer le format json dans le body de la requete 
	body, err := ioutil.ReadAll(r.Body)//-> lit le body et renvoie des données
	if err != nil {
        fmt.Println("cc2")
    }
	 // renvoie une chaine de caractère avec du json -> com pas au bon endroit
	 if err != nil {
        panic(err)
    }
	var t models.Collection
	err = json.Unmarshal(body, &t)// je veux une structure go -> a refaire 
	if err != nil {
        fmt.Println("cc3")
    }
	collections, err := collections.PostAUser(t) // qui se retrouve dans repository/service
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
	fmt.Println("cc4")
	w.WriteHeader(http.StatusOK)
	body, _ = json.Marshal(collections)// je le veux en json 
	_, _ = w.Write(body) // Bytes()
	return
}
