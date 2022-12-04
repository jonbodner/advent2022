package main

import (
	"advent2022/process"
	"fmt"
	"strconv"
	"strings"
)

type pair struct {
	start, end int
}

func main() {
	total := 0
	process.ByLine("day4/input.txt", func(line []byte) {
		p1, p2 := calcPairs(string(line))
		if p1.start <= p2.start && p1.end >= p2.end {
			fmt.Println(string(line))
			total++
		} else if p2.start <= p1.start && p2.end >= p1.end {
			fmt.Println(string(line))
			total++
		}
	})
	fmt.Println(total)

	// part 2
	total2 := 0
	process.ByLine("day4/input.txt", func(line []byte) {
		p1, p2 := calcPairs(string(line))
		if p1.start >= p2.start && p2.end >= p1.start {
			total2++
		} else if p2.start >= p1.start && p1.end >= p2.start {
			total2++
		}
	})
	fmt.Println(total2)
}

func calcPairs(line string) (pair, pair) {
	p1, p2, _ := strings.Cut(string(line), ",")
	start1, end1, _ := strings.Cut(p1, "-")
	start2, end2, _ := strings.Cut(p2, "-")
	start1i, _ := strconv.Atoi(start1)
	start2i, _ := strconv.Atoi(start2)
	end1i, _ := strconv.Atoi(end1)
	end2i, _ := strconv.Atoi(end2)
	return pair{
			start: start1i,
			end:   end1i,
		}, pair{
			start: start2i,
			end:   end2i,
		}
}
