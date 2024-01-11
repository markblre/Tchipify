package users
import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/users"
	"net/http"
)

// GetUsers
// @Accept       json
// @Produce      json
// @Tags         Users
// @Summary      Get Users.
// @Description  Get Users.
// @Success      200            {array}  models.User
// @Failure      500             "Something went wrong"
// @Router       /users [get]
func GetUsers(w http.ResponseWriter, _ *http.Request) {
	// calling service
	Users, err := users.GetAllUsers()
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
	body, _ := json.Marshal(Users)
	_, _ = w.Write(body)
	return
}
