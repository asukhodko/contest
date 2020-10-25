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

func generate(n, o, c int, res string) {
	if c == n {
		fmt.Println(res)
		return
	}
	if o < n {
		generate(n, o+1, c, res+"(")
	}
	if c < o {
		generate(n, o, c+1, res+")")
	}
}

func printParentheses(n int) {
	generate(n, 0, 0, "")
}

func main() {
	r := bufio.NewReader(os.Stdin)

	n := mustReadInt(r)

	printParentheses(n)

	fmt.Println()
}
