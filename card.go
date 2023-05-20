package deck

import "fmt"

//go:generate stringer -type Rank
type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

//go:generate stringer -type Suit
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Card struct {
	Rank
	Suit
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
	cards := make([]Card, 52)

	counter := 0
	for suit := Spade; suit <= Heart; suit++ {
		for rank := Ace; rank <= King; rank++ {
			cards[counter] = Card{Rank: rank, Suit: suit}
			counter++
		}
	}

	return cards
}
