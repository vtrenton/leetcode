package main

import (
	"strings"
)

func main() {
	testarr := []string{"flower", "flow", "flight"}
	getprefix(testarr)
}

func getprefix(array []string) string {
	var sb strings.Builder
	for _, word := range array {
		for _, v := range word {
			//what i dooo????????????
		}
	}
}
