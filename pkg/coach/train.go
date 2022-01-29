package coach

import "math/rand"

type Player struct{}

// Train given an environment, create a player
// that can play well in that environment
func Train(env Environment) Player {
	// TODO: Train the player, encode into it how to play
	return Player{}
}

// Play a single episode with a player, returning it score
func (p *Player) Play(env Environment) int {
	// Bespoke to konami for now as we figure out
	possibleActions := env.PossibleActions()
	for {

		env.Take(ActionIndex(rand.Intn(possibleActions)))
		if env.IsComplete() {
			break
		}
	}
	score := env.Score()
	env.Reset()
	return score
}
