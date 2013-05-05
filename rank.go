package main

import (
	"sort"
)

type rank struct {
	letters []rune
	counts  []int
}

func (r rank) Len() int {
	return len(r.letters)
}

func (r rank) Swap(i, j int) {
	r.letters[i], r.letters[j] = r.letters[j], r.letters[i]
	r.counts[i], r.counts[j] = r.counts[j], r.counts[i]
}

func (r rank) Less(i, j int) bool {
	a, b := r.counts[i], r.counts[j]
	if a == b {
		return r.letters[i] < r.letters[j]
	}
	return a > b
}

func Rank(raw map[rune]int) []rune {
	var r rank
	for key, value := range raw {
		r.letters = append(r.letters, key)
		r.counts = append(r.counts, value)
	}
	sort.Sort(r)
	return r.letters
}
