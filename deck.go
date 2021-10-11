package card

import (
	"math/rand"
	"time"
)

// Deck describes the deck.
// the struct contains a slice of a deck of cards
type Deck struct {
	cards []Card
}

// NewDeck generates a deck of cards and returns it as a struct
func NewDeck() (deck *Deck) {
	// suits - an array of all suits of cards
	suits := [NumberOfSuits]Suit{Spades, Hearts, Clubs, Diamonds}

	// values - an array of all values of cards
	values := [NumberOfValues]Value{Two, Three, Four, Five, Six,
		Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	// deck - pointer to a new Deck structure object
	deck = new(Deck)
	deck.cards = make([]Card, NumberOfValues*NumberOfSuits)

	// current - a variable for the current card in the deck
	current := 0

	for v := 0; v < NumberOfValues; v++ {
		for s := 0; s < NumberOfSuits; s++ {
			deck.cards[current].Value = values[v]
			deck.cards[current].Suit = suits[s]
			current += 1
		}
	}
	return deck
}

// RandomDeck shuffles the deck of cards
func (d *Deck) RandomDeck() {
	// function for randomness
	rand.Seed(time.Now().UnixNano())

	for key := range d.cards {
		randInt := rand.Intn(len(d.cards))
		d.cards[key], d.cards[randInt] = d.cards[randInt], d.cards[key]
	}
}

func (d *Deck) TakeTop() (card Card) {
	card = d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card
}

func (d *Deck) TakeBottom() (card Card) {
	card = d.cards[0]
	d.cards = d.cards[1:]
	return card
}