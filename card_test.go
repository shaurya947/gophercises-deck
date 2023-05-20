package deck

import "fmt"

func ExampleCard() {
	fmt.Println(Card{Suit: Club, Rank: Jack})
	fmt.Println(Card{Suit: Joker, Rank: Jack})

	// Output:
	// Jack of Clubs
	// Joker
}
