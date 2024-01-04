
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
// @Param        id           	 path      string  true  "User UUID formatted ID"
// @Param        body          		body   models.UserRequest   true  "User"
// @Success      200             {object}  models.User
// @Failure      404             "User not found"
// @Failure      500             "Something went wrong"
// @Failure      422             "missing fields"
// @Failure      409             "User already exists"
// @Router       /users/{id} [put]
func PutUser(w http.ResponseWriter, r *http.Request) {

	
	ctx := r.Context() 
	UserId, _ := ctx.Value("UserId").(uuid.UUID) 
	body, err := ioutil.ReadAll(r.Body)
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
	 if err != nil {
        panic(err)
    }
	var t models.User
	err = json.Unmarshal(body, &t)
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

	t.Id = &UserId

	Users, err := users.PutAUser(t) 
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
	body, _ = json.Marshal(Users)
	_, _ = w.Write(body) 
	return
}
