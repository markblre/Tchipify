package collections
//on modifie rien sur ce fichier 
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"net/http"
)
//gère les requêtes http
func Ctx(next http.Handler) http.Handler { // c'est un middleware, routeur->middleware -> appel de la fonction 
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		collectionId, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {   //si il y il y a une erreur on fait ca :
			logrus.Errorf("parsing error : %s", err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse id (%s) as UUID", chi.URLParam(r, "id")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "collectionId", collectionId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
