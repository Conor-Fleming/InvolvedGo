package letter

import (
	"fmt"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(input []string) FreqMap {
	c := make(chan map[rune]int)
	result := FreqMap{}
	for _, v := range input {
		v := v
		go func() {
			c <- Frequency(v)
		}()
	}

	x, y, z := <-c, <-c, <-c

	fmt.Println(x, y, z)

	fmt.Println("Concurrent terminated")
	return result
}
