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

func TestNewWithSortFunction(t *testing.T) {
	sortFn := func(cards []Card) func(i, j int) bool {
		// put all the 1s together, then all the 2s... all the Ks
		return func(i, j int) bool {
			return cards[i].Rank < cards[j].Rank
		}
	}

	cards := New(Sort(sortFn))

	require.Equal(t, Ace, cards[3].Rank)
	require.Equal(t, Two, cards[7].Rank)
	require.Equal(t, Three, cards[10].Rank)
}
