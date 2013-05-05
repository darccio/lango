package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
)

func readLine(r *bufio.Reader) (ln []byte, err error) {
	var (
		isPrefix bool = true
		line []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return
}

func readWords(filepath string) (words []string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for line, err := readLine(r); err == nil; line, err = readLine(r) {
		words = append(words, string(line))
	}
	return
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Usage: lng file")
	}
	c := NewChain(1)
	words := readWords(os.Args[1])
	c.Build(words)
	fmt.Println(c.Generate())
}
