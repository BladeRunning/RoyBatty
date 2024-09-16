package main

import (
	"encoding/json"
	"net/http"

	"RoyBatty/models"

	"github.com/go-chi/chi/v5"
)

type BetHandler struct {
	storage models.BetStorage
}

func (p BetHandler) ListBets(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(p.storage.List())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (p BetHandler) GetBet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Bet := p.storage.Get(id)

	if Bet == nil {
		http.Error(w, "Bet not found", http.StatusNotFound)
	}

	err := json.NewEncoder(w).Encode(Bet)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (p BetHandler) CreateBet(w http.ResponseWriter, r *http.Request) {
	var Bet models.Bet

	err := json.NewDecoder(r.Body).Decode(&Bet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.storage.Create(Bet)

	err = json.NewEncoder(w).Encode(Bet)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (p BetHandler) UpdateBet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var Bet models.Bet

	err := json.NewDecoder(r.Body).Decode(&Bet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBet := p.storage.Update(id, Bet)
	if updatedBet == nil {
		http.Error(w, "Bet not found", http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(updatedBet)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (p BetHandler) DeleteBet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Bet := p.storage.Delete(id)

	if Bet == nil {
		http.Error(w, "Bet not found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusNoContent)
}
