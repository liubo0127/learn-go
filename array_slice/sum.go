package main

import "fmt"

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

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	r1 := ArraySum(a)
	r2 := SliceSum(b)
	fmt.Println(r1, r2)
}
