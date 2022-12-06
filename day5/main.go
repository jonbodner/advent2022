package main

import (
	"advent2022/process"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	step1()
	step2()
}
func step1() {
	var stacks [9]stack[byte]
	endLine := ` 1   2   3   4   5   6   7   8   9 `
	type state int
	const (
		chart state = iota
		bottom
		blank
		moves
	)
	var curState state
	process.ByLine("day5/input.txt", func(b []byte) {
		line := string(b)
		fmt.Println(line)
		if line == endLine {
			curState = bottom
			for _, v := range stacks {
				v.flip()
			}
			return
		}
		if curState == bottom && len(strings.TrimSpace(line)) == 0 {
			curState = blank
			return
		}
		if curState == blank && strings.HasPrefix(line, "move") {
			curState = moves
		}
		if curState == chart {
			for i := 0; i < 9; i++ {
				ch := line[i*4+1]
				if ch != ' ' {
					stacks[i].push(ch)
				}
			}
		}
		if curState == moves {
			parts := strings.Split(line, " ")
			// part 1 is number to move
			numToMove, _ := strconv.Atoi(parts[1])
			// part 3 is source
			source, _ := strconv.Atoi(parts[3])
			// part 5 is target
			target, _ := strconv.Atoi(parts[5])
			for i := 0; i < numToMove; i++ {
				stacks[target-1].push(stacks[source-1].pop())
			}
		}
	})
	for _, v := range stacks {
		fmt.Print(string(v.peek()))
	}
	fmt.Println()
}

func step2() {
	var stacks [9]stack[byte]
	endLine := ` 1   2   3   4   5   6   7   8   9 `
	type state int
	const (
		chart state = iota
		bottom
		blank
		moves
	)
	var curState state
	process.ByLine("day5/input.txt", func(b []byte) {
		line := string(b)
		fmt.Println(line)
		if line == endLine {
			curState = bottom
			for _, v := range stacks {
				v.flip()
			}
			return
		}
		if curState == bottom && len(strings.TrimSpace(line)) == 0 {
			curState = blank
			return
		}
		if curState == blank && strings.HasPrefix(line, "move") {
			curState = moves
		}
		if curState == chart {
			for i := 0; i < 9; i++ {
				ch := line[i*4+1]
				if ch != ' ' {
					stacks[i].push(ch)
				}
			}
		}
		if curState == moves {
			parts := strings.Split(line, " ")
			// part 1 is number to move
			numToMove, _ := strconv.Atoi(parts[1])
			// part 3 is source
			source, _ := strconv.Atoi(parts[3])
			// part 5 is target
			target, _ := strconv.Atoi(parts[5])
			var ts stack[byte]
			for i := 0; i < numToMove; i++ {
				ts.push(stacks[source-1].pop())
			}
			for !ts.empty() {
				stacks[target-1].push(ts.pop())
			}
		}
	})
	for _, v := range stacks {
		fmt.Print(string(v.peek()))
	}
	fmt.Println()
}

type stack[T any] []T

func (s *stack[T]) push(v T) {
	*s = append(*s, v)
}

func (s *stack[T]) pop() T {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stack[T]) peek() T {
	result := (*s)[len(*s)-1]
	return result
}

func (s *stack[T]) empty() bool {
	return len(*s) == 0
}

func (s *stack[T]) flip() {
	temp := make([]T, 0, len(*s))
	for !s.empty() {
		temp = append(temp, s.pop())
	}
	for _, v := range temp {
		s.push(v)
	}
}
