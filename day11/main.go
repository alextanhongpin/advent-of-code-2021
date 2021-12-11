// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part1 (example):", part1(example, 1))    // 9
	fmt.Println("part1 (example):", part1(example2, 10))  // 204
	fmt.Println("part1 (example):", part1(example2, 100)) // 1656
	fmt.Println("part1 (input):", part1(input, 100))      // 1705
	fmt.Println("part2 (example):", part2(example2))      // 195
	fmt.Println("part2 (input):", part2(input))           // 265
}

func parse(input string) []string {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = strings.TrimSpace(line)
	}
	return result
}

type point struct {
	x, y int
}

var adjacents = []point{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func part1(input string, rounds int) int {
	lines := parse(input)

	m := make(map[point]int)
	for y, line := range lines {
		for x, ch := range line {
			p := point{x: x, y: y}
			m[p] = toInt(string(ch))
		}
	}

	flashed := make(map[point]bool)
	flash := 0

	var do_flash func(p point)
	do_flash = func(p point) {
		if flashed[p] {
			return
		}
		m[p]++
		if m[p] <= 9 {
			return
		}

		flash++
		flashed[p] = true

		for _, adj := range adjacents {
			c := p
			c.x += adj.x
			c.y += adj.y
			_, ok := m[c]
			if !ok {
				continue
			}
			do_flash(c)
		}

		m[p] = 0
	}

	for i := 0; i < rounds; i++ {
		for c := range m {
			do_flash(c)
		}
		// Reset at the end of every round
		flashed = make(map[point]bool)
	}

	return flash
}

func part2(input string) int {
	lines := parse(input)

	m := make(map[point]int)

	for y, line := range lines {
		for x, ch := range line {
			p := point{x: x, y: y}
			m[p] = toInt(string(ch))
		}
	}

	flashed := make(map[point]bool)
	round := 0

	var do_flash func(p point)
	do_flash = func(p point) {
		if flashed[p] {
			return
		}

		m[p]++
		if m[p] < 10 {
			return
		}

		flashed[p] = true

		for _, adj := range adjacents {
			c := p
			c.x += adj.x
			c.y += adj.y
			_, ok := m[c]
			if !ok {
				continue
			}
			do_flash(c)
		}

		m[p] = 0
	}

	for {
		round++
		for c := range m {
			do_flash(c)
		}

		// Reset at the end of every round
		flashed = make(map[point]bool)

		synchronized := true
		for _, n := range m {
			if n > 0 {
				synchronized = false
				break
			}
		}

		if synchronized {
			return round
		}
	}
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `11111
19991
19191
19991
11111`

var example2 = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var input = `3113284886
2851876144
2774664484
6715112578
7146272153
6256656367
3148666245
3857446528
7322422833
8152175168`
