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
	//AKey
	//BKey
)

type GameState struct {
	keys    [codeLength]Key
	pointer int
}

type KonamiCodeEnvironment struct {
	state   GameState
	allKeys []Key
	turn    int
}

func initialState() GameState {
	return GameState{
		keys:    [codeLength]Key{},
		pointer: 0,
	}
}

func New() *KonamiCodeEnvironment {
	k := KonamiCodeEnvironment{
		allKeys: []Key{
			UpKey,
			DownKey,
			LeftKey,
			RightKey,
			//AKey,
			//BKey,
		},
		turn:  0,
		state: initialState(),
	}
	ret := &k
	//ret.Reset()
	return ret

}

// func (k *KonamiCodeEnvironment) Reset() {
// 	keys := make([]Key, codeLength)
// 	k.keys = keys
// 	k.pointer = 0
// 	k.steps = 0
// }

func (k *KonamiCodeEnvironment) PossibleActions() []coach.Action {
	actions := make([]coach.Action, len(k.allKeys))
	for i, key := range k.allKeys {
		actions[i] = key
	}
	return actions
}

func (k *KonamiCodeEnvironment) PossibleStates() []coach.State {
	states := make([]coach.State, 0)
	for a1 := 0; a1 < len(k.allKeys); a1++ {
		fmt.Printf("Working on %d\n", a1)
		for a2 := 0; a2 < len(k.allKeys); a2++ {
			for a3 := 0; a3 < len(k.allKeys); a3++ {
				for a4 := 0; a4 < len(k.allKeys); a4++ {
					for a5 := 0; a5 < len(k.allKeys); a5++ {
						for a6 := 0; a6 < len(k.allKeys); a6++ {
							for a7 := 0; a7 < len(k.allKeys); a7++ {
								for a8 := 0; a8 < len(k.allKeys); a8++ {
									for a9 := 0; a9 < len(k.allKeys); a9++ {
										for a10 := 0; a10 < len(k.allKeys); a10++ {
											state := GameState{
												keys: [10]Key{
													k.allKeys[a1],
													k.allKeys[a2],
													k.allKeys[a3],
													k.allKeys[a4],
													k.allKeys[a5],
													k.allKeys[a6],
													k.allKeys[a7],
													k.allKeys[a8],
													k.allKeys[a9],
													k.allKeys[a10],
												},
												pointer: 0,
											}
											states = append(states, state)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return states
}

// func (k *KonamiCodeEnvironment) Evaluate(a coach.ActionIndex) coach.Reward {
// 	// TODO: Implement reward
// 	return 0
// }

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
