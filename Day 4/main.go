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

	var containedCount int
	var overlappedCount int
	//using scanner to read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results := (scanner.Text())
		first, last := splitLine(results)
		firstPairAsIntArray, lastPairAsIntArray := convertFirstLastToInt(first, last)
		if isContained(firstPairAsIntArray, lastPairAsIntArray) {
			containedCount += 1
		}

		if isOverlaped(firstPairAsIntArray, lastPairAsIntArray) {
			overlappedCount += 1
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	//print needed result
	fmt.Println("Contained Count:", containedCount)
	fmt.Println("Overlapped Count:", overlappedCount)
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

//part 1
func isContained(firstPair, lastPair []int) (result bool) {
	//checking if values are contained based on the first and last value in each pair.
	if (firstPair[0] <= lastPair[0] && firstPair[1] >= lastPair[1]) || (lastPair[0] <= firstPair[0] && lastPair[1] >= firstPair[1]) {
		result = true
	}
	return
}

//part 2
func isOverlaped(firstPair, lastPair []int) (result bool) {
	//checking if values are overlapping based on the first and last value in each pair
	if (firstPair[0] < lastPair[0] && firstPair[1] < lastPair[0]) || (lastPair[0] < firstPair[0] && lastPair[1] < firstPair[0]) {
		result = false
	} else {
		result = true
	}
	return

}
