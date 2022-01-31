package coach

type Action interface{}

type State interface{}

type Reward int

type Environment interface {
	// What does the environment look like
	//Reset()
	InitialState() State
	PossibleActions(State) []Action
	PossibleStates() []State

	// Training
	Evaluate(State, Action) (State, Reward)
	IsComplete(State) bool
}
