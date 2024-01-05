package users

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	
)
func GetAllUsers() ([]models.User, error) {
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
	Users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.Name, &data.Username, &data.DateInscription)
		if err != nil {
			return nil, err
		}
		Users = append(Users, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return Users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM Users WHERE id=?", id.String())
	helpers.CloseDB(db)
	var User models.User
	err = row.Scan(&User.Id, &User.Name, &User.Username, &User.DateInscription)
	if err != nil {
		return nil, err
	}
	return &User, err
}

func PostAUser(user models.User) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return  err
	}
	 _,err2 := db.Exec("INSERT INTO Users(id, name, username, date_inscription) VALUES (?,?,?,?)",user.Id, user.Name, user.Username, user.DateInscription)
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

func PutAUser(user models.User) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return  err
	}
	if (user.Name != ""){
		_,err = db.Exec("UPDATE Users SET name=? WHERE id= ?", user.Name, user.Id)
	}
	if (user.Username!= ""){
		_,err = db.Exec("UPDATE Users SET  username=? WHERE id= ? ", user.Username, user.Id)
	}
	helpers.CloseDB(db)

	return  err
}