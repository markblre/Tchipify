package collections
// le service appel le repository
// le service manage tout et devra générer les ids 
import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"
	"fmt"
)

func GetAllCollections() ([]models.Collection, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllCollections()
	// managing errors
	if err != nil {
		fmt.Println("cc,,")
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collections, nil
}

func GetCollectionById(id uuid.UUID) (*models.Collection, error) {
	collection, err := repository.GetCollectionById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{ // on peut renvoyer un nil que avec un pointeur 
				Message: "collection not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collection, err
}

func PostAUser(user models.Collection) (*models.Collection, error) { // structure -> models.Collection
	var err error
	// calling repository
	// il creer un id 
	id, err := uuid.NewV4()// génrer l'id
	user.Id= &id
	fmt.Println("%d", user.Id)
	err =repository.PostAUser(user)
	// managing errors
	if err != nil {
		fmt.Println("cc9")
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return repository.GetCollectionById(id)
}