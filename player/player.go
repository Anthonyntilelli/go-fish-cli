// Package contains players hand, points and mechanism for point accumulation.
//
// Use the `New()` function to ensure proper creation on the Player.
// Each player should have a different ID.
package player

import (
	"errors"

	"github.com/Anthonyntilelli/go-fish/deck"
)

type Player struct {
	Id     uint
	Hand   map[string][]deck.Card
	points uint
}

// Returns a Player with the specified id, a starting Hand and 0 points.
//
// An error is returned when the starting hand is empty
func New(id uint, starting_hand []deck.Card) (Player, error) {
	var p Player
	if id <= 0 || len(starting_hand) == 0 {
		return p, errors.New("invalid Player options")
	}
	p.Id = id
	p.points = 0
	p.Hand = make(map[string][]deck.Card)
	for _, c := range starting_hand {
		if _, err := p.InsertCard(c); err != nil {
			return Player{}, err
		}
	}
	return p, nil
}

// Inserts a new card into the hand of the player. If the card is part of a 4 set,
// the cards are removed from the hand and a point is increased. If a 4 set is
// made, the card value is returned, else "" is returned.
// Function will return error on empty card values
func (p *Player) InsertCard(c deck.Card) (string, error) {
	if c.Value == "" {
		return "", errors.New("attempting to an insert invalid card")
	}
	p.Hand[c.Value] = append(p.Hand[c.Value], c)
	if len(p.Hand[c.Value]) == 4 {
		delete(p.Hand, c.Value)
		p.points++
		return c.Value, nil
	}
	return "", nil
}

// Represents the hand as a one line string
func (p *Player) DisplayHand() string {
	str := "[ "
	for k, v := range p.Hand {
		str += k + ":"
		for _, val := range v {
			str += val.Suit
		}
		str += " "
	}
	str += "]"

	return str
}

// Prints current points
func (p *Player) Points() uint {
	return p.points
}

// Removed a set of cards from Players hand by value
//
// Returns cards removed from hand and If card was removed.
//
// Card will be blank if player does not have that card value.
func (p *Player) RemoveCards(cardValue string) ([]deck.Card, bool) {
	contents, ok := p.Hand[cardValue]
	if !ok {
		return []deck.Card{}, false
	}
	delete(p.Hand, cardValue)
	return contents, true
}
