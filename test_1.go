// Simple Go program to illustrate
// how to create a rune
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var resultInt []int
	var result []bool
	var lastChar rune

	args := os.Args[1:]
	arg := args[0]
	query := args[1]
	queries := strings.Split(query, ",")

	fmt.Println(fmt.Sprintf("String : %s", arg))
	fmt.Println(fmt.Sprintf("Queries : %+v", queries))
	fmt.Println("Make sure queries is the correct answer")

	for _, char := range strings.ToLower(arg) {
		if lastChar == char {
			resultInt[len(resultInt)-1] += int(char) - 96
		} else {
			resultInt = append(resultInt, int(char)-96)
		}

		lastChar = char
	}

	for idx, item := range queries {
		result = append(result, item == fmt.Sprintf("%d", resultInt[idx]))
	}

	if len(queries) != len(resultInt) {
		fmt.Println("are you sure queries is correct ?")
	}

	fmt.Println(fmt.Sprintf("Result Queries : %+v", resultInt))
	fmt.Println(result)
}
