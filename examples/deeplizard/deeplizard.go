/*

A simple game where a lizard walks around a board and tries to find crickets to eat: https://www.youtube.com/watch?v=qhRNvCVVJaA

*/
package deeplizard

import (
	"fmt"
	"strconv"

	"github.com/lukemassa/go-coach/pkg/coach"
)

type TileType int
type Tile int
type Direction int

const dimension = 3
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
	x_coordinate := state % dimension
	y_coordinate := state / dimension
	// Don't move if on an edge
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
	reward := k.rewards[k.board[newState]]
	return coach.State(newState), coach.Reward(reward.reward), reward.terminal

}

func (k *DeepLizardEvironment) Update() {
	// Nothing to do between runs
}

func (k *DeepLizardEvironment) Score(states []coach.State) coach.Score {
	// How well you did is identical to the sum of rewards gathered during the game
	var score coach.Score
	for i := 0; i < len(states); i++ {
		tile, ok := states[i].(Tile)
		if !ok {
			panic("State is not a tile")
		}
		score += coach.Score(k.rewards[k.board[tile]].reward)
	}
	return score
}

func showStrBoard(strBoard [tials]string) {
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {

			fmt.Print(strBoard[i*dimension+j])
		}
		fmt.Println()
	}
}

func (k *DeepLizardEvironment) Show(states []coach.State, interactive bool) {
	strBoard := [tials]string{}
	for i := 0; i < tials; i++ {
		var tialStr string
		switch k.board[i] {
		case Empty:
			tialStr = " "
		case OneCricket:
			tialStr = "c"
		case FiveCricket:
			tialStr = "C"
		case Bird:
			tialStr = "B"
		}
		strBoard[i] = tialStr
	}
	for i := 0; i < len(states); i++ {

		tialStr := strconv.Itoa(i)

		if i >= 10 {
			tialStr = "*"
		}
		tile, ok := states[i].(Tile)
		if !ok {
			panic("State is not a tile")
		}
		strBoard[tile] = tialStr
		if interactive {
			fmt.Print("\033[H\033[2J")
			showStrBoard(strBoard)
			fmt.Scanln()
		}
	}
	if interactive {
		fmt.Printf("Finished! Score: %f", k.Score(states))
		fmt.Scanln()
	} else {
		showStrBoard(strBoard)
	}
}
