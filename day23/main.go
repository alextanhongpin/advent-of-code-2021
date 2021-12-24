// You can edit this code!
// Click here and start typing.
package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(solve(example))  // 12521
	fmt.Println(solve(input))    // 18282
	fmt.Println(solve(example2)) // 44169
	fmt.Println(solve(input2))
}

var adjacents = []point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func solve(input string) int {
	maze := parse(input)
	hallways, paths, amps := maze.hallways, maze.paths, maze.amps
	entrances := getEntrances(paths, len(amps))
	rooms := getRooms(amps)
	fmt.Println("entrances:", entrances)
	fmt.Println("amps:", amps)
	fmt.Println("rooms:", rooms)
	fmt.Println("hallways:", hallways)

	pq := make(PriorityQueue, 1, 1e6)
	pq[0] = &move{
		amps: amps,
	}
	heap.Init(&pq)

	draw := func(amps amphipods) {
		maxX, maxY := 0, 0
		for p := range paths {
			if p.x > maxX {
				maxX = p.x
			}
			if p.y > maxY {
				maxY = p.y
			}
		}
		ampByPoint := make(map[point]string)
		for _, amp := range amps {
			ampByPoint[amp.point] = amp.label
		}
		for y := 0; y <= maxY; y++ {
			for x := 0; x <= maxX; x++ {
				p := point{x, y}
				if a, ok := ampByPoint[p]; ok {
					fmt.Print(a)
				} else if paths[p] {
					fmt.Print(".")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}

	isInRoom := func(p point) bool {
		_, ok := rooms[p]
		return ok
	}

	getRoomsForAmphipod := func(amp amphipod) []point {
		var ampRooms []point
		for p, label := range rooms {
			if label == amp.label {
				ampRooms = append(ampRooms, p)
			}
		}
		sort.Slice(ampRooms, func(i, j int) bool {
			return ampRooms[i].y < ampRooms[j].y
		})
		return ampRooms
	}

	isInDesignatedRoom := func(amp amphipod) bool {
		a, ok := rooms[amp.point]
		return ok && a == amp.label
	}

	isInRoomAndBottom := func(amp amphipod, ams amphipods) bool {
		if !isInDesignatedRoom(amp) {
			return false
		}
		m := ams.ToMap()

		rooms := getRoomsForAmphipod(amp)
		for i, room := range rooms {
			if room == amp.point {
				bottom := rooms[i+1:]
				for _, btm := range bottom {
					if m[btm] != amp.label {
						return false
					}
				}
				return true
			}
		}
		return false
	}

	isDone := func(amps amphipods) bool {
		for _, amp := range amps {
			if !isInDesignatedRoom(amp) {
				return false
			}
		}
		return true
	}

	isDesignatedRoomEmptyAndHasNoOtherAmphipods := func(amp amphipod, ams amphipods) bool {
		m := ams.ToMap()
		rooms := getRoomsForAmphipod(amp)
		for _, room := range rooms {
			empty := m[room] == ""
			match := m[room] == amp.label
			valid := empty || match
			if !valid {
				return false
			}
		}
		return true
	}

	cache := make(map[string]bool)

	iter := 0
	for pq.Len() > 0 {
		iter++
		m := heap.Pop(&pq).(*move)
		if isDone(m.amps) {
			return m.priority
		}
		if iter%1e3 == 0 {
			fmt.Println("iter", iter, "energy:", m.priority)
			draw(m.amps)
		}

		if cache[m.amps.String()] {
			continue
		}
		cache[m.amps.String()] = true

		obstacles := make(map[point]bool)
		for _, amp := range m.amps {
			obstacles[amp.point] = true
		}
		for i, amp := range m.amps {
			if isInRoom(amp.point) {
				//fmt.Println("in room", amp.point)
				if isInDesignatedRoom(amp) && isInRoomAndBottom(amp, m.amps) {
					//fmt.Println("in designated room and bottom", amp.point)
					continue
				}
				// Move to hallway.
				for hallway := range hallways {
					// Skip entrances.
					if entrances[hallway] {
						continue
					}
					steps := bfs(amp.point, hallway, paths, obstacles)
					if steps > 0 {
						//fmt.Println("move to hallway", hallway, steps)
						amp2 := amp
						amp2.point = hallway
						m2 := m.Clone()
						m2.amps[i] = amp2
						m2.priority += amp.Energy() * steps
						heap.Push(&pq, &m2)
					}
				}
			} else {
				// Is in hallway.
				if isDesignatedRoomEmptyAndHasNoOtherAmphipods(amp, m.amps) {
					// Move to room.
					for _, room := range getRoomsForAmphipod(amp) {
						steps := bfs(amp.point, room, paths, obstacles)
						if steps > 0 {
							amp2 := amp
							amp2.point = room
							m2 := m.Clone()
							m2.amps[i] = amp2
							m2.priority += amp.Energy() * steps
							heap.Push(&pq, &m2)
						}
					}
				}
			}
		}
	}
	return 0
}

func bfs(from, to point, paths, obstacles map[point]bool) int {
	type move struct {
		point point
		steps int
	}

	cache := make(map[point]bool)
	stack := []move{move{from, 0}}
	for len(stack) > 0 {
		var m move
		m, stack = stack[0], stack[1:]
		if m.point == to {
			return m.steps
		}
		if cache[m.point] {
			continue
		}
		cache[m.point] = true

		for _, adj := range adjacents {
			m2 := m
			m2.point.x += adj.x
			m2.point.y += adj.y
			m2.steps++
			if !paths[m2.point] || obstacles[m2.point] {
				continue
			}
			stack = append(stack, m2)
		}
	}
	return 0
}

type move struct {
	amps     amphipods
	priority int
}

func (m move) Clone() move {
	return move{
		amps:     m.amps.Clone(),
		priority: m.priority,
	}
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*move

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	move := x.(*move)
	*pq = append(*pq, move)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func getRooms(amps amphipods) map[point]string {
	sort.Slice(amps, func(i, j int) bool {
		if amps[i].point.x != amps[j].point.x {
			return amps[i].point.x < amps[j].point.x
		}
		return amps[i].point.y < amps[j].point.y
	})

	res := make(map[point]string)
	rooms := strings.Fields("A B C D")
	n := len(amps) / 4
	for i, amp := range amps {
		res[amp.point] = rooms[i/n]
	}
	return res
}

func getEntrances(maze map[point]bool, n int) map[point]bool {
	var paths []point
	for p := range maze {
		paths = append(paths, p)
	}
	sort.Slice(paths, func(i, j int) bool {
		if paths[i].y != paths[j].y {
			return paths[i].y > paths[j].y
		}
		return paths[i].x < paths[j].x
	})

	res := make(map[point]bool)
	for _, p := range paths[:4] {
		p.y -= n / 4
		res[p] = true
	}
	return res
}

type point struct{ x, y int }

func (p point) String() string {
	return fmt.Sprintf("<%d, %d>", p.x, p.y)
}

type amphipod struct {
	point point
	label string
}

func (a amphipod) Energy() int {
	switch a.label {
	case "A":
		return 1
	case "B":
		return 10
	case "C":
		return 100
	case "D":
		return 1000
	}
	panic("invalid label")
}

func (a amphipod) String() string { return fmt.Sprintf("%s:%s", a.label, a.point) }

type amphipods []amphipod

func (a amphipods) ToMap() map[point]string {
	res := make(map[point]string)
	for _, amp := range a {
		res[amp.point] = amp.label
	}
	return res
}

func (a amphipods) Clone() amphipods {
	res := make(amphipods, len(a))
	copy(res, a)
	return res
}

func (a amphipods) String() string {
	res := make([]string, len(a))
	for i, v := range a {
		res[i] = v.String()
	}
	sort.Strings(res)
	return strings.Join(res, ",")
}

type maze struct {
	hallways map[point]bool
	paths    map[point]bool
	amps     amphipods
}

func parse(input string) maze {
	hallways := make(map[point]bool)
	paths := make(map[point]bool)
	amps := make(amphipods, 0)

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, c := range line {
			p := point{x, y}
			switch c {
			case '.':
				hallways[p] = true
				paths[p] = true
			case 'A', 'B', 'C', 'D':
				amps = append(amps, amphipod{p, string(c)})
				paths[p] = true
			}
		}

	}
	return maze{
		paths:    paths,
		hallways: hallways,
		amps:     amps,
	}
}

var example = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

var input = `#############
#...........#
###C#C#A#B###
  #D#D#B#A#
  #########`

var example2 = `#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########`

var input2 = `#############
#...........#
###C#C#A#B###
  #D#C#B#A#
  #D#B#A#C#
  #D#D#B#A#
  #########`
