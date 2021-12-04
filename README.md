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
	"strings"
)

func main() {
	fmt.Println("part1 (example):", part1(example))
	fmt.Println("part1 (input):", part1(input))
	fmt.Println("part2 (example):", part2(example))
	fmt.Println("part2 (input):", part2(input))
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

var example = ``
var input = ``

```
