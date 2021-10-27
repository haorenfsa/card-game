package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	NewCard(CardSuitSpade, 1)
	NewCard(CardSuitSpade, 13)
	NewCard(CardSuitDiamond, 1)
	NewCard(CardSuitDiamond, 13)
	NewCard(CardSuitJoker, 1)
	NewCard(CardSuitDiamond, 13)
}

func TestNewCardPanic(t *testing.T) {
	tests := []struct {
		suit   CardSuit
		points int
	}{
		// test cases.
		{CardSuitSpade, 0},
		{CardSuitSpade, -1},
		{CardSuitSpade, 14},
		{CardSuitDiamond, 15},
		{CardSuitJoker, 1},
		{CardSuitJoker, 16},
	}
	for _, test := range tests {
		func() {
			defer func() {
				if recover() == nil {
					assert.Fail(t, "should panic")
				}
			}()
			NewCard(test.suit, test.points)
		}()
	}
}
