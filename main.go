package main

import (
	"SCMR/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Mount("/perps", PerpRoutes())

	http.ListenAndServe(":3000", r)
}

func PerpRoutes() chi.Router {
	r := chi.NewRouter()

	perpHandler := PerpHandler{
		storage: models.PerpStore{},
	}

	r.Get("/", perpHandler.ListPerps)
	r.Post("/", perpHandler.CreatePerp)
	r.Get("/{id}", perpHandler.GetPerp)
	r.Put("/{id}", perpHandler.UpdatePerp)
	r.Delete("/{id}", perpHandler.DeletePerp)

	return r
}
