package main

import (
	"bufio"
	"fmt"
	"os"
)

func mustReadLine(r *bufio.Reader) string {
	line := ""
	isPrefix := true
	for isPrefix {
		var subLine []byte
		var err error
		subLine, isPrefix, err = r.ReadLine()
		if err != nil {
			panic(err)
		}
		line += string(subLine)
	}
	return line
}

func isAnagramm(l1, l2 string) bool {
	chars1 := make(map[rune]int32)
	chars2 := make(map[rune]int32)
	for _, r := range l1 {
		chars1[r]++
	}
	for _, r := range l2 {
		chars2[r]++
	}
	isEqual := len(chars1) == len(chars2)
	if isEqual {
		for r, c := range chars1 {
			if chars2[r] != c {
				isEqual = false
				break
			}
		}
	}
	return isEqual
}

func main() {
	r := bufio.NewReader(os.Stdin)

	l1 := mustReadLine(r)
	l2 := mustReadLine(r)

	b := isAnagramm(l1, l2)
	if b {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
