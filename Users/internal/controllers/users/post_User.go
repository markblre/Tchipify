
package users

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/users"
	"net/http"
	"io/ioutil"
)

// PostUser
// @Accept       json
// @Produce      json
// @Tags         Users
// @Summary      Post a User.
// @Description  Post a User.
// @Param        body          		body   models.UserRequest    true  "User"
// @Success      201            {object}  models.User
// @Failure      500             "Something went wrong"
// @Failure      422             "missing fields"
// @Router       /users [post]
func PostUser(w http.ResponseWriter, r *http.Request) {
	//il faut recuperer le format json dans le body de la requete 
	body, err := ioutil.ReadAll(r.Body)
	 if err != nil {
        panic(err)
    }
	var t models.User
	err = json.Unmarshal(body, &t)// je veux une structure go 
	Users, err := users.PostAUser(t) 
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
	userURL := "/users/" + Users.Id.String()
	w.Header().Set("Location", userURL)
	w.WriteHeader(http.StatusCreated)
	body, _ = json.Marshal(Users)
	_, _ = w.Write(body) 
	return
}
