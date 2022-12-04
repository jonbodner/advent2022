package main

import (
	"advent2022/process"
	"fmt"
)

// A & X- Rock
// B & Y- Paper
// C & Z- Scissors
// Loss - 0
// Draw - 3
// Win -6
func main() {
	// part 1
	options := map[string]int{
		"A X": 4, // Draw (3), choose rock (1)
		"A Y": 8, // Win (6), choose paper (2)
		"A Z": 3, // Loss(0), choose scissor(3)
		"B X": 1, // Loss(0), choose rock (1)
		"B Y": 5, // Draw(3), choose paper (2)
		"B Z": 9, // Win(6), choose scissor (3)
		"C X": 7, // Win(6), choose rock (1)
		"C Y": 2, // Loss(0), choose paper (2)
		"C Z": 6, // Draw(3), choose scissor (3)
	}
	total := 0
	process.ByLine("day2/input.txt", func(line []byte) {
		total += options[string(line)]
	})
	fmt.Println(total)

	// part 2
	options2 := map[string]int{
		"A X": 3, // Lose to rock(0), choose scissor (3)
		"A Y": 4, // Draw to rock (3), choose rock (1)
		"A Z": 8, // Win to rock(6), choose paper(2)
		"B X": 1, // Lose to paper(0), choose rock (3)
		"B Y": 5, // Draw to paper (3), choose paper (2)
		"B Z": 9, // Win to paper (6), choose scissor (3)
		"C X": 2, // Lose to scissor(0), choose paper (2)
		"C Y": 6, // Draw to scissor(3), choose scissor (3)
		"C Z": 7, // Win to scissor(6), choose rock (1)
	}
	total2 := 0
	process.ByLine("day2/input.txt", func(line []byte) {
		total2 += options2[string(line)]
	})
	fmt.Println(total2)
}
