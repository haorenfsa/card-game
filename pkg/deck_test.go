package pkg

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	assert.Len(t, deck.Cards, ONE_DECK_CARDS_NUM)
	log.Print(deck.Cards)
	assert.NotNil(t, deck.Cards[ONE_DECK_CARDS_NUM-1])
	assert.Equal(t, 15, deck.Cards[ONE_DECK_CARDS_NUM-1].Points)
	for _, card := range deck.Cards {
		assert.NotNil(t, card)
	}
}
