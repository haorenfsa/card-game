package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStackPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			assert.Fail(t, "should panic")
		}
	}()
	NewStack(-1)
}

func TestNewStack(t *testing.T) {
	ret := NewStack(0)
	assert.NotNil(t, ret.Cards)
	assert.Zero(t, ret.CardsNum)
}

func TestNewStack1Deck(t *testing.T) {
	ret := NewStack(1)
	assert.Len(t, ret.Cards, ONE_DECK_CARDS_NUM)
	assert.Equal(t, ret.CardsNum, ONE_DECK_CARDS_NUM)
}

func TestNewStack2Deck(t *testing.T) {
	ret := NewStack(2)
	assert.Len(t, ret.Cards, ONE_DECK_CARDS_NUM*2)
	assert.Equal(t, ret.CardsNum, ONE_DECK_CARDS_NUM*2)
}
