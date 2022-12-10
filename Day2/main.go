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
)

var self_total int

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
		calc_points(results, results_unsplit)
	}

	if err != nil {
		log.Fatal(err)
	}

	//print result needed to for AOC22
	fmt.Println("Total points for self:", self_total)

}

func calc_points(results []string, results_unsplit string) {
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
		self_total += points[results[1]]
		self_total += game_map[results_unsplit]

	}

}
