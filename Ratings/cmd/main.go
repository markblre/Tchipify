package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/ratings"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/songs/{song_id}", func(r chi.Router) {
		r.Route("/ratings", func(r chi.Router) {
			r.Use(ratings.CtxSongID)
			r.Get("/", ratings.GetSongRatings)
			r.Post("/", ratings.PostSongRating)
			r.Route("/{rating_id}", func(r chi.Router) {
				r.Use(ratings.CtxRatingID)
				r.Get("/", ratings.GetSongRating)
				r.Put("/", ratings.PutSongRating)
				r.Delete("/", ratings.DeleteSongRating)
			})
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8082")
	logrus.Fatalln(http.ListenAndServe(":8082", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS ratings (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			comment VARCHAR(255) NOT NULL,
    		rating INT NOT NULL,
    		rating_date DATE NOT NULL,
    		song_id VARCHAR(255) NOT NULL,
    		user_id VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
