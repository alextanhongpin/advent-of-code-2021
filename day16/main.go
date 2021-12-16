// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	// Part 1.
	example  = `38006F45291200`
	example2 = `EE00D40C823060`
	example3 = `8A004A801A8002F478`
	example4 = `620080001611562C8802118E34`
	example5 = `C0015000016115A2E0802F182340`
	example6 = `A0016C880162017C3686B18A3D4780`

	// Part 2.
	example7  = `C200B40A82`
	example8  = `04005AC33890`
	example9  = `880086C3E88112`
	example10 = `CE00C43D881120`
	example11 = `D8005AC2A8F0`
	example12 = `F600BC2D8F`
	example13 = `9C005AC2F8F0`
	example14 = `9C0141080250320F1802104A08`

	input = `005173980232D7F50C740109F3B9F3F0005425D36565F202012CAC0170004262EC658B0200FC3A8AB0EA5FF331201507003710004262243F8F600086C378B7152529CB4981400B202D04C00C0028048095070038C00B50028C00C50030805D3700240049210021C00810038400A400688C00C3003E605A4A19A62D3E741480261B00464C9E6A5DF3A455999C2430E0054FCBE7260084F4B37B2D60034325DE114B66A3A4012E4FFC62801069839983820061A60EE7526781E513C8050D00042E34C24898000844608F70E840198DD152262801D382460164D9BCE14CC20C179F17200812785261CE484E5D85801A59FDA64976DB504008665EB65E97C52DCAA82803B1264604D342040109E802B09E13CBC22B040154CBE53F8015796D8A4B6C50C01787B800974B413A5990400B8CA6008CE22D003992F9A2BCD421F2C9CA889802506B40159FEE0065C8A6FCF66004C695008E6F7D1693BDAEAD2993A9FEE790B62872001F54A0AC7F9B2C959535EFD4426E98CC864801029F0D935B3005E64CA8012F9AD9ACB84CC67BDBF7DF4A70086739D648BF396BFF603377389587C62211006470B68021895FCFBC249BCDF2C8200C1803D1F21DC273007E3A4148CA4008746F8630D840219B9B7C9DFFD2C9A8478CD3F9A4974401A99D65BA0BC716007FA7BFE8B6C933C8BD4A139005B1E00AC9760A73BA229A87520C017E007C679824EDC95B732C9FB04B007873BCCC94E789A18C8E399841627F6CF3C50A0174A6676199ABDA5F4F92E752E63C911ACC01793A6FB2B84D0020526FD26F6402334F935802200087C3D8DD0E0401A8CF0A23A100A0B294CCF671E00A0002110823D4231007A0D4198EC40181E802924D3272BE70BD3D4C8A100A613B6AFB7481668024200D4188C108C401D89716A080`
)

func main() {
	s := NewSolver()
	//fmt.Println("part1 (example):", s.WithInput(example).Part1())  // 0
	//fmt.Println("part1 (example):", s.WithInput(example2).Part1()) // 0
	fmt.Println("part1 (example):", s.WithInput(example3).Part1()) // 16
	fmt.Println("part1 (example):", s.WithInput(example4).Part1()) // 12
	fmt.Println("part1 (example):", s.WithInput(example5).Part1()) // 23
	fmt.Println("part1 (example):", s.WithInput(example6).Part1()) // 31
	fmt.Println("part1 (input):", s.WithInput(input).Part1())      // 886

	fmt.Println("part2 (example):", s.WithInput(example7).Part2())  // 3
	fmt.Println("part2 (example):", s.WithInput(example8).Part2())  // 54
	fmt.Println("part2 (example):", s.WithInput(example9).Part2())  // 7
	fmt.Println("part2 (example):", s.WithInput(example10).Part2()) // 9
	fmt.Println("part2 (example):", s.WithInput(example11).Part2()) // 1
	fmt.Println("part2 (example):", s.WithInput(example12).Part2()) // 0
	fmt.Println("part2 (example):", s.WithInput(example13).Part2()) // 0
	fmt.Println("part2 (example):", s.WithInput(example14).Part2()) // 1
	fmt.Println("part2 (input):", s.WithInput(input).Part2())       // 184487454837
}

