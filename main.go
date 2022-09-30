package main

// "github.com/Anthonyntilelli/go-fish/deck"
import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Anthonyntilelli/go-fish/deck"
	"github.com/Anthonyntilelli/go-fish/player"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cards := make([]deck.Card, 5)
	d := deck.New()
	for i := 0; i != 5; i++ {
		c, err := d.DrawCard()
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, c)
	}

	p1, err := player.New(1, cards)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i != 24; i++ {
		c, _ := d.DrawCard()
		p1.InsertCard(c)
	}
	fmt.Println(p1.Hand())
	fmt.Println(p1.RemoveCards("2"))
	fmt.Println(p1.Hand())
}
