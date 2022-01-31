package coach

import (
	"fmt"
	"math/rand"
	"strings"
)

type Distribution map[Action]float64

type Player struct {
	strategy QTable
}

type QValue Reward

type QTable map[State]map[Action]QValue

func NewQTable() QTable {
	return QTable(make(map[State]map[Action]QValue))
}

func (d QTable) Choose(state State, epsilon float64) Action {
	// Map from Actions -> QValues
	qrow := d[state]
	if rand.Float64() > epsilon {
		// Choose element w highest Q Value
		var maxAction Action
		for action, qvalue := range qrow {

			if maxAction == nil || qrow[maxAction] < qvalue {
				maxAction = action
			}
		}
		return maxAction
	}

	options := make([]Action, 0)
	// Loop through samples
	for action, _ := range qrow {
		options = append(options, action)
	}
	randIndex := rand.Intn(len(options))
	return options[randIndex]
}

func (q QTable) String() string {
	sb := strings.Builder{}
	for state, qrow := range q {
		sb.WriteString(fmt.Sprintf("State %v : ", state))
		for action, qvalue := range qrow {
			sb.WriteString(fmt.Sprintf("(%v:%v)", action, qvalue))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func initialQRow(actions []Action) map[Action]QValue {
	ret := make(map[Action]QValue)
	for _, action := range actions {
		ret[action] = 0
	}
	return ret
}

func (q QTable) Update(state State, env Environment, learningRate, discountFactor float64) {
	for action, _ := range q[state] {
		// Passing epsilon as 1 to make sure we pick the value itself
		optimalFutureAction := q.Choose(state, 1)
		optimalFutureValue := float64(q[state][optimalFutureAction])
		_, reward := env.Evaluate(state, action)
		q[state][action] = QValue(1-learningRate)*q[state][action] + QValue(learningRate)*(QValue(reward)+QValue(discountFactor*optimalFutureValue))
	}
}

// Train given an environment, create a player
// that can play well in that environment
func Train(env Environment, episodes int) Player {

	qtable := NewQTable()
	possibleStates := env.PossibleStates()
	for _, state := range possibleStates {
		// Every strategy starts out with an even distributions over
		// its possible actions
		qtable[state] = initialQRow(env.PossibleActions(state))
	}
	fmt.Printf("%s\n", qtable)
	learningRate := .7
	discountFactor := .9
	epsilon := .99
	for i := 0; i < episodes; i++ {
		state := env.InitialState()
		//fmt.Printf("Working on episode %d\n", i)
		for {
			qtable.Update(state, env, discountFactor, learningRate)
			preferredAction := qtable.Choose(state, epsilon)

			state, _ = env.Evaluate(state, preferredAction)
			//fmt.Printf("Chose action %v for state %v\n", preferredAction, state)
			if env.IsComplete(state) {
				break
			}
		}

	}
	// TODO: Train
	fmt.Printf("%s\n", qtable)

	return Player{
		strategy: qtable,
	}
}

// Play a single episode with a player, returning it score
func (p *Player) Play(env Environment) Reward {

	//env.Reset()
	var score Reward
	state := env.InitialState()
	//fmt.Printf("Strategy %v\n", p.strategy)
	for {

		preferredAction := p.strategy.Choose(state, 0)
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
