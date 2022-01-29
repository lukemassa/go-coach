/*

Simple game whose only goal is to type the Konami Code (https://en.wikipedia.org/wiki/Konami_Code)

The game runs forever until the code is entered in the right order, and a player does "well" by typing it sooner

*/
package konamicode

import "github.com/lukemassa/go-coach/pkg/coach"

type Key int

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
	keys := make([]Key, 0)
	k.keys = keys
}

func (k *KonamiCodeEnvironment) Evaluate(a coach.ActionIndex) coach.Reward {
	return 0
}

func (k *KonamiCodeEnvironment) Take(index coach.ActionIndex) {
	k.keys = append(k.keys, k.allKeys[index])
}

func (k *KonamiCodeEnvironment) PossibleActions() int {
	return len(k.allKeys)
}

func (k *KonamiCodeEnvironment) Score() int {
	return len(k.keys)
}

func (k *KonamiCodeEnvironment) IsComplete() bool {
	length := len(k.keys)
	// For now just do up/down/up/down
	if len(k.keys) < 10 {
		return false
	}
	return k.keys[length-10] == UpKey &&
		k.keys[length-9] == DownKey &&
		k.keys[length-8] == UpKey &&
		k.keys[length-7] == DownKey &&
		k.keys[length-6] == LeftKey &&
		k.keys[length-5] == RightKey &&
		k.keys[length-4] == LeftKey &&
		k.keys[length-3] == RightKey &&
		k.keys[length-2] == BKey &&
		k.keys[length-1] == AKey
}
