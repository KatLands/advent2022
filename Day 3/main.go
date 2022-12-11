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
)

var total int

func main() {
	//opening the input file
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//using scanner to read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results := scanner.Text()
		index := len(results) / 2
		//splitting each line in half to compare later
		first_half := string(results[0:index])
		second_half := string(results[index:])
		rearrangement(first_half, second_half)

		if err != nil {
			log.Fatal(err)
		}
	}

	//print out sum of priorities
	fmt.Println("Sum:", total)
}

func rearrangement(first_half string, second_half string) {
	priority := make(map[string]int)
	sum := make(map[string]int)
	var count_lower int = 1
	var count_upper int = 27

	//looping through alphabet to create map
	for ch := 'a'; ch <= 'z'; ch++ {
		priority[string(ch)] = count_lower
		//sum[string(ch)] = 0
		count_lower++
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		priority[string(ch)] = count_upper
		//sum[string(ch)] = 0
		count_upper++
	}

	//loop over map keys to search for occurrence in both strings and count occurrences
	for key := range priority {
		if strings.Contains(first_half, key) {
			sum[key]++
		}
		if strings.Contains(second_half, key) {
			sum[key]++
		}
	}

	//pull letter from map that occurs most often and assign to match
	max_value := 0
	var match string
	for key, value := range sum {
		if value > max_value {
			match = key
			max_value = value
		}
	}

	//add all matches up based on priority value
	total += priority[match]

}
