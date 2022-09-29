package deck

import (
	"errors"
	"math/rand"
)

const (
	Clubs    = "♣"
	Diamonds = "♦"
	Hearts   = "♥"
	Spades   = "♠"
)

type Card struct {
	Value, Suit string
}
type Deck struct {
	pile        [52]Card
	spot        int
	initialized bool
}

func New() Deck {
	var d Deck
	d.pile = [52]Card{
		{Value: "A", Suit: Clubs}, {Value: "2", Suit: Clubs}, {Value: "3", Suit: Clubs}, {Value: "4", Suit: Clubs}, {Value: "5", Suit: Clubs}, {Value: "6", Suit: Clubs}, {Value: "7", Suit: Clubs}, {Value: "8", Suit: Clubs}, {Value: "9", Suit: Clubs}, {Value: "10", Suit: Clubs}, {Value: "J", Suit: Clubs}, {Value: "Q", Suit: Clubs}, {Value: "K", Suit: Clubs},
		{Value: "A", Suit: Diamonds}, {Value: "2", Suit: Diamonds}, {Value: "3", Suit: Diamonds}, {Value: "4", Suit: Diamonds}, {Value: "5", Suit: Diamonds}, {Value: "6", Suit: Diamonds}, {Value: "7", Suit: Diamonds}, {Value: "8", Suit: Diamonds}, {Value: "9", Suit: Diamonds}, {Value: "10", Suit: Diamonds}, {Value: "J", Suit: Diamonds}, {Value: "Q", Suit: Diamonds}, {Value: "K", Suit: Diamonds},
		{Value: "A", Suit: Hearts}, {Value: "2", Suit: Hearts}, {Value: "3", Suit: Hearts}, {Value: "4", Suit: Hearts}, {Value: "5", Suit: Hearts}, {Value: "6", Suit: Hearts}, {Value: "7", Suit: Hearts}, {Value: "8", Suit: Hearts}, {Value: "9", Suit: Hearts}, {Value: "10", Suit: Hearts}, {Value: "J", Suit: Hearts}, {Value: "Q", Suit: Hearts}, {Value: "K", Suit: Hearts},
		{Value: "A", Suit: Spades}, {Value: "2", Suit: Spades}, {Value: "3", Suit: Spades}, {Value: "4", Suit: Spades}, {Value: "5", Suit: Spades}, {Value: "6", Suit: Spades}, {Value: "7", Suit: Spades}, {Value: "8", Suit: Spades}, {Value: "9", Suit: Spades}, {Value: "10", Suit: Spades}, {Value: "J", Suit: Spades}, {Value: "Q", Suit: Spades}, {Value: "K", Suit: Spades},
	}
	rand.Shuffle(len(d.pile), func(i, j int) {
		d.pile[i], d.pile[j] = d.pile[j], d.pile[i]
	})
	d.spot = 0
	d.initialized = true

	return d
}

func (d *Deck) DrawCard() (Card, error) {
	if !d.initialized {
		return Card{}, errors.New("Deck is not initialized")
	}
	if d.spot >= 52 || d.spot < 0 {
		return Card{}, errors.New("spot is out of range")
	}
	var card Card = d.pile[d.spot]
	d.spot++
	return card, nil
}

func (d *Deck) IsDeckEmpty() bool {
	return d.spot == 52
}
