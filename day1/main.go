package main

import (
	"advent2022/process"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	type data struct {
		id    int
		total int
	}
	curId := 0
	var vals []data

	curData := data{
		id: curId,
	}
	process.ByLine("day1/input.txt", func(line []byte) {
		if len(line) == 0 {
			vals = append(vals, curData)
			curId++
			curData = data{
				id: curId,
			}
		} else {
			val, _ := strconv.Atoi(string(line))
			curData.total += val
		}

	})
	vals = append(vals, curData)
	fmt.Println(vals)
	sort.Slice(vals, func(i, j int) bool {
		return vals[i].total > vals[j].total
	})
	// part 1
	fmt.Println(vals[0])

	// part 2
	fmt.Println(vals[0].total + vals[1].total + vals[2].total)
}
