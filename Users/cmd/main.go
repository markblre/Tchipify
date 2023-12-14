package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
	
)

func main() {
	r := chi.NewRouter() // crée le routeur 
//ce qui l'Api autorise de faire 
	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUsers)
		r.Post("/", users.PostUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(users.Ctx) // le routeur appel la fonction 
			r.Get("/", users.GetUser)
			r.Delete("/", users.DeleteUser)
			r.Put("/", users.PutUser)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() { // initialise la base de données 
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL UNIQUE,
			date_inscription DATE NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
