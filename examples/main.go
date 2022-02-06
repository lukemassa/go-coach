package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/lukemassa/go-coach/examples/deeplizard"
	"github.com/lukemassa/go-coach/examples/frozenlake"
	"github.com/lukemassa/go-coach/examples/konamicode"
	"github.com/lukemassa/go-coach/examples/wordle"
	"github.com/lukemassa/go-coach/pkg/coach"
)

func getEnv(env string) coach.Environment {
	if env == "konami" {
		return konamicode.New()
	}
	if env == "frozen" {
		return frozenlake.New()
	}
	if env == "deeplizard" {
		return deeplizard.New()
	}
	if env == "wordle" {
		return wordle.New()
	}
	panic(fmt.Sprintf("unexpected environment %s", env))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var opts struct {
		// Example of a required flag
		Environment string `short:"e" long:"environment" description:"Which environment" required:"true"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}
	env := getEnv(opts.Environment)

	player := coach.Train(env, 100_000)

	trials := 10
	fmt.Println("Done training, starting trials")
	start := time.Now()
	score := player.Evaluate(env, trials)
	duration := time.Since(start)

	fmt.Printf("Ran %d trials in %v, got average score of %f\n", trials, duration, score)
}
