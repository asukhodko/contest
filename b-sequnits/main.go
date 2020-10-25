package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
func mustReadInt(r *bufio.Reader) int {
	l := mustReadLine(r)
	n, err := strconv.Atoi(l)
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	r := bufio.NewReader(os.Stdin)

	maxSeq := 0
	curSeq := 0

	n := mustReadInt(r)
	for ; n > 0; n-- {
		v := mustReadInt(r)
		if v == 1 {
			curSeq++
		} else {
			if curSeq > maxSeq {
				maxSeq = curSeq
			}
			curSeq = 0
		}
	}
	if curSeq > maxSeq {
		maxSeq = curSeq
	}

	fmt.Println(maxSeq)
}
