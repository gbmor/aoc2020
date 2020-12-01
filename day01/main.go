package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type ExpenseReport []int

func newExpenseReport(test bool) ExpenseReport {
	if test {
		return ExpenseReport{
			1721,
			979,
			366,
			299,
			675,
			1456,
		}
	}

	b, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(b), "\n")
	out := make(ExpenseReport, len(lines))
	for i, l := range lines {
		n, _ := strconv.Atoi(l)
		out[i] = n
	}
	return out
}

func (input ExpenseReport) Part1() int {
	for _, a := range input {
		for _, b := range input {
			if a+b == 2020 {
				return a * b
			}
		}
	}
	return -1
}

func (input ExpenseReport) Part2() int {
	for _, a := range input {
		for _, b := range input {
			for _, c := range input {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	return -1
}

func main() {
	er := newExpenseReport(false)

	t := time.Now()
	p1 := er.Part1()
	fmt.Printf("Part 1: %d, %s\n", p1, time.Since(t))

	t = time.Now()
	p2 := er.Part2()
	fmt.Printf("Part 2: %d, %s\n", p2, time.Since(t))
}
