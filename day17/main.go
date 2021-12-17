// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

var example = `target area: x=20..30, y=-10..-5`
var input = `target area: x=135..155, y=-102..-78`

func main() {
	s := NewSolver()

	min := point{x: 20, y: -10}
	max := point{x: 30, y: -5}
	fmt.Println("part1 (example):", s.WithInput(min, max).Part1()) // 45

	min = point{x: 135, y: -102}
	max = point{x: 155, y: -78}
	fmt.Println("part1 (input):", s.WithInput(min, max).Part1()) // 5151

	min = point{x: 20, y: -10}
	max = point{x: 30, y: -5}
	fmt.Println("part2 (example):", s.WithInput(min, max).Part2()) // 112

	min = point{x: 135, y: -102}
	max = point{x: 155, y: -78}
	fmt.Println("part2 (input):", s.WithInput(min, max).Part2()) // 968
}

type Solver struct {
	min point
	max point
}

func NewSolver() *Solver {
	return &Solver{}
}

type point struct {
	x, y int
}

func (s *Solver) WithInput(min, max point) *Solver {
	s.min = min
	s.max = max
	return s
}

func (s *Solver) Part1() int {
	res := 0
	dy := s.min.y
	if dy < 0 {
		dy = -dy
	}
	dy *= 10

	for x := 1; x <= s.max.x; x++ {
		for y := -dy; y < dy; y++ {
			vel := point{x: x, y: y}

			maxY := 0
			p := point{}
			for p.x+vel.x <= s.max.x && p.y+vel.y >= s.min.y {
				p.x += vel.x
				p.y += vel.y
				vel.y -= 1
				if vel.x < 0 {
					vel.x += 1
				} else if vel.x > 0 {
					vel.x -= 1
				}
				if p.y > maxY {
					maxY = p.y
				}
			}
			if p.x >= s.min.x && p.x <= s.max.x && p.y >= s.min.y && p.y <= s.max.y {
				if maxY > res {
					res = maxY
				}
			}
		}
	}

	return res
}

func (s *Solver) Part2() int {
	uniq := make(map[point]bool)
	dy := s.min.y
	if dy < 0 {
		dy = -dy
	}
	dy *= 10

	for x := 1; x <= s.max.x; x++ {
		for y := -dy; y < dy; y++ {
			vel := point{x: x, y: y}

			p := point{}
			for p.x+vel.x <= s.max.x && p.y+vel.y >= s.min.y {
				p.x += vel.x
				p.y += vel.y
				vel.y -= 1
				if vel.x < 0 {
					vel.x += 1
				} else if vel.x > 0 {
					vel.x -= 1
				}
			}
			if p.x >= s.min.x && p.x <= s.max.x && p.y >= s.min.y && p.y <= s.max.y {
				uniq[point{x: x, y: y}] = true
			}
		}
	}

	return len(uniq)
}
