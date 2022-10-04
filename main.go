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
		if !ok { // Go Fish
			fmt.Println("Card not in other players hand.")
			if !deck.IsDeckEmpty() {
				fmt.Println("You must GO FISH")
				c, _ := deck.DrawCard()
				fmt.Println("You picked up a " + c.Value + c.Suit)
				cv, _ := humanPlayer.InsertCard(c)
				if cv != "" {
					fmt.Println("You gained a Point :-)")
					fmt.Println("Points are now " + strconv.Itoa(humanPlayer.Points()))
				}
			} else {
				fmt.Println("The deck is empty.")
			}
		} else { // Card found
			fmt.Println("The other player gave you their card(s)")
			for _, c := range cards {
				cv, _ := humanPlayer.InsertCard(c)
				if cv != "" {
					fmt.Println("You gained a Point :-)")
					fmt.Println("Points are now " + strconv.Itoa(humanPlayer.Points()))
				}
			}
		}

		// Computer Turn
		// guess := compPlayer.Guess()
		// cards, ok := humanPlayer.RemoveCards(guess)

		if deck.IsDeckEmpty() && (compPlayer.Hand == nil || humanPlayer.Hand == nil) {
			fmt.Println("GAME IS OVER!")
			if humanPlayer.Points() == compPlayer.Points() {
				fmt.Println("GAME IS A DRAW")
			} else if humanPlayer.Points() > compPlayer.Points() {
				fmt.Println("You won :-)")
			} else {
				fmt.Println("You Lost :-(")
			}
		}

		fmt.Println("///------------------------------------------------------------------------------------------------------///")
	}

}
