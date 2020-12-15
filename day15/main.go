package main

import (
	"fmt"
	"time"
)

func getInput(test int) []int {
	if test == 1 {
		return []int{0, 3, 6}
	}
	return []int{5, 1, 9, 18, 13, 8, 0}
}

type Pos struct {
	last, prev int
}

func exec(s []int, breaker int) int {
	d := make([]int, breaker)
	copy(d, s)
	cache := make(map[int]int)
	for i, e := range s {
		cache[e] = i
	}
	for i := len(s); i < breaker-1; i++ {
		if _, ok := cache[d[i]]; !ok {
			d[i+1] = 0
			cache[d[i]] = i
			continue
		}
		pos := cache[d[i]]
		diff := i - pos
		cache[d[i]] = i
		d[i+1] = diff
	}
	return d[len(d)-1]
}

func main() {
	i := getInput(0)

	t := time.Now()
	out := exec(i, 2020)
	fmt.Printf("%d, %s\n", out, time.Since(t))

	t = time.Now()
	out = exec(i, 30000000)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
