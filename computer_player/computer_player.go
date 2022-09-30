// Computer player for the Go_fish game.  It relies on the player Struct
//
// Use `New()` function to initialize the struct
package computerplayer

import (
	"log"

	"github.com/Anthonyntilelli/go-fish/deck"
	"github.com/Anthonyntilelli/go-fish/player"
)

type ComputerPlayer struct {
	player.Player
	memory [13]bool // A 2 3 4 5 6 7 8 9 10 J Q K
}

// Returns a Computer player and error
//
// An error is returned when the starting hand is empty.
func New(player_id uint, starting_hand []deck.Card) (ComputerPlayer, error) {
	var cp ComputerPlayer
	var err error

	cp.Player, err = player.New(player_id, starting_hand)
	if err != nil {
		return cp, err
	}

	return cp, nil
}

// Ask if computer player has a particular card value.
//
// Returns the cards the player has in its hand and true if cards/ false if no card in hand.
func (cp *ComputerPlayer) Ask(cardValue string) ([]deck.Card, bool) {
	cards, found := cp.Player.RemoveCards(cardValue)
	cp.modifyMemory(cardValue, true)

	return cards, found
}

//TODO: GUESS and receive cards function

// Adds or removes the card value from memory, translates standard card values only
//
// Panics on invalid card
func (cp *ComputerPlayer) modifyMemory(cardValue string, has bool) {
	switch cardValue {
	case "A":
		cp.memory[0] = has
	case "2":
		cp.memory[1] = has
	case "3":
		cp.memory[2] = has
	case "4":
		cp.memory[3] = has
	case "5":
		cp.memory[4] = has
	case "6":
		cp.memory[5] = has
	case "7":
		cp.memory[6] = has
	case "8":
		cp.memory[7] = has
	case "9":
		cp.memory[8] = has
	case "10":
		cp.memory[9] = has
	case "J":
		cp.memory[10] = has
	case "Q":
		cp.memory[11] = has
	case "K":
		cp.memory[12] = has
	default:
		log.Fatalln("Invalid card called in modifyMemory function")
	}
}
