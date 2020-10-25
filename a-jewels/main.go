package main

import (
	"bufio"
	"fmt"
	"os"
)

func mustReadLine(r *bufio.Reader) string {
	line, isPrefix, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	if isPrefix {
		panic("input line is too long")
	}
	return string(line)
}

func countJewels(J, S string) int {
	jewels := make(map[rune]bool)
	for _, j := range J {
		jewels[j] = true
	}

	n := 0
	for _, s := range S {
		if jewels[s] {
			n++
		}
	}

	return n
}

func main() {
	r := bufio.NewReader(os.Stdin)

	J := mustReadLine(r)
	S := mustReadLine(r)

	n := countJewels(J, S)

	fmt.Println(n)
}
