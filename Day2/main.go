package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	INPUT_FILE = "input"
	LOSE       = 0
	DRAW       = 3
	WIN        = 6
	//part 2
	X = LOSE
	Y = DRAW
	Z = WIN
)

var total_part1 int
var total_part2 int

func main() {
	//opening the input file
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//using scanner to read file line by line and split each line where there is a blank space
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results := strings.Split(scanner.Text(), " ")
		results_unsplit := scanner.Text()
		calc_points_part1(results, results_unsplit)
		calc_points_part2(results, results_unsplit)
	}

	if err != nil {
		log.Fatal(err)
	}

	//print result needed to for part1
	fmt.Println("Total points for self part 1:", total_part1)

	//print result needed to for part2
	fmt.Println("Total points for self part 2:", total_part2)

}

func calc_points_part1(results []string, results_unsplit string) {
	points := make(map[string]int)
	points["A"], points["X"] = 1, 1
	points["B"], points["Y"] = 2, 2
	points["C"], points["Z"] = 3, 3

	var game_map = map[string]int{
		"A X": DRAW,
		"A Y": WIN,
		"A Z": LOSE,
		"B X": LOSE,
		"B Y": DRAW,
		"B Z": WIN,
		"C X": WIN,
		"C Y": LOSE,
		"C Z": DRAW,
	}

	//looping through the length of the input to assign points
	for i := 0; i < len(results)-1; i++ {
		total_part1 += points[results[1]]
		total_part1 += game_map[results_unsplit]

	}

}

func calc_points_part2(results []string, results_unsplit string) {
	points := make(map[string]int)
	points["A"], points["X"] = 1, 1
	points["B"], points["Y"] = 2, 2
	points["C"], points["Z"] = 3, 3

	//changing the map to reflect new game rules x loses, y draws, and z wins
	var game_map = map[string]int{
		"A X": LOSE,
		"A Y": DRAW,
		"A Z": WIN,
		"B X": LOSE,
		"B Y": DRAW,
		"B Z": WIN,
		"C X": LOSE,
		"C Y": DRAW,
		"C Z": WIN,
	}

	//key is elf choice value is player choice needed to lose round
	var player_lose = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	//key is elf choice value is player choice needed to draw round
	var player_draw = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	//key is elf choice value is player choice needed to win round
	var player_win = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	for i := 0; i < len(results)-1; i++ {
		if results[1] == "X" {
			//add points based on what I choose to play in order to lose round
			total_part2 += points[player_lose[results[0]]]
			total_part2 += game_map[results_unsplit]
		} else if results[1] == "Y" {
			//add points based on what I choose to play in order to draw round
			total_part2 += points[player_draw[results[0]]]
			total_part2 += game_map[results_unsplit]
		} else if results[1] == "Z" {
			//add points based on what I choose to play in order to win round
			total_part2 += points[player_win[results[0]]]
			total_part2 += game_map[results_unsplit]
		}

	}

}
