package main

import (
	"testing"
	"reflect"
)

func TestSort(t *testing.T) {
	letters := map[rune]int{
		'k': 1,
		'a': 4,
		'i': 1,
		'x': 1,
		'o': 4,
		'h': 2,
		'l': 3,
	}
	expected := []rune { 'a', 'o', 'l', 'h', 'i', 'k', 'x' }
	newLetters := Rank(letters)
	if !reflect.DeepEqual(newLetters, expected) {
		t.Fatalf("Expecting %v, got %v", expected, newLetters)
	}
}
