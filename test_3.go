package main

import (
	"fmt"
	"os"
)

var pairs = map[rune]rune{'{': '}', '[': ']', '(': ')'}

func main() {
	args := os.Args[1:]
	arg := args[0]

	var queue []rune
	for _, v := range arg {
		switch getBracketType(v) {
		case 1:
			queue = append(queue, pairs[v])
		case 0:
			if 0 < len(queue) && queue[len(queue)-1] == v {
				queue = queue[:len(queue)-1]
			} else {
				fmt.Println("No")
				break
			}
		}
	}

	if len(queue) > 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
	return
}

func getBracketType(char rune) int {
	for k, v := range pairs {
		switch char {
		case k:
			return 1
		case v:
			return 0
		}
	}
	return -1
}
