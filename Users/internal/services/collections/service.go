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
	"time"
)

func GetAllCollections() ([]models.User, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllCollections()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collections, nil
}

func GetCollectionById(id uuid.UUID) (*models.User, error) {
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

func PostAUser(user models.User) (*models.User, error) { // structure -> models.User
	var err error
	// calling repository
	// il creer un id 
	if user.Name == "" || user.Username == ""{
        return nil, &models.CustomError{
            Message: "missing fields",
            Code:    http.StatusUnprocessableEntity,
        }
    }
	id, err := uuid.NewV4()// génrer l'id
	user.Id= &id
	user.DateInscription=time.Now()
	err =repository.PostAUser(user)
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return repository.GetCollectionById(id)
}

func DeleteUserById(id uuid.UUID) ( error) {
	err := repository.DeleteUserById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return  &models.CustomError{ // on peut renvoyer un nil que avec un pointeur 
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return  &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return err
}

func PutAUser(user models.User) (*models.User, error) { // structure -> models.User
	var err error
	// calling repository
	// il creer un id 
	if user.Name == "" && user.Username == ""{
        return nil, &models.CustomError{
            Message: "missing fields",
            Code:    http.StatusUnprocessableEntity,
        }
    }
	err =repository.PutAUser(user)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return  nil,&models.CustomError{ // on peut renvoyer un nil que avec un pointeur 
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
	}
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return repository.GetCollectionById(*user.Id)
}