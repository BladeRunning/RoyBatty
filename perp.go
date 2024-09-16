package main

import (
	"encoding/json"
	"net/http"

	"SCMR/models"

	"github.com/go-chi/chi/v5"
)

type PerpHandler struct {
	storage models.PerpStorage
}

func (p PerpHandler) ListPerps(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(p.storage.List())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (p PerpHandler) GetPerp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	perp := p.storage.Get(id)

	if perp == nil {
		http.Error(w, "perp not found", http.StatusNotFound)
	}

	err := json.NewEncoder(w).Encode(perp)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (p PerpHandler) CreatePerp(w http.ResponseWriter, r *http.Request) {
	var perp models.Perp

	err := json.NewDecoder(r.Body).Decode(&perp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.storage.Create(perp)

	err = json.NewEncoder(w).Encode(perp)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (p PerpHandler) UpdatePerp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var perp models.Perp

	err := json.NewDecoder(r.Body).Decode(&perp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedperp := p.storage.Update(id, perp)
	if updatedperp == nil {
		http.Error(w, "perp not found", http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(updatedperp)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (p PerpHandler) DeletePerp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	perp := p.storage.Delete(id)

	if perp == nil {
		http.Error(w, "perp not found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusNoContent)
}
