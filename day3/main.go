package main

import (
	"advent2022/process"
	"fmt"
)

func main() {
	total := 0
	process.ByLine("day3/input.txt", func(line []byte) {
		part1 := line[:len(line)/2]
		part2 := line[len(line)/2:]
		m := map[byte]bool{}
		for _, v := range part1 {
			m[v] = true
		}
		for _, v := range part2 {
			if m[v] {
				//fmt.Println("same letter:", string(v))
				total += calcVal(v)
				break
			}
		}
	})
	fmt.Println(total)

	// part 2
	total2 := 0
	process.ByLines("day3/input.txt", 3, func(lines [][]byte) {
		m := map[byte]int{}
		for i := 0; i < 3; i++ {
			for _, v := range lines[i] {
				m[v] |= 1 << i
			}
		}
		//fmt.Println(m)
		for k, v := range m {
			if v == 7 {
				//fmt.Println(string(k))
				total2 += calcVal(k)
				break
			}
		}
	})
	fmt.Println(total2)
}

func calcVal(v byte) int {
	switch {
	case 'a' <= v && v <= 'z':
		val := int(v - 'a' + 1)
		return val
	case 'A' <= v && v <= 'Z':
		val := int(v - 'A' + 27)
		return val
	default:
		panic(string(v))
	}
}
