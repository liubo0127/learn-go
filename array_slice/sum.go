package main

import (
	"fmt"
	"reflect"
)

func ArraySum(numbers [5]int) int {
	sum := 0

	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SliceSum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SliceSumAll(numbers ...[]int) []int {
	r := make([]int, 0)
	for _, s := range numbers {
		r = append(r, SliceSum(s))
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
}
