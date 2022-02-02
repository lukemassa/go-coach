/*

https://gym.openai.com/envs/FrozenLake-v0/

*/
package frozenlake

import (
	"github.com/lukemassa/go-coach/pkg/coach"
)

type TileType int
type Tile int
type Direction int

const tials = 16
const dimension = 4

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
	Frozen TileType = iota
	Hole
	Goal
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

type FrozenHoleEnvironment struct {
	board   [tials]TileType
	rewards map[TileType]Reward
}

func New() *FrozenHoleEnvironment {
	f := FrozenHoleEnvironment{
		board: [tials]TileType{
			Frozen,
			Frozen,
			Frozen,
			Frozen,
			Frozen,
			Hole,
			Frozen,
			Hole,
			Frozen,
			Frozen,
			Frozen,
			Hole,
			Hole,
			Frozen,
			Frozen,
			Goal,
		},
		rewards: map[TileType]Reward{
			Hole: {
				reward:   0,
				terminal: true,
			},
			Frozen: {
				reward:   0,
				terminal: false,
			},
			Goal: {
				reward:   1,
				terminal: true,
			},
		},
	}
	ret := &f
	//ret.Reset()
	return ret

}

func (f *FrozenHoleEnvironment) InitialState() coach.State {
	return Tile(0) // Put the lizard in the bottom corner
}

func (f *FrozenHoleEnvironment) PossibleActions() []coach.Action {
	return []coach.Action{
		Left,
		Right,
		Up,
		Down,
	}
}

func (f *FrozenHoleEnvironment) MaxSteps() int {
	return 100000
}

func (f *FrozenHoleEnvironment) Evaluate(currentState coach.State, action coach.Action) (coach.State, coach.Reward, bool) {
	state, ok := currentState.(Tile)
	if !ok {
		panic("State is not a tile")
	}
	direction, ok := action.(Direction)
	if !ok {
		panic("Action is not a direction")
	}
	x_coordinate := state % dimension
	y_coordinate := state / dimension
	switch direction {
	case Up:
		y_coordinate--
		if y_coordinate < 0 {
			y_coordinate = 0
		}
	case Down:
		y_coordinate++
		if y_coordinate > dimension-1 {
			y_coordinate = dimension - 1
		}
	case Left:
		x_coordinate--
		if x_coordinate < 0 {
			x_coordinate = 0
		}
	case Right:
		x_coordinate++
		if x_coordinate > dimension-1 {
			x_coordinate = dimension - 1
		}
	}
	newState := Tile(y_coordinate*dimension + x_coordinate)
	//.Printf("Action %v moves from %v to %v\n", action, currentState, newState)
	reward := f.rewards[f.board[newState]]
	return coach.State(newState), coach.Reward(reward.reward), reward.terminal

}
