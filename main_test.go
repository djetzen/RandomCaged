package main

import (
	"testing"
	"unicode/utf8"
)

func TestStrummingPrint(t *testing.T) {
	testCases := map[string]struct {
		randomInput   int
		strumStep     int
		expectedStrum string
	}{
		"random input 1 prints pattern with arrow up if step is odd": {
			randomInput:   1,
			strumStep:     1,
			expectedStrum: "↑",
		},
		"random input 1 prints pattern with arrow down if step is even": {
			randomInput:   1,
			strumStep:     2,
			expectedStrum: "↓",
		},
		"a dot is printed if random input is 0": {
			randomInput:   0,
			strumStep:     1,
			expectedStrum: ".",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actualStrum := getStrum(testCase.randomInput, testCase.strumStep)
			if actualStrum != testCase.expectedStrum {
				t.Fatalf("expected other strum. Got %v but should be %v", actualStrum, testCase.expectedStrum)
			}
		})
	}
}

func TestStrummingPatternIsAlwaysAFourFourTime(t *testing.T) {
	strummingPattern := getStrummingPattern()
	if utf8.RuneCountInString(strummingPattern) != 8 {
		t.Fatalf("lenght must always be eight digits, but was %v", len(strummingPattern))
	}
}
