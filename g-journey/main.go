package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func mustReadIntPair(r *bufio.Reader) (int, int) {
	l := mustReadLine(r)
	pair := strings.Split(l, " ")
	if len(pair) != 2 {
		panic("pair expected")
	}
	a, err := strconv.Atoi(pair[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(pair[1])
	if err != nil {
		panic(err)
	}
	return a, b
}

type city struct {
	n       int
	x       int
	y       int
	roads   []*road
	metric  int
	checked bool
}

type road struct {
	a, b     *city
	distance int
}

func mod(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)

	n := mustReadInt(r)
	cities := make([]*city, 0, n)
	for i := 1; i <= n; i++ {
		x, y := mustReadIntPair(r)
		newCity := &city{
			n:      i,
			x:      x,
			y:      y,
			metric: math.MaxInt32,
		}
		for _, c := range cities {
			distance := mod(newCity.x-c.x) + mod(newCity.y-c.y)
			newRoad := &road{
				a:        newCity,
				b:        c,
				distance: distance,
			}
			newCity.roads = append(newCity.roads, newRoad)
			c.roads = append(c.roads, newRoad)
		}
		cities = append(cities, newCity)
	}

	fuel := mustReadInt(r)
	from, to := mustReadIntPair(r)
	if from > len(cities)+1 {
		panic("from city number is to big")
	}
	if to > len(cities)+1 {
		panic("to city number is to big")
	}

	cityFrom := cities[from-1]
	cityTo := cities[to-1]

	cityFrom.metric = 0
	nRoads := findPath(cityFrom, cityTo, fuel)

	fmt.Println(nRoads)
}

func findPath(from *city, to *city, fuel int) int {
	if from == to {
		return 0
	}

	for _, r := range from.roads {
		if r.distance > fuel {
			continue
		}
		next := nextCity(from, r)
		if next.checked {
			continue
		}
		metric := from.metric + 1
		if next.metric > metric {
			next.metric = metric
		}
	}

	from.checked = true

	for _, r := range from.roads {
		if r.distance > fuel {
			continue
		}
		next := nextCity(from, r)
		if next.checked {
			continue
		}
		findPath(next, to, fuel)
	}
	nRoads := to.metric
	if nRoads == math.MaxInt32 {
		nRoads = -1
	}
	return nRoads
}

func nextCity(from *city, r *road) *city {
	if r.a != from {
		return r.a
	} else {
		return r.b
	}
}
