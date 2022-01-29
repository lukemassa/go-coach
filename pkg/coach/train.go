package coach

type Player struct{}

// Train given an environment, create a player
// that can play well in that environment
func Train(env Environment) Player {
	return Player{}
}

// Play a single episode with a player, returning it score
func (p *Player) Play(env Environment) int {
	return 0
}
