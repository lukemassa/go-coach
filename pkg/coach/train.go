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
	for action := range q[state] {

		newState, reward := env.Evaluate(state, action)
		// Passing epsilon as 0 to make sure we pick the value itself
		optimalFutureAction := q.Choose(newState, 0)
		optimalFutureValue := float64(q[newState][optimalFutureAction])

		q[state][action] = QValue(1-learningRate)*q[state][action] + QValue(learningRate)*(QValue(reward)+QValue(discountFactor*optimalFutureValue))
	}
}

// Train given an environment, create a player
// that can play well in that environment
func Train(env Environment, episodes int) Player {

	qtable := NewQTable()

	// Fill in q state for initial state
	state := env.InitialState()
	qtable[state] = initialQRow(env.PossibleActions(state))

	learningRate := .9
	discountFactor := .9
	epsilon := .99
	decay := .0000000000001
	for i := 0; i < episodes; i++ {
		// Decay the learning
		learningRate = learningRate / (1 + decay*float64(i))
		//fmt.Printf("Learning rate is now %f\n", learningRate)
		state := env.InitialState()
		for {
			qtable.Update(state, env, discountFactor, learningRate)
			preferredAction := qtable.Choose(state, epsilon)

			state, _ = env.Evaluate(state, preferredAction)

			// Learn about new states, initialize the q state
			if _, ok := qtable[state]; !ok {
				qtable[state] = initialQRow(env.PossibleActions(state))
				//do something here
			}
			fmt.Printf("Chose action %v for state %v\n", preferredAction, state)
			if env.IsComplete(state) {
				break
			}
		}

	}
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
	maxSteps := env.MaxSteps()
	steps := 0
	for {
		steps += 1
		if steps > maxSteps {
			break
		}
		preferredAction := p.strategy.Choose(state, .1)
		//fmt.Printf("Preferred action for state %v is %v\n", state, preferredAction)
		newState, incrementalReward := env.Evaluate(state, preferredAction)
		score += incrementalReward
		state = newState

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
