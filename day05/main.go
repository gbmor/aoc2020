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
	seatz := make([]int, 0)
	for _, v := range s {
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
		seatz = append(seatz, z)
	}
	sort.Ints(seatz)
	return seatz
}

func highestSeatID(s []int) int {
	id := 0
	for _, i := range s {
		if i > id {
			id = i
		}
	}
	return id
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
	out := highestSeatID(s)
	fmt.Printf("%d, %s\n", out, time.Since(t))
	out = missingSeat(s)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}