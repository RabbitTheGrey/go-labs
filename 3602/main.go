package main

import "fmt"

var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9}

func main() {
	var buf = make(map[int]int)

	for _, i := range nums {
		buf[i]++

		if buf[i] > 1 {
			fmt.Println("true")
			return
		}
	}

	fmt.Println("false")
}

