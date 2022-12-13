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
	var results_arr []string
	for scanner.Scan() {
		//part 2
		results_arr = append(results_arr, scanner.Text())

		if err != nil {
			log.Fatal(err)
		}
	}

	rearrangement(results_arr)

	//print out sum of priorities
	fmt.Println("Sum:", total)
}

func rearrangement(results_arr []string) {
	priority := make(map[string]int)

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

	var match string
	//since we need to to group the elves, loop increments by 3
	for i := 0; i < len(results_arr)-1; i += 3 {
		//loop over map keys to search for occurrence of a similar letter in all 3 groups.
		for key := range priority {
			if strings.Contains(results_arr[i], key) && strings.Contains(results_arr[i+1], key) && strings.Contains(results_arr[i+2], key) {
				match = key
			}
		}

		//add all matches up based on priority value
		total += priority[match]
	}

}
