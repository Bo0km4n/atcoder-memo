package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var input = make([]int64, 0)

var A, B, K int64

func main() {
	for sc.Scan() {
		in := sc.Text()
		ins := strings.Split(in, " ")
		for i := range ins {
			inInt, _ := strconv.ParseInt(ins[i], 10, 64)
			input = append(input, inInt)
		}

		// first
		A = input[0]
		B = input[1]
		K = input[2]
		eat()
		fmt.Println(A, B)
		input = make([]int64, 0)
	}
}

func eat() {

	for i := int64(1); i <= K; i++ {
		switch i % 2 {
		// 青木
		case 0:
			if B%2 == 0 {
				// nop
			} else {
				B = B - 1
			}
			pass := B / 2
			B = B / 2
			A = A + pass
		// 高橋
		case 1:
			if A%2 == 0 {
				// nop
			} else {
				A = A - 1
			}
			pass := A / 2
			A = A / 2
			B = B + pass
		}
	}
}
