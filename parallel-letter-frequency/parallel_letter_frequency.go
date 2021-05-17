//Letter package gives us ways to investigate how letters are being used in texts
package letter

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

//ConcurrentFrequency takes a slice of strings and uses a goroutine to get a map of the various runes being used in the text of each value in the slice
//it then merges the resulting maps to return the number of each rune in the given slice of strings
func ConcurrentFrequency(input []string) FreqMap {
	c := make(chan map[rune]int)
	for _, v := range input {
		v := v
		go func() {
			c <- Frequency(v)
		}()
	}

	result := FreqMap{}
	for i := 0; i < len(input); i++ {
		x := <-c
		for k, v := range x {
			result[k] += v
		}
	}
	return result
}
