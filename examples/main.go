package main

import (
	"fmt"

	"github.com/lukemassa/go-coach/pkg/coach"
)

func main() {
	env := KonamiCodeEnvironment{}

	player := coach.Train(env)

	env.Reset()

	score := player.Play(env)

	fmt.Printf("Got score of %d (best possible score is 10)\n", score)
}
