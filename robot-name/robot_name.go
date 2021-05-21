package robotname

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type Robot struct {
	name string
}

//namesUsed := make(map[string]int)

func (r *Robot) Name() (string, error) {
	//create letters/numbers for name and check if name is in use in map namesUsed

	name := createName()
	return name, nil
}

func (r *Robot) Reset() {

}
func createLetters() string {
	letters := make([]byte, 2)
	for i := range letters {
		letters[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(letters)
}
func createNumbers

func robotNameCheck() bool {
	return true
}
