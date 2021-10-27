package pkg

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

func init() {
	// for stack shuffle
	rand.Seed(time.Now().Unix())
}

type StackInterface interface {
	// Shuffle the stack of cards into random order
	Shuffle()
	// Draw a card from the stack top
	Draw() *Card
	// Add a card to stack top
	Add(*Card)
	// CardsNum of stack
	CardsNum() int
	// GetCardByIndex return card if exists, nil if not, panic if index < 0
	GetCardByIndex(int) *Card
}

// Stack card stack
// implements sort.Interface for shuffle
type Stack struct {
	Cards []*Card
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

func (s *Stack) Add(card *Card) {
	s.Cards = append(s.Cards, card)
}

func (s *Stack) Shuffle() {
	sort.Sort(s)
}

func (s *Stack) Draw() *Card {
	if len(s.Cards) < 1 {
		panic("no cards")
	}
	ret := s.Cards[0]
	s.Cards = s.Cards[1:]
	return ret
}

func (s *Stack) CardsNum() int {
	return len(s.Cards)
}

func (s *Stack) GetCardByIndex(i int) *Card {
	if i < 0 {
		panic("bad card index")
	}
	log.Print(s.CardsNum())
	if i <= s.CardsNum()-1 {
		log.Print(s.Cards[i])
		return s.Cards[i]
	}
	return nil
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
