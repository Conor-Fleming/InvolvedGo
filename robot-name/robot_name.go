//Package robotname package gives us utilities for naming robots
package robotname

import (
	"fmt"
	"math/rand"
)

//Robot type creates a robot object with a property of name
type Robot struct {
	name string
}

var namesUsed = map[string]bool{}

//Name function checks to see if a Robot has a name
func (r *Robot) Name() (string, error) {
	//create letters/numbers for name and check if name is in use in map namesUsed
	if r.name == "" {
		r.name = newName()
		for namesUsed[r.name] {
			r.name = newName()
		}
	}
	return r.name, nil
}

//Reset function helps with the clearing of the robots name and assignment of a new name with a call to Name()
func (r *Robot) Reset() (string, error) {
	r.name = ""
	return r.Name()
}

func newName() string {
	r1 := rand.Intn(26) + 'A'
	r2 := rand.Intn(26) + 'A'
	num := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
