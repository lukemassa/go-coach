package coach

import (
	"fmt"
	"math/rand"
)

type Distribution map[Action]float64

type Player struct {
	strategy map[State]Distribution
}

func initialDistribution(actions []Action) Distribution {
	ret := make(map[Action]float64)

	probability := 1.0 / float64(len(actions))

	hasRight := false
	for _, action := range actions {
		actionString := fmt.Sprintf("%v", action)
		if actionString == "0" {
			hasRight = true
		}
		ret[action] = probability
	}
	if hasRight {
		for _, action := range actions {
			actionString := fmt.Sprintf("%v", action)
			if actionString == "0" {
				ret[action] = 1
			} else {
				ret[action] = 0
			}
		}
	}
	return ret
}

// Sample returns an action based on the distribution
func (d Distribution) Sample() Action {
	val := rand.Float64()
	inc := 0.0

	// Loop through samples
	for action, probability := range d {

		inc += probability
		if inc > val {
			return action
		}

	}
	panic("UNREACHABLE CODE")
}

// Train given an environment, create a player
// that can play well in that environment
func Train(env Environment) Player {

	strategy := make(map[State]Distribution)
	possibleStates := env.PossibleStates()
	fmt.Printf("There are %d states\n", len(possibleStates))
	for _, state := range possibleStates {
		// Every strategy starts out with an even distributions over
		// its possible actions
		strategy[state] = initialDistribution(env.PossibleActions(state))
		fmt.Printf("Strategy for %v is %v\n", state, strategy[state])
	}
	// TODO: Train :)
	return Player{
		strategy: strategy,
	}
}

// Play a single episode with a player, returning it score
func (p *Player) Play(env Environment) Reward {

	//env.Reset()
	var score Reward
	state := env.InitialState()
	//fmt.Printf("Strategy %v\n", p.strategy)
	for {

		preferredAction := p.strategy[state].Sample()
		score += env.Evaluate(state, preferredAction)

		//fmt.Printf("Preferred action is %v\n", preferredAction)

		state = env.Take(state, preferredAction)
		if env.IsComplete(state) {
			break
		}
	}

	return score
}

func (p *Player) Evaluate(env Environment, episodes int) Reward {
	var total Reward
	for i := 0; i < episodes; i++ {
		total += p.Play(env)
	}
	return total / Reward(episodes)
}
