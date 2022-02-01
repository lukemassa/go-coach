package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lukemassa/go-coach/examples/frozenlake"
	"github.com/lukemassa/go-coach/pkg/coach"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	env := frozenlake.New()

	player := coach.Train(env, 100000)

	trials := 1

	start := time.Now()
	score := player.Evaluate(env, trials)
	duration := time.Since(start)

	fmt.Printf("Ran %d trials in %v, got average score of %f\n", trials, duration, score)
}
