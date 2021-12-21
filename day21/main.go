// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	fmt.Println("part 1:", part1(4-1, 8-1, 1000)) // 739785
	fmt.Println("part 1:", part1(7-1, 5-1, 1000)) // 798147
	l, r := part2(4-1, 8-1, 0, 0)
	fmt.Println("part 2:", max(l, r)) // 798147
	l, r = part2(7-1, 5-1, 0, 0)
	fmt.Println("part 2:", max(l, r)) // 798147
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(i, j, tgt int) int {
	score := roll(tuple{i, j}, 1000)
	if score.i > score.j {
		return score.j * score.k
	}
	return score.i * score.k
}

var dp = make(map[tuple3]tuple)

func part2(p0, p1, s0, s1 int) (int, int) {
	if s0 >= 21 {
		return 1, 0
	}
	if s1 >= 21 {
		return 0, 1
	}

	key := tuple3{p0, p1, s0, s1}
	if v, exists := dp[key]; exists {
		return v.i, v.j
	}

	left, right := 0, 0
	rng := []int{1, 2, 3}
	for _, a := range rng {
		for _, b := range rng {
			for _, c := range rng {
				newp0 := (p0 + a + b + c) % 10
				news0 := s0 + newp0 + 1
				l, r := part2(p1, newp0, s1, news0)
				left += r
				right += l
			}
		}
	}

	dp[key] = tuple{left, right}
	return left, right
}

type tuple struct{ i, j int }
type tuple2 struct{ i, j, k int }
type tuple3 struct{ i, j, k, l int }

func roll(pos tuple, tgt int) (score tuple2) {
	incr := 9
	step := 6
	for {
		if score.k%2 == 0 {
			pos.i += step
			pos.i %= 10
			score.i += pos.i + 1
		} else {
			pos.j += step
			pos.j %= 10
			score.j += pos.j + 1
		}
		score.k++
		if score.i >= tgt || score.j >= tgt {
			score.k *= 3
			return
		}
		step += incr
	}
}
