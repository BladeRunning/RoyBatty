package models

import "time"

type PerpStorage interface {
	List() []*Perp
	Get(string) *Perp
	Update(string, Perp) *Perp
	Create(Perp)
	Delete(string) *Perp
}

type PerpStore struct{}

type Perp struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Country  string    `json:"country"`
	City     *string   `json:"city"`
}

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var usa = Country{ID: 1, Name: "USA"}
var canada = Country{ID: 2, Name: "Canada"}
var country = Country{ID: 3, Name: "New Zealand"}

var city = "Wellington"
var perps = []*Perp{
	{
		ID:       "1",
		Name:     "Lark Davis",
		Birthday: time.Date(1990, 5, 18, 0, 0, 0, 0, time.UTC),
		Country:  country.Name,
		City:     &city,
	},
}

func (b PerpStore) Get(id string) *Perp {
	for _, perps := range perps {
		if perps.ID == id {
			return perps
		}
	}

	return nil
}

func (b PerpStore) List() []*Perp {
	return perps
}

func (b PerpStore) Create(book Perp) {
	perps = append(perps, &book)
}

func (b PerpStore) Delete(id string) *Perp {
	for i, perp := range perps {
		if perp.ID == id {
			perps = append(perps[:i], (perps)[i+1:]...)
			return &Perp{}
		}
	}

	return nil
}

func (b PerpStore) Update(id string, bookUpdate Perp) *Perp {
	for i, book := range perps {
		if book.ID == id {
			perps[i] = &bookUpdate
			return book
		}
	}

	return nil
}
