// Computer player for the Go_fish game.  It relies on the player Struct
//
// Use `New()` function to initialize the struct
package computerplayer

import (
	"math/rand"

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
func New(player_id int, starting_hand []deck.Card) (ComputerPlayer, error) {
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

// Chooses a card value from memory and hand or select a card value at random from the hand
//
// It panics when hand is empty
func (cp *ComputerPlayer) Guess() string {
	var foundList []string
	var hnd []string

	if cp.EmptyHand() {
		panic("Cannot guess on an empty hand")
	}

	// Check Memory for cards in hand
	memory := cp.cardListFromMemory()
	for _, c := range memory {
		cValue := string(c)
		_, found := cp.Hand[cValue] // Checking if card value is in hand
		if found {
			foundList = append(foundList, cValue)
		}
	}

	// Item found in list
	if foundList != nil {
		i := rand.Intn(len(foundList)-0) + 0 // (MAX-MIN)-MIN
		cp.modifyMemory(foundList[i], false)
		return foundList[i]
	}
	// Nothing found in hand
	for k := range cp.Hand {
		hnd = append(hnd, k)
	}
	j := rand.Intn(len(hnd)-0) + 0 // (MAX-MIN)-MIN
	return hnd[j]
}

// Adds or removes the card value from memory, translates standard card values only
//
// Panics on invalid card (Capital Letters Only)
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
		panic("Invalid card called in modifyMemory function")
	}
}

// Creates a string based on memory
func (cp *ComputerPlayer) cardListFromMemory() string {
	cardList := ""

	if cp.memory[0] {
		cardList += "A"
	}
	if cp.memory[1] {
		cardList += "2"
	}
	if cp.memory[2] {
		cardList += "3"
	}
	if cp.memory[3] {
		cardList += "4"
	}
	if cp.memory[4] {
		cardList += "5"
	}
	if cp.memory[5] {
		cardList += "6"
	}
	if cp.memory[6] {
		cardList += "7"
	}
	if cp.memory[7] {
		cardList += "8"
	}
	if cp.memory[8] {
		cardList += "9"
	}
	if cp.memory[9] {
		cardList += "10"
	}
	if cp.memory[10] {
		cardList += "J"
	}
	if cp.memory[11] {
		cardList += "Q"
	}
	if cp.memory[12] {
		cardList += "K"
	}
	return cardList
}
