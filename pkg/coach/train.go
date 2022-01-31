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

	for _, action := range actions {
		ret[action] = probability
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
	fmt.Println("Strategies before training:")
	for _, state := range possibleStates {
		// Every strategy starts out with an even distributions over
		// its possible actions
		strategy[state] = initialDistribution(env.PossibleActions(state))
		fmt.Printf("  %v is %v\n", state, strategy[state])
	}
	// TODO: Train :

	fmt.Println("Strategies after training:")
	for _, state := range possibleStates {
		// Every strategy starts out with an even distributions over
		// its possible actions
		fmt.Printf("  %v is %v\n", state, strategy[state])
	}

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
		newState, incrementalReward := env.Evaluate(state, preferredAction)
		score += incrementalReward
		state = newState

		//fmt.Printf("Preferred action is %v\n", preferredAction)

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
