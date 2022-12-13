package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	InputFile = "input"
)

func main() {
	//opening the input file
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//using scanner to read file line by line
	scanner := bufio.NewScanner(file)
	var resultsArr []string
	var totalP1 int
	priority := createPriorityMap()

	for scanner.Scan() {

		//part 1
		results := scanner.Text()
		firstHalf, secondHalf := splitString(results)
		totalP1 += rearrangementP1(firstHalf, secondHalf, priority)

		//part 2
		resultsArr = append(resultsArr, scanner.Text())

		if err != nil {
			log.Fatal(err)
		}
	}

	total_p2 := rearrangementP2(resultsArr, createPriorityMap())

	//print out sum of priorities p1
	fmt.Println("Sum:", totalP1)

	//print out sum of priorities p2
	fmt.Println("Sum:", total_p2)
}

func splitString(results string) (firstHalf, secondHalf string) {
	index := len(results) / 2
	//splitting each line in half to compare later
	firstHalf = string(results[:index])
	secondHalf = string(results[index:])
	return
}

func createPriorityMap() map[string]int {
	priority := make(map[string]int)
	var countLower int = 1
	var countUpper int = 27

	//looping through alphabet to create map
	for ch := 'a'; ch <= 'z'; ch++ {
		priority[string(ch)] = countLower
		countLower++
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		priority[string(ch)] = countUpper
		countUpper++
	}
	return priority
}

func rearrangementP1(firstHalf string, secondHalf string, priority map[string]int) int {
	sum := make(map[string]int)
	var total int
	//loop over map keys to search for occurrence in both strings and count occurrences
	for key := range priority {
		if strings.Contains(firstHalf, key) {
			sum[key]++
		}
		if strings.Contains(secondHalf, key) {
			sum[key]++
		}
	}

	//pull letter from map that occurs most often and assign to match
	maxValue := 0
	var match string
	for key, value := range sum {
		if value > maxValue {
			match = key
			maxValue = value
		}
	}

	//add all matches up based on priority value
	total = priority[match]
	return total
}

func rearrangementP2(resultsArr []string, priority map[string]int) int {

	var match string
	var total int
	//since we need to to group the elves, loop increments by 3
	for i := 0; i < len(resultsArr)-1; i += 3 {
		//loop over map keys to search for occurrence of a similar letter in all 3 groups.
		for key := range priority {
			if strings.Contains(resultsArr[i], key) && strings.Contains(resultsArr[i+1], key) && strings.Contains(resultsArr[i+2], key) {
				match = key
			}
		}

		//add all matches up based on priority value
		total += priority[match]
	}
	return total
}
