package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input int64
	var err error
	if len(os.Args) < 2 {
		fmt.Printf("Please enter a number to test: ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
	} else {
		input, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting string to int: %v\n", err)
			os.Exit(1)
		}
	}

	//fmt.Printf("sanity check: input = %d", input)
	//input := 1001
	if ispalindrome(input) {
		fmt.Printf("%d is a palindrome!\n", input)
	} else {
		fmt.Printf("%d is NOT a palindrome.\n", input)
	}
}

// no need to convert to string :)
func ispalindrome(input int64) bool {
	var myslice []int64

	if input < 0 {
		return false
	}

	length := 1
	var base int64
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

	// no need for ceil
	split := (length + 2 - 1) / 2
	for i := range split {
		if myslice[i] != myslice[length-1-i] {
			return false
		}
	}
	return true
}
