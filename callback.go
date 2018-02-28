package main

import "fmt"

type Callback func(i int) int

func sum(ints []int, callback Callback) int {
	var sum int
	for _, v := range ints {
		sum += v
	}
	return callback(sum)
}

func main() {
	var ints = []int{1, 2, 3, 4, 5}
	a := sum(ints, func(i int) int {
		return i * 2
	})
	fmt.Println(a)
}
