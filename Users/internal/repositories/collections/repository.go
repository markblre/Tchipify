package collections

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	
)
//"database/sql" 
func GetAllCollections() ([]models.Collection, error) {
	//ouvrir la base de données 
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM Users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	collections := []models.Collection{}
	for rows.Next() {
		var data models.Collection
		err = rows.Scan(&data.Id, &data.Name, &data.Username, &data.DateInscription)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return collections, err
}

func GetCollectionById(id uuid.UUID) (*models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM Users WHERE id=?", id.String())
	helpers.CloseDB(db)
	var collection models.Collection
	err = row.Scan(&collection.Id, &collection.Name, &collection.Username, &collection.DateInscription)
	if err != nil {
		return nil, err
	}
	return &collection, err
}

func PostAUser(user models.Collection) (error) {// un seul User 
	//ouvrir la base de données 
	db, err := helpers.OpenDB()
	if err != nil {
		return  err
	}
	 _,err2 := db.Exec("INSERT INTO Users(id, name, username, date_inscription) VALUES (?,?,?,?)",user.Id, user.Name, user.Username, user.DateInscription)
	 //regarder si ca génère des erreurs 
	//La valeur 'sql.Result' fournit des informations supplémentaires sur l'opération de mise à jour ou d'insertion, telles que le nombre de lignes affectées ou le dernier ID inséré.
	helpers.CloseDB(db)

	return  err2
}

func DeleteUserById(id uuid.UUID) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return  err
	}
	_,err = db.Exec("DELETE FROM Users WHERE id=?", id.String())
	helpers.CloseDB(db)
	
	if err != nil {
		return  err
	}
	return err
}