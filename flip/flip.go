package flip

import (
	"math/rand"
)

func Flip() string {
	answers := []string{
		"Heads",
		"Tails",
	}

	return answers[rand.Intn(len(answers))]
}
