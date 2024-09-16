package main

import (
	"RoyBatty/models"
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

	r.Mount("/Bets", BetRoutes())

	http.ListenAndServe(":3000", r)
}

func BetRoutes() chi.Router {
	r := chi.NewRouter()

	BetHandler := BetHandler{
		storage: models.BetStore{},
	}

	r.Get("/", BetHandler.ListBets)
	r.Post("/", BetHandler.CreateBet)
	r.Get("/{id}", BetHandler.GetBet)
	r.Put("/{id}", BetHandler.UpdateBet)
	r.Delete("/{id}", BetHandler.DeleteBet)

	return r
}
