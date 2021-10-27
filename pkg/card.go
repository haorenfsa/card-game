package pkg

type CardSuit int

const (
	CardSuitSpade CardSuit = iota
	CardSuitHeart
	CardSuitDiamond
	CardSuitClub
	CardSuitJoker
)

var ALL_SUITS = []CardSuit{
	CardSuitSpade,
	CardSuitHeart,
	CardSuitDiamond,
	CardSuitClub,
	CardSuitJoker,
}

type Card struct {
	Suit   CardSuit
	Points int // 1-13; black joker 14; red joker 15
}

func NewCard(suit CardSuit, points int) *Card {
	// assertion
	if suit == CardSuitJoker {
		RuntimeAssert(points == 14 || points == 15, "bad joker points")
	} else {
		RuntimeAssert(points > 0 && points <= 13, "bad none joker points")
	}
	return &Card{Suit: suit, Points: points}
}

func RuntimeAssert(assertion bool, msg string) {
	if !assertion {
		panic(msg)
	}
}
