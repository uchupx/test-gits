package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type arrString []string

func loopCheckPalindrome(j int, i int, arrStr []string) string {
	if j <= i {
		return "0"
	}

	if arrStr[i] == arrStr[j] {
		i++
		j--
	} else {
		return "-1"
	}

	loopCheckPalindrome(j, i, arrStr)
	return "0"
}

func checkPalindrome(str string) string {
	i := 0
	j := len(str) - 1
	arrStr := strings.Split(str, "")

	if loopCheckPalindrome(j, i, arrStr) == "-1" {
		return "-1"
	}

	return str
}

func loopDiff(arrStr []string, j *int, i *int, diff *int) {
	if *i >= len(arrStr)/2 {
		return
	}

	if arrStr[*i] != arrStr[*j] {
		*diff++
	}

	*j--
	*i++

	loopDiff(arrStr, j, i, diff)
}

func loopDoSubString(arrStr *arrString, lo *int, hi *int, k *int, diff *int) {
	if *k <= 0 || *hi <= *lo {
		return
	}

	if (*arrStr)[*lo] == (*arrStr)[*hi] {
		if *k > 1 && (*k-2) >= *diff && (*arrStr)[*lo] != "9" {
			(*arrStr)[*lo] = "9"
			(*arrStr)[*hi] = "9"
			*k -= 2
		}
	} else {
		if *k > 1 && (*k-2) >= *diff-1 {
			if (*arrStr)[*lo] != "9" {
				(*arrStr)[*lo] = "9"
				*k--
			}

			if (*arrStr)[*hi] != "9" {
				(*arrStr)[*hi] = "9"
				*k--
			}
		} else {
			if stringToInt((*arrStr)[*lo]) > stringToInt((*arrStr)[*hi]) {
				(*arrStr)[*hi] = (*arrStr)[*lo]
			} else {
				(*arrStr)[*lo] = (*arrStr)[*hi]
			}
			*k--
		}
		*diff--
	}

	if *lo == *hi && *k > 0 {
		(*arrStr)[*lo] = "9"
		*k--
	}

	*lo++
	*hi--

	fmt.Println(arrStr)
	loopDoSubString(arrStr, lo, hi, k, diff)
}

func execute(str string, k int) string {
	lo := 0
	hi := len(str) - 1
	arrStr := strings.Split(str, "")
	diff := 0

	j := hi
	i := 0

	loopDiff(arrStr, &j, &i, &diff)

	if diff > k {
		return "-1"
	}

	tempArr := arrString(arrStr)
	loopDoSubString(&tempArr, &lo, &hi, &k, &diff)

	return strings.Join(tempArr, "")
}

func stringToInt(val string) int {
	res, _ := strconv.Atoi(val)

	return res
}

func main() {
	args := os.Args[1:]
	arg := args[0]
	k, _ := strconv.Atoi(args[1])

	fmt.Println(fmt.Sprintf("String : %s", arg))
	fmt.Println(fmt.Sprintf("K : %+v", k))

	fmt.Println(checkPalindrome(execute(arg, k)))
}
