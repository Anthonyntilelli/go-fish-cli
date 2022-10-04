package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	computerplayer "github.com/Anthonyntilelli/go-fish/computer_player"
	"github.com/Anthonyntilelli/go-fish/deck"
	"github.com/Anthonyntilelli/go-fish/player"
)

// Draws count number of cards from the deck AND return drawn cards
func drawMultipleCards(count int, d *deck.Deck) []deck.Card {
	var cards []deck.Card
	for i := 0; i < count; i++ {
		c, err := d.DrawCard()
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, c)
	}
	return cards
}

// Draw card from the deck and adds it to the players hand.
// It will also print all related text to the terminal.
func goFish(d *deck.Deck, p *player.Player) {
	if d.IsDeckEmpty() {
		fmt.Println("The deck is empty.")
		return
	}
	pl := "Player " + strconv.Itoa(p.Id)
	fmt.Println(pl + " must GO FISH")
	c, _ := d.DrawCard()
	fmt.Println(pl + " picked up a " + c.Value + c.Suit)
	cv, _ := p.InsertCard(c)
	if cv != "" {
		fmt.Println(pl + " gained a Point :-)")
		fmt.Println("Point: " + strconv.Itoa(p.Points()))
	}
}

// transfers cards to players hand, and outputs points to terminal if needed.
func transferCards(p *player.Player, cards []deck.Card) {
	pl := "Player " + strconv.Itoa(p.Id)
	fmt.Println(pl + " guessed Correctly ")
	for _, c := range cards {
		cv, _ := p.InsertCard(c)
		if cv != "" {
			fmt.Println("You gained a Point :-)")
			fmt.Println("Points are now " + strconv.Itoa(p.Points()))
		}
	}
}

func main() {
	// Initialize Game
	rand.Seed(time.Now().UnixNano())
	deck := deck.New()

	humanPlayer, err := player.New(1, drawMultipleCards(5, &deck))
	if err != nil {
		log.Fatal(err)
	}

	compPlayer, err := computerplayer.New(2, drawMultipleCards(5, &deck))
	if err != nil {
		log.Fatal(err)
	}

	// Game Loop
	var input string
	for input != "exit" {
		fmt.Println("Your hand is: " + humanPlayer.DisplayHand())
		fmt.Println("Comp hand is: " + compPlayer.DisplayHand())
		fmt.Println("Enter a card value to ask other player (or type exit to leave): ")

		// Players turn
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}

		if !humanPlayer.CheckHand(input) {
			fmt.Println("///------------------------------------------------------------------------------------------------------///")
			fmt.Println("Error: Card Value must be in hand to ask for it.")
			continue
		}
		cards, ok := compPlayer.Ask(input)
		if ok { // Card found
			transferCards(&humanPlayer, cards)
		} else { // Go Fish
			goFish(&deck, &humanPlayer)
		}

		// // Computer Turn
		// guess := compPlayer.Guess()
		// fmt.Println("The Computer player Guesses " + guess)
		// cards, ok = humanPlayer.RemoveCards(guess)
		// if ok { // card found
		// 	fmt.Println("Your cards were removed from your hand.")
		// 	for _, c := range cards {
		// 		cv, _ := humanPlayer.InsertCard(c)
		// 		if cv != "" {
		// 			fmt.Println("Computer player gained a Point :-o")
		// 			fmt.Println("Points are now " + strconv.Itoa(compPlayer.Points()))
		// 		}
		// 	}
		// } else { // compPlayer GO FISH

		// }

		if (deck.IsDeckEmpty() && (compPlayer.EmptyHand() || humanPlayer.EmptyHand())) || (compPlayer.EmptyHand() && humanPlayer.EmptyHand()) {
			fmt.Println("GAME IS OVER!")
			if humanPlayer.Points() == compPlayer.Points() {
				fmt.Println("GAME IS A DRAW")
			} else if humanPlayer.Points() > compPlayer.Points() {
				fmt.Println("You won :-)")
			} else {
				fmt.Println("You Lost :-(")
			}
			break
		}

		fmt.Println("///------------------------------------------------------------------------------------------------------///")
	}

}
