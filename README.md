# advent-of-code-2021
AOC in go


## Template for Playground

I use [golang playground](https://go.dev/play) to avoid auto-completion from Github CoPilot.

```go
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part1 (example):", part1(example))
	//fmt.Println("part1 (input):", part1(input))
	//fmt.Println("part2 (example):", part2(example))
	//fmt.Println("part2 (input):", part2(input))
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

func part1(input string) int {
	lines := parse(input)
	_ = lines
	return 0
}

func part2(input string) int {
	lines := parse(input)
	_ = lines
	return 0
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = ``
var input = ``
```


## Template V2 for Playground

Since there's always a lot of logic reuse between part 1 and part 2, I end up copying the logic rather than writing functions to reuse. Testing if using a single struct to share state would be better.

```go
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := NewSolver()
	fmt.Println("part1 (example):", s.WithInput(example).Part1()) // 0
	fmt.Println("part1 (input):", s.WithInput(input).Part1())     // 0
	fmt.Println("part2 (example):", s.WithInput(example).Part2()) // 0
	fmt.Println("part2 (input):", s.WithInput(input).Part2())     // 0
}

type Solver struct {
	input string
	state []string
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) WithInput(input string) *Solver {
	s.input = input
	s.state = parseLines(parse(s.input))
	return s
}

func (s *Solver) Part1() int {
	state := s.state
	return len(state)
}

func (s *Solver) Part2() int {
	state := s.state
	return len(state)
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

func parseLines(lines []string) []string {
	return lines
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = ``
var input = ``
```
