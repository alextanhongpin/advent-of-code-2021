// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1", magnitude(sum(input))) // 4033
	fmt.Println("Part 2", max(input))            // 4864
}

func max(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	result := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			m := magnitude(sum(strings.Join([]string{lines[i], lines[j]}, "\n")))
			if m > result {
				result = m
			}
		}
	}
	return result
}

func explode(input string) string {
	tokens := tokenize(input)
	depth := 0

	for i, ch := range tokens {
		switch ch {
		case "[":
			depth++
		case "]":
			depth--
		case ",":
		default:
			if depth > 4 && isInt(tokens[i]) && isInt(tokens[i+2]) {
				left, right := toInt(tokens[i]), toInt(tokens[i+2])
				for j := i - 1; j > -1; j-- {
					if isInt(tokens[j]) {
						tokens[j] = strconv.Itoa(toInt(tokens[j]) + left)
						break
					}
				}

				for j := i + 3; j < len(tokens); j++ {
					if isInt(tokens[j]) {
						tokens[j] = strconv.Itoa(toInt(tokens[j]) + right)
						break
					}
				}

				// Explode.
				tokens = append(tokens[:i-1], append([]string{"0"}, tokens[i+4:]...)...)
				return strings.Join(tokens, "")
			}
		}
	}
	return input
}

func split(input string) string {
	tokens := tokenize(input)

	for i, ch := range tokens {
		if !isInt(ch) {
			continue
		}
		n := toInt(ch)
		if n > 9 {
			replacement := []string{
				"[", strconv.Itoa(n / 2), ",", strconv.Itoa(n - n/2), "]",
			}
			tokens = append(tokens[:i], append(replacement, tokens[i+1:]...)...)
			return strings.Join(tokens, "")
		}
	}
	return input
}

func sum(input string) string {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	line := lines[0]
	for i := 1; i < len(lines); i++ {
		next := strings.TrimSpace(lines[i])
		line = fmt.Sprintf("[%s,%s]", line, next)
		line = solve(line)
	}
	return solve(line)
}

