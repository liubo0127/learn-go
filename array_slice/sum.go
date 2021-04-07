package main

import "fmt"

func Sum(numbers [5]int) int {
	sum := 0

	for _, i := range numbers {
		sum += i
	}
	return sum
}

func main() {
	r := Sum([5]int{1, 2, 3, 4, 5})
	fmt.Println(r)
}
