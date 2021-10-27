package pkg

type Deck struct {
	Cards []*Card
}

const ONE_DECK_CARDS_NUM = 54
const ONE_SUIT_CARDS_NUM = 13

func NewDeck() *Deck {
	cards := make([]*Card, ONE_DECK_CARDS_NUM)
	index := 0
	for _, suit := range ALL_SUITS {
		if suit != CardSuitJoker {
			for points := 1; points <= ONE_SUIT_CARDS_NUM; points++ {
				card := NewCard(suit, points)
				cards[index] = card
				index++
			}
		} else {
			// add jokers
			card := NewCard(suit, 14)
			cards[index] = card
			index++
			card = NewCard(suit, 15)
			cards[index] = card
			index++
		}
	}
	return &Deck{
		Cards: cards,
	}
}
