// Package contains players hand, points and mechanism for point accumulation.
//
// Use the `New()` function to ensure proper creation on the Player.
// Each player should have a different ID.
package player

import (
	"github.com/Anthonyntilelli/go-fish/deck"
)

type Player struct {
	Id     int
	hand   map[string][]deck.Card
	points int
}

// Returns a Player with the specified id, an empty hand and 0 points.
func New(id int) Player {
	var p Player
	p.Id = id
	p.points = 0
	p.hand = make(map[string][]deck.Card)
	return p
}

// Inserts a new card into the hand of the player. If the card is part of a 4 set,
// the cards are removed from the hand and a point is increased. If a 4 set is
// made, the card value is returned, else "" is returned
func (p *Player) InsertCard(c deck.Card) string {
	p.hand[c.Value] = append(p.hand[c.Value], c)
	if len(p.hand[c.Value]) == 4 {
		delete(p.hand, c.Value)
		p.points++
		return c.Value
	}
	return ""
}

// Represents the hand as a one line string
func (p *Player) Hand() string {
	str := "[ "
	for k, v := range p.hand {
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
func (p *Player) Points() int {
	return p.points
}

// Removed a set of cards from Players hand by value
// Returns cards removed from hand and If card was removed.
// card will be blank if player does not have that card value
func (p *Player) RemoveCards(cardValue string) ([]deck.Card, bool) {
	contents, ok := p.hand[cardValue]
	if !ok {
		return []deck.Card{}, false
	}
	delete(p.hand, cardValue)
	return contents, true
}
