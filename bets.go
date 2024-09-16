package models

import "time"

type BetStorage interface {
	List() []*Bet
	Get(string) *Bet
	Update(string, Bet) *Bet
	Create(Bet)
	Delete(string) *Bet
}

type BetStore struct{}

type Bet struct {
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
var Bets = []*Bet{
	{
		ID:       "1",
		Name:     "Lark Davis",
		Birthday: time.Date(1990, 5, 18, 0, 0, 0, 0, time.UTC),
		Country:  country.Name,
		City:     &city,
	},
}

func (b BetStore) Get(id string) *Bet {
	for _, Bets := range Bets {
		if Bets.ID == id {
			return Bets
		}
	}

	return nil
}

func (b BetStore) List() []*Bet {
	return Bets
}

func (b BetStore) Create(book Bet) {
	Bets = append(Bets, &book)
}

func (b BetStore) Delete(id string) *Bet {
	for i, Bet := range Bets {
		if Bet.ID == id {
			Bets = append(Bets[:i], (Bets)[i+1:]...)
			return &Bet{}
		}
	}

	return nil
}

func (b BetStore) Update(id string, bookUpdate Bet) *Bet {
	for i, book := range Bets {
		if book.ID == id {
			Bets[i] = &bookUpdate
			return book
		}
	}

	return nil
}
