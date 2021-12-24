// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var example = `inp z
inp x
mul z 3
eql z x`

func main() {
	fmt.Println("part1 1:", part1(input, 91141123491111))

}

// w, 0, 0, w+11+26
func in12(wi, _xi, _yi, zi int) (w, x, y, z int) {
	w = wi
	x = zi % 26 // Doesn't matter
	y = 0       // Doesn't matter
	z = zi / 26

	x -= 11
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	z *= (25 * x) + 1
	y = (w + 6) * x
	z += y
	return
}

// w, 0, 0, w + 26
func in13(wi, _xi, _yi, zi int) (w, x, y, z int) {
	w = wi
	x = zi % 26 // Doesn't matter
	y = 0       // Doesn't matter
	z = zi / 26

	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	z *= (25 * x) + 1
	y = (w + 15) * x
	z += y
	return
}

func copyMap(m map[string]int) map[string]int {
	res := make(map[string]int)
	for k, v := range m {
		res[k] = v
	}
	return res
}

type tuple struct {
	w, z int
}

func part1(input string, _n int) int {
	ins := parse(input)

	state := make(map[int]map[tuple]int)
	for i := 0; i < len(ins); i++ {
		state[i] = make(map[tuple]int)
		prev, ok := state[i-1]
		for w := 1; w < 10; w++ {
			res := make(map[string]int)
			if ok {
				for _, prevz := range prev {
					res["z"] = prevz
					exec(ins[i], res, w)
					state[i][tuple{w, prevz}] = res["z"]
				}
			} else {
				fmt.Println("i", i, "w", w)
				fmt.Println("before", res)
				exec(ins[i], res, w)
				fmt.Println("after", res)
				fmt.Println()
				z := res["z"]
				state[i][tuple{w, 0}] = z
			}
		}
	}
	fmt.Println(state)
	return 0
}

func exec(ins []instruction, res map[string]int, n int) {
	for _, in := range ins {
		lhs := res[in.lhs]
		rhs := 0
		if in.rhs != nil {
			if in.isRhsVar {
				rhs = res[*in.rhs]
			} else {
				rhs = toInt(*in.rhs)
			}
		}
		switch in.op {
		case "inp":
			res[in.lhs] = n
		case "add":
			res[in.lhs] = lhs + rhs
		case "div":
			res[in.lhs] = lhs / rhs
		case "mod":
			res[in.lhs] = lhs % rhs
		case "mul":
			res[in.lhs] = lhs * rhs
		case "eql":
			if lhs == rhs {
				res[in.lhs] = 1
			} else {
				res[in.lhs] = 0
			}
		default:
			panic("not handled: " + in.op)
		}
	}
}

type instruction struct {
	op       string
	lhs      string
	rhs      *string
	isRhsVar bool
}

var re *regexp.Regexp

func init() {
	var err error
	re, err = regexp.Compile(`-?\d+`)
	if err != nil {
		panic(err)
	}
}

func parse(input string) map[int][]instruction {
	lines := strings.Split(input, "\n")
	res := make(map[int][]instruction)
	i := -1
	for _, line := range lines {
		parts := strings.Fields(line)
		switch len(parts) {
		case 2:
			i++
			if parts[0] != "inp" {
				panic("invalid instruction:" + line)
			}
			lhs := parts[1]
			res[i] = append(res[i], instruction{op: "inp", lhs: lhs})
		case 3:
			op := parts[0]
			lhs := parts[1]
			rhs := parts[2]
			isRhsVar := re.MatchString(rhs)
			res[i] = append(res[i], instruction{op: op, lhs: lhs, rhs: &rhs, isRhsVar: !isRhsVar})
		}
	}
	return res
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var input = `inp w
mul x 0
add x z
mod x 26
div z 1
add x 11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 1
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 10
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 10
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 13
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 2
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -10
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 5
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 6
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 0
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 12
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 16
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 12
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -7
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 15
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 13
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 7
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -13
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 6
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x 0
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 5
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 6
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x 0
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 15
mul y x
add z y`
