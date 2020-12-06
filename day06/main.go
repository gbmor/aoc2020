package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func getInput() []string {
	b, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(b), "\n\n")
}

func calc(s []string) (int, int) {
	sum := 0
	p2sum := 0
	for _, v := range s {
		qs := make(map[rune]int)
		g := strings.Split(v, "\n")
		for _, e := range g {
			for _, c := range e {
				qs[c]++
			}
		}
		sum += len(qs)
		for _, e := range qs {
			if e == len(g) {
				p2sum++
			}
		}
	}
	return sum, p2sum
}

func main() {
	input := getInput()
	t := time.Now()
	p1, p2 := calc(input)
	fmt.Printf("P1: %d, P2: %d, %s\n", p1, p2, time.Since(t))
}
