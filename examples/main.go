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

	player := coach.Train(env)

	trials := 10000

	start := time.Now()
	score := player.Evaluate(env, 10000)
	duration := time.Now().Sub(start)

	fmt.Printf("Ran %d trials in %v, got average score of %d\n", trials, duration, score)
}
