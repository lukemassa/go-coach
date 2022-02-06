/*

Simple game whose only goal is to type the Konami Code (https://en.wikipedia.org/wiki/Konami_Code)

The game runs forever until the code is entered in the right order, and a player does "well" by typing it sooner

*/
package konamicode

import (
	"fmt"

	"github.com/lukemassa/go-coach/pkg/coach"
)

type Key int8

func (k Key) String() string {
	if k == UpKey {
		return "U"
	}
	if k == DownKey {
		return "D"
	}
	if k == LeftKey {
		return "L"
	}
	if k == RightKey {
		return "R"
	}
	if k == AKey {
		return "A"
	}
	if k == BKey {
		return "B"
	}
	panic("Unreachable")
}

const codeLength = 10

const (
	UpKey Key = iota
	DownKey
	LeftKey
	RightKey
	AKey
	BKey
)

type GameState [codeLength]Key

type KonamiCodeEnvironment struct {
	allKeys []Key
}

func New() *KonamiCodeEnvironment {
	k := KonamiCodeEnvironment{
		allKeys: []Key{
			UpKey,
			DownKey,
			LeftKey,
			RightKey,
			AKey,
			BKey,
		},
	}
	ret := &k
	return ret

}

func (k *KonamiCodeEnvironment) MaxSteps() int {
	return 100
}

func (k *KonamiCodeEnvironment) PossibleActions() []coach.Action {
	actions := make([]coach.Action, len(k.allKeys))
	for i, key := range k.allKeys {
		actions[i] = key
	}
	return actions
}

func (k *KonamiCodeEnvironment) InitialState() coach.State {
	return GameState([codeLength]Key{})
}

func (k *KonamiCodeEnvironment) Evaluate(currentState coach.State, action coach.Action) (coach.State, coach.Reward, bool) {
	currentGame, ok := currentState.(GameState)
	if !ok {
		panic("State is not a game")
	}
	newKey, ok := action.(Key)
	if !ok {
		panic(fmt.Sprintf("Action is not a key, it is '%v'", action))
	}
	// Bump one off the end
	newGame := GameState([codeLength]Key{
		newKey,
		currentGame[0],
		currentGame[1],
		currentGame[2],
		currentGame[3],
		currentGame[4],
		currentGame[5],
		currentGame[6],
		currentGame[7],
		currentGame[8],
	})
	reward := currentRun(newGame)
	if reward > 9 {
		fmt.Printf("Got reward %v for %v\n", reward, newGame)
	}

	return newGame, coach.Reward(reward), reward == codeLength
}

func currentRun(game GameState) int {
	correctCode := [codeLength]Key{
		UpKey,
		DownKey,
		UpKey,
		DownKey,
		LeftKey,
		RightKey,
		LeftKey,
		RightKey,
		AKey,
		BKey,
	}
	//fmt.Printf("Looking for matches in %v\n", game)
	for i := codeLength - 1; i >= 0; i-- {
		matches := true

		for j := 0; j <= i; j++ {
			if correctCode[i-j] != game[j] {
				//fmt.Printf("Sorry, %v is not %v\n", correctCode[j], game[i])
				matches = false
				break
			}
		}
		if matches {
			return i + 1

		}
	}
	return 0

}

func (k *KonamiCodeEnvironment) Update() {
	// Nothing to do between runs
}

func (k *KonamiCodeEnvironment) Score(states []coach.State) coach.Score {
	// How well you did is merely how many buttons you had to press
	return coach.Score(len(states))
}
