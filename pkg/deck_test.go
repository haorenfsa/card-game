package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	assert.Len(t, deck.Cards, ONE_DECK_CARDS_NUM)
	assert.Equal(t, 15, deck.Cards[53].Points)
}
