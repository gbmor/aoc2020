package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

func getInput(test bool) []string {
	if test {
		return []string{
			"BFFFBBFRRR",
			"FFFBBBFRRR",
			"BBFFBBFRLL",
		}
	}
	b, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(b), "\n")
}

func makeRows(s int) []int {
	rows := make([]int, s)
	for i := range rows {
		rows[i] = i
	}
	return rows
}

func getSeats(s []string) []int {
	seatz := make([]int, len(s))
	for i, v := range s {
		rows := makeRows(128)
		seats := makeRows(8)
		for _, e := range v {
			switch e {
			case 'F':
				l := len(rows) / 2
				rows = rows[:l]
			case 'B':
				l := len(rows) / 2
				rows = rows[l:]
			case 'R':
				l := len(seats) / 2
				seats = seats[l:]
			case 'L':
				l := len(seats) / 2
				seats = seats[:l]
			}
		}
		z := rows[0]*8 + seats[0]
		seatz[i] = z
	}
	sort.Ints(seatz)
	return seatz
}

func missingSeat(s []int) int {
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i]-1 {
			return s[i] - 1
		}
	}
	return -1
}

func main() {
	input := getInput(false)

	t := time.Now()
	s := getSeats(input)
	fmt.Printf("%d, %s\n", s[len(s)-1], time.Since(t))
	out := missingSeat(s)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
