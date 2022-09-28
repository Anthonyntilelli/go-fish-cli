package deck

const (
	Clubs    = "♣"
	Diamonds = "♦"
	Hearts   = "♥"
	Spades   = "♠"
)

type Card struct {
	Value, Suit string
}
type Deck struct {
	Pile []Card
}

var Starter_deck = Deck{
	Pile: []Card{
		{Value: "A", Suit: Clubs}, {Value: "2", Suit: Clubs}, {Value: "3", Suit: Clubs}, {Value: "4", Suit: Clubs}, {Value: "5", Suit: Clubs}, {Value: "6", Suit: Clubs}, {Value: "7", Suit: Clubs}, {Value: "8", Suit: Clubs}, {Value: "9", Suit: Clubs}, {Value: "10", Suit: Clubs}, {Value: "J", Suit: Clubs}, {Value: "Q", Suit: Clubs}, {Value: "K", Suit: Clubs},
		{Value: "A", Suit: Diamonds}, {Value: "2", Suit: Diamonds}, {Value: "3", Suit: Diamonds}, {Value: "4", Suit: Diamonds}, {Value: "5", Suit: Diamonds}, {Value: "6", Suit: Diamonds}, {Value: "7", Suit: Diamonds}, {Value: "8", Suit: Diamonds}, {Value: "9", Suit: Diamonds}, {Value: "10", Suit: Diamonds}, {Value: "J", Suit: Diamonds}, {Value: "Q", Suit: Diamonds}, {Value: "K", Suit: Diamonds},
		{Value: "A", Suit: Hearts}, {Value: "2", Suit: Hearts}, {Value: "3", Suit: Hearts}, {Value: "4", Suit: Hearts}, {Value: "5", Suit: Hearts}, {Value: "6", Suit: Hearts}, {Value: "7", Suit: Hearts}, {Value: "8", Suit: Hearts}, {Value: "9", Suit: Hearts}, {Value: "10", Suit: Hearts}, {Value: "J", Suit: Hearts}, {Value: "Q", Suit: Hearts}, {Value: "K", Suit: Hearts},
		{Value: "A", Suit: Spades}, {Value: "2", Suit: Spades}, {Value: "3", Suit: Spades}, {Value: "4", Suit: Spades}, {Value: "5", Suit: Spades}, {Value: "6", Suit: Spades}, {Value: "7", Suit: Spades}, {Value: "8", Suit: Spades}, {Value: "9", Suit: Spades}, {Value: "10", Suit: Spades}, {Value: "J", Suit: Spades}, {Value: "Q", Suit: Spades}, {Value: "K", Suit: Spades},
	},
}
