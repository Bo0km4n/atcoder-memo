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
var inputLimit = 2

var maxSlime = 100
var colorMax = 10000
var change = 0

type houseSlimes struct {
	slimes []int
	idx    int
	next   int
	prev   int
}

func main() {
	i := 0
	for sc.Scan() {
		i++
		input = append(input, sc.Text())
		if i == inputLimit {
			break
		}
	}

	h := houseSlimes{}
	slimeStrings := splitSpace(input[1])

	for i := range slimeStrings {
		v, _ := strconv.Atoi(slimeStrings[i])
		h.slimes = append(h.slimes, v)
	}

	indexes := h.getChangeIndexes()
	fmt.Println(h.changeColor(indexes))
}

func (h *houseSlimes) getChangeIndexes() []int {
	buf := 0
	queue := make([]int, 0)

	for i := range h.slimes {
		if buf == 0 {
			buf = h.slimes[i]
			continue
		} else {
			if h.slimes[i] == buf {
				queue = append(queue, i)
				buf = 0
			} else {
				buf = h.slimes[i]
			}
		}
	}
	return queue
}

func (h *houseSlimes) changeColor(indexes []int) int {
	return len(indexes)
}

func splitSpace(input string) []string {
	return strings.Split(input, " ")
}
