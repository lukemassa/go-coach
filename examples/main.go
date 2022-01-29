package main

import (
	"fmt"

	"github.com/lukemassa/go-coach/examples/konamicode"
	"github.com/lukemassa/go-coach/pkg/coach"
)

func main() {
	env := konamicode.New()

	player := coach.Train(env)

	score := player.Play(env)

	fmt.Printf("Got score of %d\n", score)
}
