package coach

type Environment interface {
	Reset()
}

type Player struct{}

func Train(env Environment) Player {
	return Player{}
}

func (p *Player) Play(env Environment) int {
	return 0
}
