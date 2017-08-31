package main

func main() {
	//cards := newDeck()
	//
	//hand, remainingCards := deal(cards, 5)
	//
	//hand.print()
	//
	//remainingCards.print()

	//cards :=newDeck()
	//fmt.Println(cards.toString())
	//
	//cards.saveToFile("my-cards")

	//cards := newDeckFromFile("my")
	//cards.print()

	cards := newDeck()

	cards.shuffle()

	cards.print()

	//i:=1
	//remainingCards := cards
	//for i<5{
	//
	//	hand, remainingCards := deal(remainingCards,13)
	//	fmt.Println("Hand %v",i)
	//	hand.print()
	//	fmt.Println("Kalanlar")
	//	i++
	//	remainingCards.print()
	//}
}
