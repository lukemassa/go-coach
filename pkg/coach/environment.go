package coach

type ActionIndex int

type StateIndex int

type Reward float64

type Environment interface {
	Reset()
	// What does the environment look like
	InitialState() StateIndex
	PossibleActions() int
	PossibleStates() int

	// Training
	Evaluate(ActionIndex) Reward
	Take(ActionIndex)
	CurrentState() StateIndex

	// Playing
	IsComplete() bool
	Score() int
}
