package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var notes = []string{"c", "c#", "d", "d#", "e", "f", "f#", "g", "g#", "a", "a#", "b"}
var pattern = []string{"C", "A", "G", "E", "D"}

func main() {
	numberOfPatterns := os.Args[1]

	r, err := strconv.Atoi(numberOfPatterns)
	if err != nil {
		fmt.Printf("please pass an integer %v", err)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < r; i++ {
		note := getNote()
		pattern := getPattern()

		fmt.Printf("play the scale for %s in %s pattern\n", note, pattern)

	}

}

func getNote() string {
	i := randWithRange(0, len(notes))
	return notes[i]
}

func getPattern() string {
	i := randWithRange(0, len(pattern))
	return pattern[i]
}

func randWithRange(min, max int) int {
	return min + rand.Intn(max-min)

}