type Solver struct {
	input  string
	packet *packet
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) WithInput(input string) *Solver {
	s.input = hex2Bin(input)
	s.packet, _ = parse(s.input)
	return s
}

type packet struct {
	version      int
	typeID       int
	lengthTypeID *int
	subpackets   []packet
	literal      *int
}

func parse(b string) (*packet, int) {
	version := bin2Dec(b[:3])
	typeID := bin2Dec(b[3:6])
	var subpackets []packet
	end := len(b)
	var literal, lengthTypeID *int
	if typeID == 4 {
		bb := ""
		for i := 6; i < len(b); i += 5 {
			end = i + 5
			bb += b[i+1 : i+5]
			if b[i] == '0' {
				break
			}
		}
		l := bin2Dec(bb)
		literal = &l
	} else {
		s := 7
		r := toInt(b[6:7])
		lengthTypeID = &r
		switch r {
		case 0:
			e := s + 15
			totalBits := bin2Dec(b[s:e])
			maxLen := e + totalBits
			for e < maxLen {
				sub, end := parse(b[e:])
				if sub != nil {
					subpackets = append(subpackets, *sub)
				}
				if end == 0 {
					break
				}
				e += end
			}
			end = e
		case 1:
			e := s + 11
			nsubs := bin2Dec(b[s:e])
			for nsubs > 0 {
				sub, end := parse(b[e:])
				if sub != nil {
					subpackets = append(subpackets, *sub)
					nsubs--
				}
				if end == 0 {
					break
				}
				e += end
			}
			end = e
		}
	}

	p := &packet{
		version:      version,
		typeID:       typeID,
		literal:      literal,
		lengthTypeID: lengthTypeID,
		subpackets:   subpackets,
	}

	return p, end
}

func (s *Solver) Part1() int {
	version := 0
	packets := []packet{*s.packet}
	for len(packets) > 0 {
		var p packet
		p, packets = packets[0], packets[1:]
		version += p.version
		packets = append(packets, p.subpackets...)
	}
	return version
}

func (s *Solver) Part2() int {
	return expr(s.packet)
}

func expr(p *packet) int {
	startValue := map[int]int{
		0: 0,
		1: 1,
		2: math.MaxInt,
		3: math.MinInt,
	}

	res := startValue[p.typeID]
	packets := []packet{*p}
	for len(packets) > 0 {
		var p packet
		p, packets = packets[0], packets[1:]
		switch p.typeID {
		case 0:
			for _, sub := range p.subpackets {
				res += expr(&sub)
			}
		case 1:
			for _, sub := range p.subpackets {
				res *= expr(&sub)
			}
		case 2:
			for _, sub := range p.subpackets {
				lit := expr(&sub)
				if lit < res {
					res = lit
				}
			}
		case 3:
			for _, sub := range p.subpackets {
				lit := expr(&sub)
				if lit > res {
					res = lit
				}
			}
		case 4:
			res = *p.literal
		case 5:
			if len(p.subpackets) != 2 {
				panic("invalid subpacket for greater than")
			}
			l, r := expr(&p.subpackets[0]), expr(&p.subpackets[1])
			if l > r {
				res = 1
			} else {
				res = 0
			}
		case 6:
			if len(p.subpackets) != 2 {
				panic("invalid subpacket for less than")
			}
			l, r := expr(&p.subpackets[0]), expr(&p.subpackets[1])
			if l < r {
				res = 1
			} else {
				res = 0
			}
		case 7:
			if len(p.subpackets) != 2 {
				panic("invalid subpacket for equal than")
			}
			l, r := expr(&p.subpackets[0]), expr(&p.subpackets[1])
			if l == r {
				res = 1
			} else {
				res = 0
			}
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

func hex2Bin(h string) string {
	res := ""
	for _, n := range h {
		b, err := strconv.ParseInt(string(n), 16, 64)
		if err != nil {
			panic(err)
		}
		res += fmt.Sprintf("%04b", b)
	}
	return res
}

func bin2Dec(b string) int {
	v, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic(err)
	}
	n, err := strconv.Atoi(fmt.Sprintf("%d", v))
	if err != nil {
		panic(err)
	}
	return n
}
