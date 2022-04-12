package usecase

import (
	"rest/substr/usecase"
	"testing"
)

func TestMultiply(t *testing.T) {

	testCase := map[string]string{
		"hello":       "hel",
		"qwertyuiop":  "qwertyuiop",
		"bbbbbbbbbb":  "b",
		"qwzwsx":      "zwsx",
		"GreadisGood": "readisGo",
		"":            "",
	}
	for value, Expected := range testCase {
		Sub := usecase.NewSubstrUseCase()
		got := Sub.FindLongestSubstring(value)
		if Expected != got {
			t.Errorf("Expected '%v', but got '%v'", Expected, got)
		}
	}
}
