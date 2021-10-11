package card

import (
	"math/rand"
	"time"
)

// Deck describes the deck.
// the struct contains a slice of a deck of Cards
type Deck struct {
	Cards []Card
}

// NewDeck generates a deck of Cards and returns it as a struct
func NewDeck() (deck *Deck) {
	// suits - an array of all suits of Cards
	suits := [NumberOfSuits]Suit{Spades, Hearts, Clubs, Diamonds}

	// values - an array of all values of Cards
	values := [NumberOfValues]Value{Two, Three, Four, Five, Six,
		Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	// deck - pointer to a new Deck structure object
	deck = new(Deck)
	deck.Cards = make([]Card, NumberOfValues*NumberOfSuits)

	// current - a variable for the current card in the deck
	current := 0

	for v := 0; v < NumberOfValues; v++ {
		for s := 0; s < NumberOfSuits; s++ {
			deck.Cards[current].Value = values[v]
			deck.Cards[current].Suit = suits[s]
			current += 1
		}
	}
	return deck
}

// RandomDeck shuffles the deck of Cards
func (d *Deck) RandomDeck() {
	// function for randomness
	rand.Seed(time.Now().UnixNano())

	for key := range d.Cards {
		randInt := rand.Intn(len(d.Cards))
		d.Cards[key], d.Cards[randInt] = d.Cards[randInt], d.Cards[key]
	}
}

func (d *Deck) TakeTop() (card Card) {
	card = d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card
}

func (d *Deck) TakeBottom() (card Card) {
	card = d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}