package coach

type Action interface{}

type State interface{}

type Reward float64

// Environment where the agent will learn
// Note that the environment itself is stateless, it's up to
// any code interacting with the environment to keep track of
// which state the agent is in.
type Environment interface {
	// What does the environment look like
	MaxSteps() int
	InitialState() State
	PossibleActions(State) []Action

	// Training
	Evaluate(State, Action) (State, Reward)
	IsComplete(State) bool
}
