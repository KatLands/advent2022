package main

import (
	"reflect"
	"testing"
)

func TestSplitLineAtComma(test *testing.T) {
	first, last := splitLine("2-4,6-8")
	if first != "2-4" || last != "6-8" {
		test.Fail()
	}
}

func TestConvertFirstLastToInt(test *testing.T) {
	first, last := convertFirstLastToInt("2-8", "3-7")
	if !reflect.DeepEqual(first, []int{2, 8}) || !reflect.DeepEqual(last, []int{3, 7}) {
		test.Fail()
	}
}

// func TestFindIntRange(test *testing.T) {
// 	first, last := findIntRange([]int{2, 8}, []int{3, 7})
// 	if !reflect.DeepEqual(first, []int{2, 3, 4, 5, 6, 7, 8}) || !reflect.DeepEqual(last, []int{3, 4, 5, 6, 7}) {
// 		test.Fail()
// 	}
// }

func TestIsContained(test *testing.T) {
	result := isContained([]int{2, 8}, []int{3, 7})
	if result != true {
		test.Fail()
	}
}

func TestIsContainedWithDuplicateValueInPair(test *testing.T) {
	result := isContained([]int{6, 6}, []int{4, 6})
	if result != true {
		test.Fail()
	}

}
