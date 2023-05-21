package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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

func New(options ...func([]Card) []Card) []Card {
	cards := make([]Card, 52)

	counter := 0
	for suit := Spade; suit <= Heart; suit++ {
		for rank := Ace; rank <= King; rank++ {
			cards[counter] = Card{Rank: rank, Suit: suit}
			counter++
		}
	}

	for _, option := range options {
		cards = option(cards)
	}

	return cards
}

func Sort(less func([]Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Shuffle(cards []Card) []Card {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	newRand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}
