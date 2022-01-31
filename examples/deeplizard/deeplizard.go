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

func (k *DeepLizardEvironment) PossibleActions(currentState coach.State) []coach.Action {
	actions := make([]coach.Action, 0)
	state, ok := currentState.(Tile)
	if !ok {
		panic("State is not a tile")
	}
	x_coordinate := state % 3
	y_coordinate := state / 3

	if x_coordinate > 0 {
		actions = append(actions, Left)
	}
	if x_coordinate < 2 {
		actions = append(actions, Right)
	}
	if y_coordinate > 0 {
		actions = append(actions, Up)
	}
	if y_coordinate < 2 {
		actions = append(actions, Down)
	}
	return actions
}

func (k *DeepLizardEvironment) PossibleStates() []coach.State {
	states := make([]coach.State, tials)
	for i := 0; i < tials; i++ {
		states[i] = Tile(i)
	}
	return states
}

func (k *DeepLizardEvironment) Evaluate(currentState coach.State, action coach.Action) (coach.State, coach.Reward) {
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
	switch direction {
	case Up:
		y_coordinate--
		if y_coordinate < 0 {
			panic("Moved up when on top row")
		}
	case Down:
		y_coordinate++
		if y_coordinate > 2 {
			panic("Moved down when on bottom row")
		}
	case Left:
		x_coordinate--
		if x_coordinate < 0 {
			panic("Moved left when on leftmost column")
		}
	case Right:
		x_coordinate++
		if x_coordinate > 2 {
			panic("Moved right when on rightmost column")
		}
	}
	newState := Tile(y_coordinate*3 + x_coordinate)
	//.Printf("Action %v moves from %v to %v\n", action, currentState, newState)
	tileType := k.board[newState]
	return coach.State(newState), coach.Reward(k.rewards[tileType].reward)

}

func (k *DeepLizardEvironment) IsComplete(currentState coach.State) bool {
	state, ok := currentState.(Tile)
	if !ok {
		panic("State is not a tile")
	}
	tileType := k.board[state]
	return k.rewards[tileType].terminal
}

// func (k *KonamiCodeEnvironment) Take(a coach.ActionIndex) {

// 	k.pointer += 1
// 	k.steps += 1
// 	if k.pointer == codeLength {
// 		k.pointer = 0
// 	}
// 	k.keys[k.pointer] = k.allKeys[a]

// }

// // How much of the code is currently done
// func (k *KonamiCodeEnvironment) currentRun() int {
// 	correctCode := []Key{
// 		UpKey,
// 		DownKey,
// 		UpKey,
// 		DownKey,
// 		LeftKey,
// 		RightKey,
// 		LeftKey,
// 		RightKey,
// 		AKey,
// 		BKey,
// 	}
// 	length := len(k.keys)
// 	i := 0
// 	for ; i < len(correctCode); i++ {
// 		index := length - i - 1
// 		if index < 0 {
// 			break
// 		}
// 		if k.keys[index] != correctCode[i] {
// 			break
// 		}

// 	}
// 	fmt.Printf("%v has run of %d\n", k.keys, i)
// 	return i

// }

// func (k *KonamiCodeEnvironment) IsComplete() bool {
// 	return k.currentRun() == 10
// }

// func (k *KonamiCodeEnvironment) Score() int {
// 	return k.steps
// }
