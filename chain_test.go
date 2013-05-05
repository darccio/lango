package main

import (
	"testing"
)

func TestSpainLanguages(t *testing.T) {
	words := []string{ "kaixo", "hola", "ola", "hola" }
	c := NewChain(2)
	c.Build(words)
	if word := c.Generate(); word != "hola" {
		t.Fatalf("wrong generation: %s instead of hola", word)
	}
}
