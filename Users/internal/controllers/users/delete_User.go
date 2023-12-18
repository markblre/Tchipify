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
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Summary      Delete a User.
// @Description  Delete a User.
// @Success      200           
// @Failure      404             "User not found"
// @Failure      500             "Something went wrong"
// @Router        /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // context de la requête (on met l'id dans l'url)
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID) 

	err := users.DeleteUserById(collectionId)
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
