package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbset = "1234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type Robot struct {
	name string
}

var namesUsed = map[string]bool{}

func (r *Robot) Name() (string, error) {
	//create letters/numbers for name and check if name is in use in map namesUsed
	if r.name == "" {
		name := createLetters() + createNumbers()
		if robotNameCheck(name) {
			return "", fmt.Errorf("This robot name is already being used")
		}
		namesUsed[name] = true
		r.name = name
	}
	return r.name, nil
}

func (r *Robot) Reset() (string, error) {
	r.name = ""
	return r.Name()
}

func createLetters() string {
	letters := make([]byte, 2)
	for i := range letters {
		letters[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(letters)
}
func createNumbers() string {
	numbers := make([]byte, 3)
	for i := range numbers {
		numbers[i] = numbset[seededRand.Intn(len(numbset))]
	}
	return string(numbers)
}

func robotNameCheck(name string) bool {
	if _, present := namesUsed[name]; present {
		return true
	}
	return false
}
