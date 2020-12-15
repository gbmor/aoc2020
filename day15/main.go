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

func exec(s []int, breaker int) int {
	cache := make(map[int]int)
	for i, e := range s {
		cache[e] = i
	}
	current := 0
	for i := len(s); i < breaker-1; i++ {
		if _, ok := cache[current]; !ok {
			cache[current] = i
			current = 0
			continue
		}
		pos := cache[current]
		diff := i - pos
		cache[current] = i
		current = diff
	}
	return current
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