func solve(input string) string {
	output := input
	for {
		input = output
		output = explode(input)
		if output != input {
			continue
		}
		input = output
		output = split(input)
		if output != input {
			continue
		}
		return output
	}
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func magnitude(input string) int {
	tokens := tokenize(input)
	if len(tokens) == 1 {
		return toInt(tokens[0])
	}
	for i := range tokens {
		if isInt(tokens[i]) && isInt(tokens[i+2]) {
			left, right := toInt(tokens[i]), toInt(tokens[i+2])
			mg := 3*left + 2*right
			tokens = append(tokens[:i-1], append([]string{strconv.Itoa(mg)}, tokens[i+4:]...)...)
			return magnitude(strings.Join(tokens, ""))
		}
	}
	return 0
}

func tokenize(input string) []string {
	input = strings.Split(strings.TrimSpace(input), "\n")[0]
	re, err := regexp.Compile(`(\[|\]|\d+|,)`)
	if err != nil {
		panic(err)
	}
	matches := re.FindAllStringSubmatch(input, -1)
	result := make([]string, len(matches))
	for i, match := range matches {
		result[i] = match[1]
	}
	return result
}

var (
	example = `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`

	input = `[[[[7,1],[0,0]],[6,[8,2]]],[8,[3,8]]]
[[[3,6],[9,4]],[[[5,9],5],[8,0]]]
[[[2,2],2],[1,[[1,6],7]]]
[[[[0,9],7],[[3,2],8]],[6,[7,9]]]
[[[[4,1],6],[[7,6],[2,2]]],[[[1,1],9],4]]
[[[8,[3,7]],3],[[4,4],[[9,1],[3,5]]]]
[[4,[8,2]],[1,[0,5]]]
[8,[8,7]]
[[[[2,2],7],[3,[4,5]]],[[4,6],[[2,5],4]]]
[[[5,5],[[5,1],3]],[[2,[8,2]],[[6,9],[1,5]]]]
[0,7]
[[[[5,1],3],[8,[5,3]]],7]
[[5,[2,[0,6]]],[[[5,5],2],[9,[8,0]]]]
[[[[3,4],2],0],4]
[[[[5,3],[2,7]],6],[[4,0],[9,[7,2]]]]
[[[3,[2,5]],[3,3]],7]
[[[[5,1],1],[4,8]],[[5,[8,3]],2]]
[[4,[[8,1],[8,5]]],[[[4,1],0],6]]
[[[5,5],[5,9]],[0,[[6,8],[0,1]]]]
[4,[[[7,9],4],0]]
[[[[0,1],7],[[3,6],5]],[8,[5,[6,1]]]]
[[[7,7],[8,0]],[6,[8,[7,9]]]]
[[[9,2],1],6]
[[[4,4],[2,[5,0]]],[[[2,6],6],[5,[4,3]]]]
[[2,[[4,7],5]],1]
[[8,7],[[[2,0],7],[1,[0,3]]]]
[[9,[[9,3],[9,5]]],[[8,7],[[4,1],[6,5]]]]
[[3,4],[[9,4],5]]
[[5,[[8,3],5]],1]
[[0,[[9,0],[3,2]]],[2,[7,[5,1]]]]
[[9,[[9,5],[8,6]]],[[4,4],[[3,8],[1,6]]]]
[[[1,[5,2]],9],[[4,6],[3,[8,0]]]]
[[1,7],[[1,7],9]]
[[[[3,4],3],[[7,5],[9,1]]],[[[5,0],[3,0]],[[7,9],6]]]
[[[7,2],[[1,0],[5,6]]],[[[3,7],[8,9]],6]]
[[[[1,1],1],[[8,6],[9,8]]],[[[1,8],4],[8,9]]]
[[[8,9],0],3]
[[[1,7],[1,[3,9]]],[6,[0,[8,5]]]]
[[0,5],[6,5]]
[[[[6,8],[4,5]],[[7,4],6]],[[3,6],5]]
[[8,[[0,9],8]],[9,[7,[7,9]]]]
[0,[[[7,1],2],[[0,4],4]]]
[[0,[[9,1],5]],[1,4]]
[3,4]
[[[9,3],[1,3]],[[[4,8],3],[[1,3],[9,0]]]]
[[[[5,1],7],[[9,2],8]],[[[6,8],[5,4]],[0,1]]]
[8,[[1,[3,0]],[[7,9],4]]]
[[[6,4],[[2,9],[9,0]]],[7,[[0,0],3]]]
[[3,[[9,6],6]],2]
[[5,[[3,1],[7,5]]],[[[6,7],9],[[4,6],[5,2]]]]
[[[4,[6,5]],8],[[6,[8,0]],[[9,3],3]]]
[[[[4,9],[2,8]],9],[[[5,0],0],[[3,4],[2,8]]]]
[[3,[7,1]],[9,[[1,8],7]]]
[[9,1],[0,[[0,7],[7,1]]]]
[[7,[0,[7,6]]],[[[5,3],1],[6,[4,5]]]]
[8,[[[2,1],[6,9]],[[3,3],[4,6]]]]
[0,[7,[3,0]]]
[[[[1,6],3],[5,[8,0]]],[[[6,6],7],1]]
[[[7,[8,3]],3],[[[2,8],5],[0,[9,5]]]]
[[[[5,1],4],[[1,2],1]],7]
[[[3,[7,5]],7],3]
[[9,[6,[1,1]]],[[[4,1],[2,2]],[[9,5],[7,7]]]]
[2,7]
[[[9,[8,6]],[[9,0],[6,5]]],[[[6,7],5],[[7,7],[2,3]]]]
[[[0,[6,4]],2],[4,[7,[7,5]]]]
[[[[6,1],[9,1]],[[6,1],9]],[[2,6],0]]
[[0,[[1,8],[3,5]]],[4,[[8,2],[4,2]]]]
[[[[9,3],[4,2]],2],[[[2,1],[7,1]],[4,8]]]
[[[3,[0,2]],3],8]
[[[4,[4,9]],9],[[[4,4],5],9]]
[[[[8,2],7],9],[[[1,0],[3,8]],[[7,7],0]]]
[[[3,2],[9,7]],[[9,[8,2]],[[5,5],3]]]
[[[7,[3,1]],[[8,3],1]],[[[8,6],[7,0]],4]]
[[9,[[9,1],5]],[[4,[1,1]],2]]
[[[[7,4],[0,3]],7],[8,[6,[3,3]]]]
[5,5]
[[6,7],[1,[7,[8,1]]]]
[[1,[0,4]],7]
[[[4,0],[[0,1],[2,2]]],[9,[[9,9],[3,0]]]]
[[[6,0],[[8,6],3]],[[5,1],[[8,1],[2,7]]]]
[[[[8,3],7],5],[9,[[5,1],8]]]
[[[[4,0],[5,2]],[[0,0],7]],2]
[[[[0,1],6],2],[[8,2],6]]
[[[[2,4],1],[[6,7],9]],[[[1,6],9],3]]
[[5,5],[[8,[7,7]],[5,8]]]
[[6,[[9,2],[9,7]]],[[[8,5],[4,4]],7]]
[[[9,[7,7]],[6,0]],[7,[[8,7],[1,2]]]]
[[7,[6,2]],[[9,[5,2]],[1,4]]]
[[[7,[5,9]],[[3,9],[4,5]]],[0,6]]
[[9,[8,[2,2]]],[[9,7],[1,1]]]
[[[[2,3],4],[[4,8],9]],[[9,[8,6]],[[0,9],0]]]
[[0,[[9,3],0]],[8,8]]
[[[[2,9],6],[[2,8],9]],[[[0,5],6],[[6,1],7]]]
[[9,[[8,3],[5,8]]],[[7,[3,0]],3]]
[[[4,[4,2]],0],1]
[[[[9,6],[5,8]],[6,2]],[[[8,0],[7,0]],[[5,6],4]]]
[[[8,0],[[4,3],[7,4]]],[[3,[7,9]],[[7,3],6]]]
[[3,[5,[0,3]]],[5,4]]
[[[[1,2],[6,3]],1],[[7,[5,2]],[[8,8],7]]]
[[4,[[8,0],[7,1]]],[[8,[8,0]],[[1,5],3]]]`
)
