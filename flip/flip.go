package flip

import (
	"math/rand/v2"
)

const (
	HEADS_IMG string = ` 
	████████████        
    ██████      ████████    
  ████              ████    
  ██     ██     ██    ████  
████     ██     ██    ████  
██       ██     ██      ████
██       ██     ██      ████
██       █████████      ████
██       ██     ██      ████
████     ██     ██    ████  
  ██     ██     ██   ████   
  ████              ████    
    ██████      ████████    
        ████████████   
          HEADS`

	TAILS_IMG string = ` 
        ████████████        
    ██████      ████████    
  ████              ████    
  ██    ██████████    ████  
████        ██        ████  
██          ██          ████
██          ██          ████
██          ██          ████
██          ██          ████
████        ██        ████  
  ██        ██       ████   
  ████              ████    
    ██████      ████████    
        ████████████
          TAILS`

	BLUE  string = "\033[34m"
	GREEN string = "\033[32m"
	RESET string = "\033[0m"
)

func Flip() string {
	// 4 options are better than 2
	// computers aren't actually random ¯\_(ツ)_/¯
	options := []string{
		"Heads",
		"Tails",
		"Heads",
		"Tails",
	}

	r := rand.IntN(len(options))

	return options[r]
}
