package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

func getInput(test bool) []int {
	if test {
		out := []int{
			28,
			33,
			18,
			42,
			31,
			14,
			46,
			20,
			48,
			47,
			24,
			23,
			49,
			45,
			19,
			38,
			39,
			11,
			1,
			32,
			25,
			35,
			8,
			17,
			7,
			9,
			4,
			2,
			34,
			10,
			3,
		}
		sort.Ints(out)
		return out
	}
	b, _ := ioutil.ReadFile("input.txt")
	bs := strings.Split(string(b), "\n")
	ints := make([]int, len(bs))
	for i, e := range bs {
		ints[i], _ = strconv.Atoi(e)
	}
	sort.Ints(ints)
	return ints
}

func part1(s []int) int {
	ones := 1
	threes := 1
	for i := 1; i < len(s); i++ {
		switch s[i] - s[i-1] {
		case 1:
			ones++
		case 3:
			threes++
		}
	}
	return ones * threes
}

func getChains(hist map[int]int, s []int, n int) int {
	if n == s[len(s)-1] {
		return 1
	}
	if v, ok := hist[n]; ok {
		return v
	}
	count := 0
	for _, e := range s {
		switch e {
		case n + 1:
			count += getChains(hist, s, n+1)
		case n + 2:
			count += getChains(hist, s, n+2)
		case n + 3:
			count += getChains(hist, s, n+3)
		}
	}
	hist[n] = count
	return count
}

func main() {
	input := getInput(false)

	t := time.Now()
	out := part1(input)
	fmt.Printf("%d, %s\n", out, time.Since(t))

	t = time.Now()
	out = getChains(make(map[int]int), input, 0)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
