package users

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/gofrs/uuid"
	"middleware/example/internal/services/users"
	"net/http"
	"middleware/example/internal/models"
)

// 	DeleteUser
// @Accept       json
// @Produce      json
// @Tags         Users
// @Param        id           	path      string  true  "User UUID formatted ID"
// @Summary      Delete a User.
// @Description  Delete a User.
// @Success      204           
// @Failure      500             "Something went wrong"
// @Router        /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // context de la requête (on met l'id dans l'url)
	UserId, _ := ctx.Value("UserId").(uuid.UUID) 

	err := users.DeleteUserById(UserId)
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

	w.WriteHeader(http.StatusNoContent) 
	return
}
