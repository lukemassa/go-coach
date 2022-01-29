package coach

type ActionIndex int

type Reward float64

type Environment interface {
	Reset()
	PossibleActions() int
	Evaluate(ActionIndex) Reward
	Take(ActionIndex)
	IsComplete() bool
	Score() int
}
