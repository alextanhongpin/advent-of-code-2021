// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := NewSolver()
	fmt.Println("part1 (example):", s.WithInput(example).Part1(14)) // 1588
	fmt.Println("part1 (input):", s.WithInput(input).Part1(10))     // 2068
	//fmt.Println("part2 (example):", s.WithInput(example).Part1(40)) // 0
	//fmt.Println("part2 (input):", s.WithInput(input).Part1(40))     // 0
}

type Solver struct {
	input    string
	reaction string
	react    map[string]string
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) WithInput(input string) *Solver {
	s.input = input
	s.reaction, s.react = parseLines(parse(s.input))
	return s
}

func (s *Solver) Part1(rounds int) int {
	r := s.reaction
	for i := 0; i < rounds; i++ {
		r = s.run(r)
	}

	count := make(map[rune]int)
	for _, ch := range r {
		count[ch]++
	}

	var result []int
	for _, v := range count {
		result = append(result, v)
	}
	sort.Ints(result)
	return result[len(result)-1] - result[0]
}

func (s *Solver) run(reaction string) string {
	var sb strings.Builder
	for i := 0; i < len(reaction)-1; i++ {
		curr := string(reaction[i])
		next := string(reaction[i+1])
		v, ok := s.react[curr+next]
		if ok {
			sb.WriteString(curr + v)
			if i == len(reaction)-2 {
				sb.WriteString(next)
			}
		} else {
			sb.WriteString(curr + next)
		}
	}
	return sb.String()
}

func (s *Solver) runReaction(counter map[string]int) map[string]int {
	newCounter := make(map[string]int)
	for k, v := range counter {
		newCounter[k] = v
	}
	for k := range counter {
		v, ok := s.react[k]
		if ok {
			parts := strings.Split(k, "")
			left, right := parts[0], parts[1]
			newCounter[left+v]++
			newCounter[v+right]++
		}
	}
	return newCounter
}

func (s *Solver) Part2() int {
	return 0
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

func parseLines(lines []string) (string, map[string]string) {
	reaction := ""
	isFormula := false
	react := make(map[string]string)
	for _, line := range lines {
		if len(line) == 0 {
			isFormula = true
			continue
		}
		if isFormula {
			parts := strings.Split(line, " -> ")
			from, to := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
			react[from] = to
		} else {
			reaction = line
		}
	}
	return reaction, react
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
var input = `KFFNFNNBCNOBCNPFVKCP

PB -> F
KC -> F
OB -> H
HV -> N
FS -> S
CK -> K
CC -> V
HF -> K
VP -> C
CP -> S
HO -> N
OS -> N
HS -> O
HB -> F
OH -> V
PP -> B
BS -> N
VS -> F
CN -> B
KB -> O
KH -> B
SS -> K
NS -> B
BP -> V
FB -> S
PV -> O
NB -> S
FC -> F
VB -> P
PC -> O
VF -> K
BV -> K
OO -> B
PN -> N
NH -> H
SP -> B
KF -> O
BN -> F
OF -> C
VV -> H
BB -> P
KN -> H
PO -> C
BH -> O
HC -> B
VO -> O
FV -> B
PK -> V
KO -> H
BK -> V
SC -> S
KV -> B
OV -> S
HK -> F
NP -> V
VH -> P
OK -> S
SO -> C
PF -> C
SH -> N
FP -> V
CS -> C
HH -> O
KK -> P
BF -> S
NN -> O
OC -> C
CB -> O
BO -> V
ON -> F
BC -> P
NO -> N
KS -> H
FF -> V
FN -> V
HP -> N
VC -> F
OP -> K
VN -> S
NV -> F
SV -> F
FO -> V
PS -> H
VK -> O
PH -> P
NF -> N
KP -> S
CF -> S
FK -> P
FH -> F
CO -> H
SN -> B
NC -> H
SK -> P
CV -> P
CH -> H
HN -> N
SB -> H
NK -> B
SF -> H`
