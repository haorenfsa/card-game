package pkg

type PokerGame interface {
	Start()
}

type MyGamePhase int

type MyGame struct {
	players   []*Player
	turnIndex int
	round     int
	stack     *Stack
}

func (g MyGame) Start() {

}
