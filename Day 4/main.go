package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var count int
	//using scanner to read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results := (scanner.Text())
		first, last := splitLine(results)
		firstPairAsIntArray, lastPairAsIntArray := convertFirstLastToInt(first, last)
		if isContained(firstPairAsIntArray, lastPairAsIntArray) {
			count += 1
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	//print needed result
	fmt.Println(count)
}

func splitLine(results string) (first, last string) {
	splitPair := strings.Split(results, ",")
	first, last = splitPair[0], splitPair[1]
	return
}

func convertFirstLastToInt(first, last string) (firstPairAsIntArray []int, lastPairAsIntArray []int) {
	//this splits the first and last string on a given line with a - delimiter
	firstPair := strings.Split(first, "-")
	lastPair := strings.Split(last, "-")

	//now that we have two seperate numbers, we can convert each number into an int
	firstIndexOfFirstPair, _ := strconv.Atoi(firstPair[0])
	lastIndexOfFirstPair, _ := strconv.Atoi(firstPair[1])
	firstIndexOfLastPair, _ := strconv.Atoi(lastPair[0])
	lastIndexOfLastPair, _ := strconv.Atoi(lastPair[1])

	//creating an array with first and last numerical values
	firstPairAsIntArray = []int{firstIndexOfFirstPair, lastIndexOfFirstPair}
	lastPairAsIntArray = []int{firstIndexOfLastPair, lastIndexOfLastPair}

	return

}

//not needed anymore, leaving in case needed for part 2

// func findIntRange(firstPairAsIntArray []int, lastPairAsIntArray []int) (firstRange []int, lastRange []int) {
// 	//finding the range for first and last part of a line.
// 	//for example an array of 2 4 should return an slice of 2, 3, 4
// 	for i := firstPairAsIntArray[0]; i <= firstPairAsIntArray[1]; i++ {
// 		firstRange = append(firstRange, i)
// 	}
// 	for i := lastPairAsIntArray[0]; i <= lastPairAsIntArray[1]; i++ {
// 		lastRange = append(lastRange, i)
// 	}
// 	return

// }

func isContained(firstPair, lastPair []int) (result bool) {
	//checking if values are contained based on the first and last value in each pair.
	if (firstPair[0] <= lastPair[0] && firstPair[1] >= lastPair[1]) || (lastPair[0] <= firstPair[0] && lastPair[1] >= firstPair[1]) {
		result = true
	}
	return
}
