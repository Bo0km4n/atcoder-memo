package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var input = make([]string, 0)

func main() {
	for sc.Scan() {
		in := sc.Text()
		switch len(in) {
		case 2:
			fmt.Println(in)
		case 3:
			fmt.Println(reverse(in))
		}
	}

}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
