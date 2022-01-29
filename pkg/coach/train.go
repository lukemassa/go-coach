package coach

import (
	"math/rand"
)

type Player struct {
	strategy map[StateIndex]ActionIndex
}

// Train given an environment, create a player
// that can play well in that environment
func Train(env Environment) Player {
	// TODO: Train the player, encode into it how to play
	strategy := make(map[StateIndex]ActionIndex)
	possibleActions := env.PossibleActions()
	for i := 0; i < env.PossibleStates(); i++ {
		strategy[StateIndex(i)] = ActionIndex(rand.Intn(possibleActions))
	}
	return Player{
		strategy: strategy,
	}
}

// Play a single episode with a player, returning it score
func (p *Player) Play(env Environment) int {

	state := env.InitialState()
	for {

		preferredAction := p.strategy[state]

		state = env.Take(preferredAction)
		if env.IsComplete(state) {
			break
		}
	}
	score := env.Score()
	env.Reset()
	return score
}
