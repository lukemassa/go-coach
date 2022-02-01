package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lukemassa/go-coach/examples/konamicode"
	"github.com/lukemassa/go-coach/pkg/coach"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	env := konamicode.New()

	player := coach.Train(env, 10000)

	trials := 1

	start := time.Now()
	score := player.Evaluate(env, trials)
	duration := time.Since(start)

	fmt.Printf("Ran %d trials in %v, got average score of %f\n", trials, duration, score)
}
