// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part1 (example):", solve(example, 18))  // 26
	fmt.Println("part1 (input):", solve(input, 80))      // 380243
	fmt.Println("part2 (example):", solve(example, 256)) // 26984457539
	fmt.Println("part2 (input):", solve(input, 256))     // 1708791884591
}

func parse(input string) []string {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, ",")
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = strings.TrimSpace(line)
	}
	return result
}

func solve(input string, days int) int {
	lines := parse(input)

	result := make(map[int]int)
	for _, line := range lines {
		result[toInt(line)] += 1
	}

	for n := 0; n < days; n++ {
		created := make(map[int]int)
		for k := 0; k < 10; k++ {
			v, ok := result[k]
			if !ok {
				continue
			}
			result[k] -= v
			if k == 0 {
				created[6] += v
				created[8] += v
			} else {
				result[k-1] += v
			}
		}
		for k, v := range created {
			result[k] += v
		}
	}

	var total int
	for _, v := range result {
		total += v
	}
	return total
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `3,4,3,1,2 `
var input = `5,1,1,1,3,5,1,1,1,1,5,3,1,1,3,1,1,1,4,1,1,1,1,1,2,4,3,4,1,5,3,4,1,1,5,1,2,1,1,2,1,1,2,1,1,4,2,3,2,1,4,1,1,4,2,1,4,5,5,1,1,1,1,1,2,1,1,1,2,1,5,5,1,1,4,4,5,1,1,1,3,1,5,1,2,1,5,1,4,1,3,2,4,2,1,1,4,1,1,1,1,4,1,1,1,1,1,3,5,4,1,1,3,1,1,1,2,1,1,1,1,5,1,1,1,4,1,4,1,1,1,1,1,2,1,1,5,1,2,1,1,2,1,1,2,4,1,1,5,1,3,4,1,2,4,1,1,1,1,1,4,1,1,4,2,2,1,5,1,4,1,1,5,1,1,5,5,1,1,1,1,1,5,2,1,3,3,1,1,1,3,2,4,5,1,2,1,5,1,4,1,5,1,1,1,1,1,1,4,3,1,1,3,3,1,4,5,1,1,4,1,4,3,4,1,1,1,2,2,1,2,5,1,1,3,5,2,1,1,1,1,1,1,1,4,4,1,5,4,1,1,1,1,1,2,1,2,1,5,1,1,3,1,1,1,1,1,1,1,1,1,1,2,1,3,1,5,3,3,1,1,2,4,4,1,1,2,1,1,3,1,1,1,1,2,3,4,1,1,2`
