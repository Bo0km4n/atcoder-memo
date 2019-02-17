package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var input = make([]string, 0)
var initA = 0
var inputLimit = 1

func main() {
	i := 0
	for sc.Scan() {
		i++
		input = append(input, sc.Text())

		l, _ := strconv.Atoi(input[0])
		inputLimit = l + 1
		if i == inputLimit {
			break
		}
	}

	in := input[1:]
	for i := range in {
		a, b, c, d := splitSpace(in[i])
		initA = a
		days(a, b, c, d)
	}
}

// func oneDay(a, b, c, d int) {
// 	var nextA int
// 	if a-b <= c {
// 		nextA = a - b + d
// 	} else {
// 		nextA = a - b
// 	}

// 	if nextA < b {
// 		fmt.Println("No")
// 	} else if nextA == initA {
// 		fmt.Println("Yes")
// 	} else {
// 		oneDay(nextA, b, c, d)
// 	}
// }

func days(a, b, c, d int) {
	var nextA int
	for {

		if c+d == a {
			fmt.Println("Yes")
			return
		} else if c+d == a-1 {
			fmt.Println("No")
			return
		}

		if (a - b) < 0 {
			fmt.Println("No")
			return
		}

		if a-b <= c {
			nextA = a - b + d
		} else {
			if a-b == c {
				fmt.Println("Yes")
				return
			}
			nextA = a - b
		}

		if nextA < b {
			fmt.Println("No")
			break
		} else if nextA == initA {
			fmt.Println("Yes")
			break
		} else {
			a = nextA
		}
	}
}

func splitSpace(input string) (a, b, c, d int) {
	s := strings.Split(input, " ")
	a, _ = strconv.Atoi(s[0])
	b, _ = strconv.Atoi(s[1])
	c, _ = strconv.Atoi(s[2])
	d, _ = strconv.Atoi(s[3])
	return a, b, c, d
}
