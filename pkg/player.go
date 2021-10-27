package pkg

type PlayerInterface interface {
	DrawCard(StackInterface)
}

type Player struct {
	hand StackInterface
}

func NewPlayer() *Player {
	return &Player{
		hand: NewStack(0),
	}
}

// DrawCard from a stack
func (p *Player) DrawCard(stack StackInterface) {
	card := stack.Draw()
	p.hand.Add(card)
}
