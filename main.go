package main

import (
	"fmt"

	"github.com/wowTerrific/coinflip/flip"
)

func main() {
	newFlip := flip.Flip()
	if newFlip == "Heads" {
		fmt.Println(flip.HEADS)
	} else {
		fmt.Println(flip.TAILS)
	}

}
