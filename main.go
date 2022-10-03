package main

// "github.com/Anthonyntilelli/go-fish/deck"
import (
	"fmt"
	"log"
	"math/rand"
	"time"

	computerplayer "github.com/Anthonyntilelli/go-fish/computer_player"
	"github.com/Anthonyntilelli/go-fish/deck"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var cards []deck.Card
	d := deck.New()
	for i := 0; i != 5; i++ {
		c, err := d.DrawCard()
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, c)
	}

	cp, err := computerplayer.New(1, cards)
	if err != nil {
		log.Fatal(err)
	}

	// Ask BEFORE more cards added
	fmt.Println(cp.Ask("3"))
	fmt.Println(cp.Ask("8"))
	fmt.Println(cp.Ask("5"))

	for i := 0; i != 15; i++ {
		c, _ := d.DrawCard()
		_, err = cp.InsertCard(c)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(cp.DisplayHand())
	fmt.Println(cp.Guess())
}
