package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var input = make([]int64, 0)
var inputLimit int64
var big = make([]int64, 0)
var min = make([]int64, 0)

func main() {
	var i int64
	for sc.Scan() {
		in := sc.Text()
		if i == 0 {
			inputLimit, _ = strconv.ParseInt(in, 10, 64)
			i++
			continue
		}

		in64, _ := strconv.ParseInt(in, 10, 64)
		input = append(input, in64)
		if i >= inputLimit {
			break
		}
		i++
	}

	eval()
}

func eval() {

	// find max nums
	for i := int64(0); i < inputLimit/2; i++ {
		var tmp int64
		for _, v := range input {
			if len(big) == 0 && tmp < v {
				tmp = v
			} else if tmp < v && !findBig(v) {
				tmp = v
			}
		}
		big = append(big, tmp)
	}

	// find min nums
	for _, v := range input {
		if !findBig(v) {
			min = append(min, v)
		}
	}

	fmt.Println(big, min)
	// 並び替え
	result := make([]int64, inputLimit)
	for _, v := range big {

	}
}

func findBig(in int64) bool {
	for _, v := range big {
		if v == in {
			return true
		}
	}

	return false
}

func permutation(n int64, k int64) int64 {
	v := int64(1)
	if 0 < k && k <= n {
		for i := int64(0); i < int64(k); i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int64) int64 {
	return permutation(n, n-1)
}

func permutationAndCombinations(array, used []int64) {
	if len(used) == 0 {
		used = append(used, array[0])
		array = array[1:]
		permutationAndCombinations(array, used)
	}

	if len(array) == 0 {
		fmt.Println(used)
		return
	}

	for i := range array {
		if array[i] != used[len(used)-1] {
			used = append(used, array[i])
			array = array[i:]
		}
	}
}
