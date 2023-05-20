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
