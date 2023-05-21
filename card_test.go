package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Club, Rank: Jack})
	fmt.Println(Card{Suit: Joker, Rank: Jack})

	// Output:
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()

	require.Equal(t, Card{Suit: Spade, Rank: Five}, cards[4])
	require.Equal(t, Card{Suit: Diamond, Rank: Queen}, cards[24])
	require.Equal(t, Card{Suit: Heart, Rank: Six}, cards[44])
}

func TestNewSort(t *testing.T) {
	sortFn := func(cards []Card) func(i, j int) bool {
		// put all the 1s together, then all the 2s... all the Ks
		return func(i, j int) bool {
			return cards[i].Rank < cards[j].Rank
		}
	}

	cards := New(Sort(sortFn), Jokers(4))

	require.Equal(t, Ace, cards[3].Rank)
	require.Equal(t, Two, cards[7].Rank)
	require.Equal(t, Three, cards[10].Rank)
}

func TestNewJokers(t *testing.T) {
	cards := New(Jokers(4))

	require.Equal(t, len(cards), 56)
	require.Equal(t, Joker, cards[52].Suit)
	require.Equal(t, Joker, cards[53].Suit)
	require.Equal(t, Joker, cards[54].Suit)
	require.Equal(t, Joker, cards[55].Suit)
}

func TestNewFilter(t *testing.T) {
	keepFn := func(c Card) bool {
		// remove aces, twos and threes
		return c.Rank > Three
	}

	cards := New(Filter(keepFn))

	require.Equal(t, len(cards), 40)
	require.Equal(t, Four, cards[0].Rank)
	require.Equal(t, Five, cards[1].Rank)
	require.Equal(t, Six, cards[2].Rank)
}

func TestNewDecks(t *testing.T) {
	cards := New(Decks(3))

	require.Equal(t, len(cards), 156)
}
