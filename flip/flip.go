package flip

import (
	"math/rand"
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
	answers := []string{
		"Heads",
		"Tails",
	}

	return answers[rand.Intn(len(answers))]
}
