package main

import (
	"fmt"
	"time"

	"github.com/wowTerrific/coinflip/flip"
)

func main() {
	newFlip := flip.Flip()

	chars := []rune{'_', '/', '-', '\\', '|', '/', '-', '\\', '-', '_'}

	for i := 0; i < len(chars); i++ {
		fmt.Printf("\rFLIP!\t%c\t", chars[i%len(chars)])
		time.Sleep(175 * time.Millisecond)
	}
	fmt.Printf("\rFlip Result: \n")

	if newFlip == "Heads" {
		fmt.Printf("\r%s%s%s", flip.BLUE, flip.HEADS_IMG, flip.RESET)
	} else {
		fmt.Printf("\r%s%s%s", flip.GREEN, flip.TAILS_IMG, flip.RESET)
	}

}
