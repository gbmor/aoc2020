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
		return []int{
			35,
			20,
			15,
			25,
			47,
			40,
			62,
			55,
			65,
			95,
			102,
			117,
			150,
			182,
			127,
			219,
			299,
			277,
			309,
			576,
		}
	}
	b, _ := ioutil.ReadFile("input.txt")
	s := strings.Split(string(b), "\n")
	d := make([]int, len(s))
	for i, e := range s {
		d[i], _ = strconv.Atoi(e)
	}
	return d
}

func part1(s []int, preambleLength int) int {
	var pre []int
	for i := preambleLength; i < len(s); i++ {
		pre = s[i-preambleLength : i]
		found := false
		for _, j := range pre {
			for _, l := range pre {
				if j != l && j+l == s[i] {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return s[i]
		}
	}
	return -1
}

func part2(s []int, num int) int {
	var set []int
	for i := 0; i < len(s); i++ {
		sum := 0
		for j := i; j < len(s); j++ {
			sum += s[j]
			if sum == num {
				set = s[i : j+1]
				break
			}
			if sum > num {
				break
			}
		}
		if set != nil {
			break
		}
	}
	if set == nil {
		return -1
	}
	sort.Ints(set)
	return set[0] + set[len(set)-1]
}

func main() {
	input := getInput(false)

	t := time.Now()
	out := part1(input, 25)
	fmt.Printf("%d, %s\n", out, time.Since(t))

	t = time.Now()
	out = part2(input, out)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
