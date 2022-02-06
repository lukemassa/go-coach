/*

https://www.powerlanguage.co.uk/wordle/

*/
package wordle

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/lukemassa/go-coach/pkg/coach"
)

const wordLength = 5
const numLetters = 3

type word [wordLength]rune

type WordleEnvironment struct {
	letters    [numLetters]rune
	allWords   []word
	targetWord word
}

type gameState struct {
	previousWord   word
	correctLetters word
}

func (w word) String() string {
	var sb strings.Builder
	for i := 0; i < wordLength; i++ {
		sb.WriteRune(w[i])
	}
	return sb.String()
}

func (g gameState) String() string {
	return fmt.Sprintf("Correct letters: %s", g.correctLetters)
}

func getEnglishLetters() [numLetters]rune {
	return [numLetters]rune{
		'A',
		'B',
		'C',
	}
}

func getEnglishWords() []word {
	return []word{
		{'A', 'A', 'A', 'B', 'B'},
		{'A', 'A', 'C', 'A', 'A'},
		{'B', 'A', 'C', 'A', 'A'},
		{'C', 'A', 'A', 'A', 'B'},
		{'A', 'B', 'C', 'B', 'A'},
		{'A', 'A', 'C', 'A', 'B'},
		{'B', 'C', 'B', 'C', 'B'},
	}
}

func New() *WordleEnvironment {
	w := WordleEnvironment{
		letters:  getEnglishLetters(),
		allWords: getEnglishWords(),
	}
	ret := &w
	w.Update()
	return ret
}

func (w *WordleEnvironment) Update() {
	// Pick a new word!
	index := rand.Intn(len(w.allWords))
	w.targetWord = w.allWords[index]
}

func (w *WordleEnvironment) InitialState() coach.State {
	return gameState{
		correctLetters: word{' ', ' ', ' ', ' ', ' '},
	}
}

func (w *WordleEnvironment) PossibleActions() []coach.Action {
	actions := make([]coach.Action, len(w.allWords))
	for i := 0; i < len(w.allWords); i++ {
		actions[i] = w.allWords[i]
	}
	return actions
}

func (w *WordleEnvironment) MaxSteps() int {
	return 10
}

func (w *WordleEnvironment) Evaluate(currentState coach.State, action coach.Action) (coach.State, coach.Reward, bool) {
	gameState, ok := currentState.(gameState)
	if !ok {
		panic("State is not a gameState")
	}
	newWord, ok := action.(word)
	if gameState.previousWord == newWord {
		return coach.State(gameState), -10, false
	}
	gameState.previousWord = newWord
	if !ok {
		panic("Action is not a direction")
	}
	correctLetters := 0
	for i := 0; i < numLetters; i++ {
		if newWord[i] == w.targetWord[i] {
			gameState.correctLetters[i] = newWord[i]
			correctLetters += 1
		}
	}
	return coach.State(gameState), coach.Reward(correctLetters), correctLetters == numLetters

}

func (w *WordleEnvironment) Score(states []coach.State) coach.Score {
	// How well you did is just how many words you had to guess
	return coach.Score(len(states))
}

func (w *WordleEnvironment) Show(states []coach.State, interactive bool) {
	for i := 0; i < len(states); i++ {
		gameState, ok := states[i].(gameState)
		if !ok {
			panic("State is not a gameState")
		}
		fmt.Println(gameState.correctLetters)
	}
	fmt.Printf("Took %d attempts to get to correct word of %v\n", len(states), w.targetWord)
}
