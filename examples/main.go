package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lukemassa/go-coach/examples/deeplizard"
	"github.com/lukemassa/go-coach/pkg/coach"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	env := deeplizard.New()

	player := coach.Train(env, 0)

	trials := 10000

	start := time.Now()
	score := player.Evaluate(env, trials)
	duration := time.Since(start)

	fmt.Printf("Evaluated untrained player: Ran %d trials in %v, got average score of %f\n", trials, duration, score)

	player = coach.Train(env, 100)

	start = time.Now()
	score = player.Evaluate(env, trials)
	duration = time.Since(start)

	fmt.Printf("Evaluated trained player: Ran %d trials in %v, got average score of %f\n", trials, duration, score)
}
