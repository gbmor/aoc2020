package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func getInput(test bool) []string {
	if test {
		return []string{
			"1-3 a: abcde",
			"1-3 b: cdefg",
			"2-9 c: ccccccccc",
		}
	}
	b, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(b), "\n")
}

func parse(line string) (int, int, rune, string) {
	var low, high int
	var lets, pass string
	_, err := fmt.Sscanf(line, "%d-%d %s %s", &low, &high, &lets, &pass)
	if err != nil {
		panic(err)
	}
	return low, high, rune(lets[0]), pass
}

func check(input []string) (int, int) {
	part1 := 0
	part2 := 0
	for _, e := range input {
		low, high, letR, password := parse(e)
		letMap := make(map[rune]int)
		for _, c := range password {
			letMap[c]++
		}
		if letMap[letR] >= low && letMap[letR] <= high {
			part1++
		}
		let := uint8(letR)
		if password[low-1] == let && password[high-1] == let {
			continue
		}
		if password[low-1] == let || password[high-1] == let {
			part2++
		}
	}
	return part1, part2
}

func main() {
	input := getInput(false)

	t := time.Now()
	p1, p2 := check(input)
	fmt.Printf("%d, %d, %s\n", p1, p2, time.Since(t))
}
