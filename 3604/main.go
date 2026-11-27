package main

import "fmt"

const str string = "съешь еще этих мягких французских булок да выпей чаю"

func main() {
	letters := make(map[rune]uint8)

	for _, r := range []rune(str) {
		if r == ' ' {
			continue
		}
		letters[r]++
	}

	for letter, count := range letters {
		fmt.Printf("%c: %d\n", letter, count)
	}
}
