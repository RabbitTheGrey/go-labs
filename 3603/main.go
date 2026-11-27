package main

import (
	"fmt"
	"slices"
)

var firstSlice = []string{"я", "помню", "чудное", "мгновенье"}
var secondSlice = []string{"ананас", "банан", "груша", "яблоко"}

func check(slice []string) {
	cp := slices.Clone(slice)
	slices.Sort(slice)

	fmt.Println(slices.Equal(cp, slice))
}

func main() {
	check(firstSlice)
	check(secondSlice)
}
