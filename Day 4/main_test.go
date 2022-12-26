package main

import (
	"reflect"
	"testing"
)

func TestSplitLineAtComma(t *testing.T) {
	first, last := splitLine("2-4,6-8")
	if first != "2-4" || last != "6-8" {
		t.Fail()
	}
}

func TestConvertFirstLastToInt(t *testing.T) {
	first, last := convertFirstLastToInt("2-8", "3-7")
	if !reflect.DeepEqual(first, []int{2, 8}) || !reflect.DeepEqual(last, []int{3, 7}) {
		t.Fail()
	}
}

func TestIsContained(t *testing.T) {
	result := isContained([]int{2, 8}, []int{3, 7})
	if result != true {
		t.Fail()
	}
}

func TestIsContainedWithDuplicateValueInPair(t *testing.T) {
	result := isContained([]int{6, 6}, []int{4, 6})
	if result != true {
		t.Fail()
	}

}

func TestIsNotContained(t *testing.T) {
	result := isContained([]int{2, 4}, []int{6, 8})
	if result != false {
		t.Fail()
	}

}

func TestIsOverlapped(t *testing.T) {
	result := isOverlaped([]int{2, 8}, []int{3, 7})
	if result != true {
		t.Fail()
	}
}

func TestIsOverlappedWithDuplicateValueInPair(t *testing.T) {
	result := isOverlaped([]int{6, 6}, []int{4, 6})
	if result != true {
		t.Fail()
	}
}

func TestIsNotOverlapped(t *testing.T) {
	result := isOverlaped([]int{2, 4}, []int{6, 8})
	if result != false {
		t.Fail()
	}

}
