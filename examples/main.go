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

	score := player.Evaluate(env, 10000)

	fmt.Printf("Got average score of %d\n", score)
}
