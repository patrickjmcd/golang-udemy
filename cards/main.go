package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	testHand := newDeckFromFile("remainingDeck.csv")
	testHand.print()
	testHand.shuffle(5000)
	testHand.print()
}
