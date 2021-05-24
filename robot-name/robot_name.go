//Package robotname package gives us utilities for naming robots
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

//Robot type creates a robot object with a property of name
type Robot struct {
	name string
}

var namesUsed = map[string]bool{}

//Name function checks to see if a Robot has a name
func (r *Robot) Name() (string, error) {
	//create name and check if name is in use in map namesUsed
	if r.name != "" {
		return r.name, nil
	}
	if len(namesUsed) == 26*26*10*10*10 {
		return "", fmt.Errorf("There are no more possible unique combinations of robot names :(")
	}
	r.name = newName()
	for namesUsed[r.name] {
		r.name = newName()
	}
	namesUsed[r.name] = true
	return r.name, nil
}

//Reset function helps with the clearing of the robots name
func (r *Robot) Reset() {
	r.name = ""
}

//NewName helps with generation of robot names
func newName() string {
	r1 := seededRand.Intn(26) + 'A'
	r2 := seededRand.Intn(26) + 'A'
	num := seededRand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
