package pkg

import (
	"math/rand"
	"sort"
	"time"
)

type StackInterface interface {
	// shuffle the stack of cards into random order
	Shuffle()
}

// Stack card stack
// implements sort.Interface for shuffle
type Stack struct {
	Cards    []*Card
	CardsNum int
}

func NewShuffledStack(deckNum int) *Stack {
	ret := NewStack(deckNum)
	ret.Shuffle()
	return ret
}

func NewStack(deckNum int) *Stack {
	if deckNum < 0 {
		panic("bad deckNum")
	}
	if deckNum == 0 {
		return &Stack{Cards: make([]*Card, 0)}
	}
	var cards = []*Card{}
	for i := 0; i < deckNum; i++ {
		deck := NewDeck()
		cards = append(cards, deck.Cards...)
	}
	return &Stack{
		Cards: cards,
	}
}

func init() {
	// for shuffle
	rand.Seed(time.Now().Unix())
}

func (s *Stack) Shuffle() {
	sort.Sort(s)
}

func (s *Stack) Len() int {
	return len(s.Cards)
}

func (s *Stack) Less(i, j int) bool {
	return rand.Int()%2 == 0
}

func (s *Stack) Swap(i, j int) {
	s.Cards[i], s.Cards[j] = s.Cards[j], s.Cards[i]
}
