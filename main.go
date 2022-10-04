package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var notes = []string{"c", "c#", "d", "d#", "e", "f", "f#", "g", "g#", "a", "a#", "b"}
var cagedPatternBase = []string{"C", "A", "G", "E", "D"}

func main() {
	var argNumberOfPatterns string
	if len(os.Args) < 2 {
		argNumberOfPatterns = "1"
	} else {
		argNumberOfPatterns = os.Args[1]
	}

	numberOfPatterns, err := strconv.Atoi(argNumberOfPatterns)
	if err != nil {
		fmt.Printf("please pass an integer as argument %v", err)
	}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numberOfPatterns; i++ {
		note := getNote()
		pattern := getCagedPattern()
		strummingPattern := getStrummingPattern()
		fmt.Printf("play the scale for %s in %s caged pattern with %s strumming pattern\n", note, pattern, strummingPattern)
	}

}

func getNote() string {
	i := rand.Intn(len(notes))
	return notes[i]
}

func getCagedPattern() string {
	i := rand.Intn(len(cagedPatternBase))
	return cagedPatternBase[i]
}

func getStrummingPattern() string {
	strummingPattern := ""
	for strum := 0; strum < 8; strum++ {
		randomStrum := rand.Intn(2)
		strummingPattern = strummingPattern + getStrum(randomStrum, strum)
	}
	return strummingPattern
}

func getStrum(randomStrum int, strum int) string {
	if randomStrum == 1 {
		if strum%2 == 0 {
			return "↓"
		} else {
			return "↑"
		}
	} else {
		return "."
	}
}
