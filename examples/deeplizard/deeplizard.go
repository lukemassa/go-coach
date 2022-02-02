/*

A simple game where a lizard walks around a board and tries to find crickets to eat: https://www.youtube.com/watch?v=qhRNvCVVJaA

*/
package deeplizard

import (
	"github.com/lukemassa/go-coach/pkg/coach"
)

type TileType int
type Tile int
type Direction int

const tials = 9

/*
c = one cricket, +1
C = 5 crickets, +10 (and game over)
B = Bird, -10 (and game over)
s = start
Empty squares are -1
    | c |   |   |
	|   | B |   |
    | s |   | C |
*/

const (
	Empty TileType = iota
	OneCricket
	FiveCricket
	Bird
)
const (
	Up Direction = iota
	Down
	Left
	Right
)

func (d Direction) String() string {
	if d == Up {
		return "Up"
	}
	if d == Down {
		return "Down"
	}
	if d == Left {
		return "Left"
	}
	if d == Right {
		return "Right"
	}
	panic("wtf")
}

type Reward struct {
	reward   int
	terminal bool
}

type DeepLizardEvironment struct {
	board   [tials]TileType
	rewards map[TileType]Reward
}

func New() *DeepLizardEvironment {
	k := DeepLizardEvironment{
		board: [tials]TileType{
			OneCricket,
			Empty,
			Empty,
			Empty,
			Bird,
			Empty,
			Empty,
			Empty,
			FiveCricket,
		},
		rewards: map[TileType]Reward{
			Empty: {
				reward:   -1,
				terminal: false,
			},
			OneCricket: {
				reward:   1,
				terminal: false,
			},
			FiveCricket: {
				reward:   10,
				terminal: true,
			},
			Bird: {
				reward:   -10,
				terminal: true,
			},
		},
	}
	ret := &k
	//ret.Reset()
	return ret

}

func (k *DeepLizardEvironment) InitialState() coach.State {
	return Tile(6) // Put the lizard in the bottom corner
}

func (k *DeepLizardEvironment) PossibleActions() []coach.Action {
	return []coach.Action{
		Left,
		Right,
		Up,
		Down,
	}
}

func (k *DeepLizardEvironment) MaxSteps() int {
	return 100
}

func (k *DeepLizardEvironment) PossibleStates() []coach.State {
	states := make([]coach.State, tials)
	for i := 0; i < tials; i++ {
		states[i] = Tile(i)
	}
	return states
}

func (k *DeepLizardEvironment) Evaluate(currentState coach.State, action coach.Action) (coach.State, coach.Reward, bool) {
	state, ok := currentState.(Tile)
	if !ok {
		panic("State is not a tile")
	}
	direction, ok := action.(Direction)
	if !ok {
		panic("Action is not a direction")
	}
	x_coordinate := state % 3
	y_coordinate := state / 3
	// Don't move if on an edge
	switch direction {
	case Up:
		y_coordinate--
		if y_coordinate < 0 {
			y_coordinate = 0
		}
	case Down:
		y_coordinate++
		if y_coordinate > 2 {
			y_coordinate = 2
		}
	case Left:
		x_coordinate--
		if x_coordinate < 0 {
			x_coordinate = 0
		}
	case Right:
		x_coordinate++
		if x_coordinate > 2 {
			x_coordinate = 2
		}
	}
	newState := Tile(y_coordinate*3 + x_coordinate)
	//.Printf("Action %v moves from %v to %v\n", action, currentState, newState)
	reward := k.rewards[k.board[newState]]
	return coach.State(newState), coach.Reward(reward.reward), reward.terminal

}
