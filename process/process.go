package process

import (
	"bufio"
	"os"
)

func ByLine(filename string, f func(b []byte)) {
	file, _ := os.Open(filename)
	defer file.Close()
	s := bufio.NewScanner(file)
	for s.Scan() {
		f(s.Bytes())
	}
}

func ByLines(filename string, count int, f func(b [][]byte)) {
	file, _ := os.Open(filename)
	defer file.Close()
	s := bufio.NewScanner(file)
	for {
		var lines [][]byte
		var end bool
		for i := 0; i < count; i++ {
			end = s.Scan()
			lines = append(lines, s.Bytes())
		}
		f(lines)
		if !end {
			break
		}
	}
}
