package main

import "fmt"

func main() {
	input := 10
	if ispalendrome(input) {
		fmt.Printf("%d is a palendrome!\n", input)
	}
}

func ispalendrome(input int) bool {
	var myslice []int

	if input < 0 {
		return false
	}

	length := 1
	var base int
	for base = 1; input/base >= 10; base *= 10 {
		length++
	}

	for base > 0 {
		curr := input / base // ints so we strip remainder
		input = input - (curr * base)
		base /= 10
		if curr == 0 {
			myslice = append(myslice, 0)
			continue
		}
		myslice = append(myslice, curr)
	}

	split := (length + 2 - 1) / 2
	for i := 0; i < split; i++ {
		if myslice[i] != myslice[length-1-i] {
			return false
		}
	}
	return true
}
