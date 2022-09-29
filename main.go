package main

// "github.com/Anthonyntilelli/go-fish/deck"
import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Anthonyntilelli/go-fish/deck"
	"github.com/Anthonyntilelli/go-fish/player"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	d := deck.New()
	p1 := player.New(1)
	for i := 0; i != 24; i++ {
		c, _ := d.DrawCard()
		p1.InsertCard(c)
	}
	fmt.Println(p1.Hand())
	fmt.Println(p1.RemoveCards("2"))
	fmt.Println(p1.Hand())
}
