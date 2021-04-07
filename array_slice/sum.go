package main

import (
	"fmt"
)

func ArraySum(numbers [5]int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SliceSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SliceSumAll(numbers ...[]int) []int {
	r := make([]int, 0)
	for _, number := range numbers {
		r = append(r, SliceSum(number))
	}
	return r
}

// func to calculate sum for number[1:]
func SliceSumTail(numbers ...[]int) []int {
	var r []int
	for _, number := range numbers {
		switch len(number) {
		case 0:
			r = append(r, 0)
		default:
			r = append(r, SliceSum(number[1:]))
		}
	}
	return r
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	r1 := ArraySum(a)
	r2 := SliceSum(b)
	fmt.Println(r1, r2)

	c := []int{1, 2, 3}
	d := []int{4, 5, 6}
	r3 := SliceSumAll(c, d)
	fmt.Println(r3)

	r4 := SliceSumTail(c, d)
	fmt.Println(r4)

	r5 := SliceSumTail([]int{}, []int{1})
	fmt.Println(r5)
}
