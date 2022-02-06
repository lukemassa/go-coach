package coach

type Action interface{}

type State interface{}

type Reward float64

type Score float64

// Environment where the agent will learn
// Note that the environment itself is stateless, it's up to
// any code interacting with the environment to keep track of
// which state the agent is in.
type Environment interface {
	// What does the environment look like
	MaxSteps() int
	InitialState() State
	PossibleActions() []Action

	// Some environments may need to change *between*
	// episodes. For example, wordle has a new "goal"
	// for each episode, but the rest of the environment
	// remains the same. Environments like chess or even
	// random ones like tetris don't need updates
	Update()

	// Training
	// Given a state and an action, which state will this bring
	// us to, how much reward do we get there, and is it a terminal state
	Evaluate(State, Action) (State, Reward, bool)

	// Validation
	// How "well" did a player do that produced this particular
	// stream of states. Assume that the first state is the iniitial
	// state and the last state is a terminal one
	// This may differ from simply the sum of all the rewards determined
	// by Evaluate(), or it may not
	Score([]State) Score
}
