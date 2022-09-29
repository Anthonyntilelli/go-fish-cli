package main

// "github.com/Anthonyntilelli/go-fish/deck"
import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Anthonyntilelli/go-fish/deck"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	d := deck.New()
	fmt.Println(d)
	for i := 0; i != 55; i++ {
		fmt.Println(d.DrawCard())
	}
	fmt.Println(d.IsDeckEmpty())
}
