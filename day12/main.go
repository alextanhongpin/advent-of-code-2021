// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part1 (example):", part1(example))  // 10
	fmt.Println("part1 (example):", part1(example1)) // 19
	fmt.Println("part1 (example):", part1(example2)) // 226
	fmt.Println("part1 (input):", part1(input))      // 5333
	fmt.Println("part2 (example):", part2(example))  // 36
	fmt.Println("part2 (example):", part2(example1)) // 103
	fmt.Println("part2 (example):", part2(example2)) // 3509
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

	m := make(map[string]map[string]bool)
	for _, line := range lines {
		paths := strings.Split(line, "-")
		from, to := paths[0], paths[1]
		if m[from] == nil {
			m[from] = make(map[string]bool)
		}
		if m[to] == nil {
			m[to] = make(map[string]bool)
		}
		m[from][to] = true
		m[to][from] = true
	}

	var count int
	stack := [][]string{{"start"}}
	for len(stack) > 0 {
		var last []string
		stack, last = stack[:len(stack)-1], stack[len(stack)-1]
		tail := last[len(last)-1]
		for k := range m[tail] {
			visited := make(map[string]bool)
			for _, door := range last {
				visited[door] = true
			}
			if strings.ToLower(k) == k && visited[k] {
				continue
			}

			t := make([]string, len(last))
			copy(t, last)
			t = append(t, k)

			if k == "end" {
				count++
			}
			stack = append(stack, t)
		}
	}
	return count
}

func part2(input string) int {
	lines := parse(input)

	m := make(map[string]map[string]bool)
	for _, line := range lines {
		paths := strings.Split(line, "-")
		from, to := paths[0], paths[1]
		if m[from] == nil {
			m[from] = make(map[string]bool)
		}
		if m[to] == nil {
			m[to] = make(map[string]bool)
		}
		m[from][to] = true
		m[to][from] = true
	}

	stack := [][]string{{"start"}}
	var count int
	for len(stack) > 0 {
		var last []string
		stack, last = stack[:len(stack)-1], stack[len(stack)-1]
		tail := last[len(last)-1]
		for k := range m[tail] {
			visited := make(map[string]int)
			valid := true
			twice := 0
			for _, door := range last {
				visited[door]++
				if strings.ToLower(door) == door && visited[door] >= 2 {
					twice++
					if twice > 1 {
						valid = false
						break
					}
				}
				if door == "start" && visited["start"] > 1 {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			t := make([]string, len(last))
			copy(t, last)
			t = append(t, k)

			if k == "end" {
				count++
				continue
			}
			stack = append(stack, t)
		}
	}
	return count
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var example1 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var example2 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

var input = `pq-GX
GX-ah
mj-PI
ey-start
end-PI
YV-mj
ah-iw
te-GX
te-mj
ZM-iw
te-PI
ah-ZM
ey-te
ZM-end
end-mj
te-iw
te-vc
PI-pq
PI-start
pq-ey
PI-iw
ah-ey
pq-iw
pq-start
mj-GX`
