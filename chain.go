// Copyright 2013 Dario Castañé.  All rights reserved.
// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Generating words in data-driven way: a Modified Markov chain algorithm

Based on the program presented in the "Design and Implementation" chapter
of The Practice of Programming (Kernighan and Pike, Addison-Wesley 1999).
See also Computer Recreations, Scientific American 260, 122 - 125 (1989).

A Markov chain algorithm generates text by creating a statistical model of
potential textual suffixes for a given prefix.

To generate text using this table we select an initial prefix ("d", for
example), choose one of the suffixes associated with that prefix by Lango
Algorithm (check select() function), and then create a new prefix by removing
the first word from the prefix and appending the suffix (making the new prefix is "da",
supposing next is "a").

Repeat this process until we can't find any suffixes for the current prefix 
or we exceed the length limit.
*/
package main

import (
	"math"
	"fmt"
	"strings"
)

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain      map[string][]rune
	prefixLen  int
	vowels     int
	consonants int
}

// Prefix is a Markov chain prefix of one or more runes.
type Prefix []rune

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return string(p)
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(letter rune) {
	copy(p, p[1:])
	p[len(p) - 1] = letter
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]rune), prefixLen, 0, 0}
}

func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func (c *Chain) Build(words []string) {
	vowels, consonants := 0.0, 0.0
	for _, word := range words {
		p := make(Prefix, c.prefixLen)
		for _, letter := range word {
			key := p.String()
			if isVowel(letter) {
				vowels++
			} else {
				consonants++
			}
			c.chain[key] = append(c.chain[key], letter)
			p.Shift(letter)
		}
	}
	processed := float64(len(words))
	c.vowels = int(math.Floor(vowels / processed))
	c.consonants = int(math.Floor(consonants / processed))
	for prefix, letters := range c.chain {
		counts := make(map[rune]int)
		for _, letter := range letters {
			counts[letter]++
		}
		c.chain[prefix] = Rank(counts)
	}
}

func (c *Chain) chooseNext(candidates []rune, pool []int) (next rune) {
	next = 0
	for _, candidate := range candidates {
		if isVowel(candidate) && pool[0] > 0 {
			next = candidate
			pool[0]--
			break
		}
		if !isVowel(candidate) && pool[1] > 0 {
			next = candidate
			pool[1]--
			break
		}
	}
	return
}

func (c *Chain) Generate() string {
	p := make(Prefix, c.prefixLen)
	pool := []int{c.vowels, c.consonants}
	fmt.Printf("%v\n", pool)
	var word []rune
	for i := 0; i < c.vowels + c.consonants; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := c.chooseNext(choices, pool)
		if next == 0 {
			fmt.Printf("'%s' => '%v' => '(%v)'\n", p.String(), strings.Replace(string(choices), "\n", "\\n", -1), pool)
			p.Shift(choices[0])
			i--
			continue
		}
		fmt.Printf("'%s' => '%v' => '%c(%v)'\n", p.String(), strings.Replace(string(choices), "\n", "\\n", -1), next, pool)
		word = append(word, next)
		p.Shift(next)
	}
	return string(word)
}
