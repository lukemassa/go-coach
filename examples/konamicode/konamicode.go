/*

Simple game whose only goal is to type the Konami Code (https://en.wikipedia.org/wiki/Konami_Code)

The game runs forever until the code is entered in the right order, and a player does "well" by typing it sooner

*/
package konamicode

import (
	"fmt"

	"github.com/lukemassa/go-coach/pkg/coach"
)

type Key int

const codeLength = 10

const (
	UpKey Key = iota
	DownKey
	LeftKey
	RightKey
	AKey
	BKey
)

type KonamiCodeEnvironment struct {
	keys    []Key
	allKeys []Key
	pointer int
	steps   int
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
	ret.Reset()
	return ret

}
func (k *KonamiCodeEnvironment) Reset() {
	keys := make([]Key, codeLength)
	k.keys = keys
	k.pointer = 0
	k.steps = 0
}

func (k KonamiCodeEnvironment) InitialState() coach.StateIndex {
	return 0
}

func (k *KonamiCodeEnvironment) PossibleActions() int {
	return len(k.allKeys)
}

func (k *KonamiCodeEnvironment) Evaluate(a coach.ActionIndex) coach.Reward {
	// TODO: Implement reward
	return 0
}

func (k *KonamiCodeEnvironment) Take(a coach.ActionIndex) {

	k.pointer += 1
	k.steps += 1
	if k.pointer == codeLength {
		k.pointer = 0
	}
	k.keys[k.pointer] = k.allKeys[a]

}

func (k *KonamiCodeEnvironment) PossibleStates() int {
	// Anywhere along the code is a different "state"
	return codeLength
}

// How much of the code is currently done
func (k *KonamiCodeEnvironment) currentRun() int {
	correctCode := []Key{
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
	length := len(k.keys)
	i := 0
	for ; i < len(correctCode); i++ {
		index := length - i - 1
		if index < 0 {
			break
		}
		if k.keys[index] != correctCode[i] {
			break
		}

	}
	fmt.Printf("%v has run of %d\n", k.keys, i)
	return i

}

func (k *KonamiCodeEnvironment) IsComplete() bool {
	return k.currentRun() == 10
}

func (k *KonamiCodeEnvironment) Score() int {
	return k.steps
}
