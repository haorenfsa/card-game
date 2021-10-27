package pkg

// Stack card stack
type Stack struct {
	Cards    []*Card
	CardsNum int
}

func NewStack(deckNum int) *Stack {
	if deckNum < 0 {
		panic("bad deckNum")
	}
	if deckNum == 0 {
		return &Stack{Cards: make([]*Card, 0)}
	}
	var cards = []*Card{}
	num := 0
	for i := 0; i < deckNum; i++ {
		deck := NewDeck()
		cards = append(cards, deck.Cards...)
		num += ONE_DECK_CARDS_NUM
	}
	return &Stack{
		Cards:    cards,
		CardsNum: num,
	}
}
