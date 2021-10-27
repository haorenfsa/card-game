package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_DrawCard(t *testing.T) {
	player := NewPlayer()
	stack := NewStack(1)
	player.DrawCard(stack)
	assert.Equal(t, player.hand.CardsNum(), 1)
	assert.Equal(t, 1, player.hand.GetCardByIndex(0).Points)
}
