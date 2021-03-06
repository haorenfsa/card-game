package pkg

import (
	"math/rand"
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
}

func TestNewStack1Deck(t *testing.T) {
	ret := NewStack(1)
	assert.Len(t, ret.Cards, ONE_DECK_CARDS_NUM)
}

func TestNewStack2Deck(t *testing.T) {
	ret := NewStack(2)
	assert.Len(t, ret.Cards, ONE_DECK_CARDS_NUM*2)
}

func TestStackShuffle(t *testing.T) {
	rand.Seed(2)
	stack := NewStack(1)
	assert.True(t, stack.Cards[0].Points == 1)
	stack.Shuffle()
	assert.True(t, stack.Cards[0].Points != 1)
}

func TestNewShuffledStack(t *testing.T) {
	rand.Seed(2)
	stack := NewShuffledStack(1)
	assert.True(t, stack.Cards[0].Points != 1)
}

func TestStackDraw(t *testing.T) {
	stack := NewStack(1)
	top := stack.Cards[0]
	ret := stack.Draw()
	assert.Equal(t, top, ret)
	assert.Len(t, stack.Cards, ONE_DECK_CARDS_NUM-1)
}

func TestStack_GetCardByIndex(t *testing.T) {
	stack := NewStack(1)
	assert.Nil(t, stack.GetCardByIndex(ONE_DECK_CARDS_NUM))
	assert.NotNil(t, stack.GetCardByIndex(0))
	assert.NotNil(t, stack.GetCardByIndex(15))
	assert.NotNil(t, stack.GetCardByIndex(ONE_DECK_CARDS_NUM-1))
}
