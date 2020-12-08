package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func getInput(test bool) []string {
	if test {
		return []string{
			"nop +0",
			"acc +1",
			"jmp +4",
			"acc +3",
			"jmp -3",
			"acc -99",
			"acc +1",
			"jmp -4",
			"acc +6",
		}
	}
	b, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(b), "\n")
}

func ACCBeforeFirstRepeat(s []string) (int, int) {
	acc := 0
	hist := make(map[int]int)
	i := 0
	for i < len(s) {
		if hist[i] > 0 {
			return acc, -1
		}
		hist[i]++
		ss := strings.Fields(s[i])
		d, _ := strconv.Atoi(ss[1])
		switch ss[0] {
		case "acc":
			acc += d
			i++
		case "jmp":
			i += d
		case "nop":
			i++
		}
	}
	return acc, i
}

func ACCAfterEnd(s []string) int {
	for i, e := range s {
		if strings.HasPrefix(e, "acc") {
			continue
		}
		f := make([]string, len(s))
		copy(f, s)
		ss := strings.Fields(e)
		if strings.HasPrefix(e, "jmp") {
			f[i] = fmt.Sprintf("nop %s", ss[1])
		}
		if strings.HasPrefix(e, "nop") {
			f[i] = fmt.Sprintf("jmp %s", ss[1])
		}
		out, j := ACCBeforeFirstRepeat(f)
		if j == len(f) {
			return out
		}
	}
	return -1
}

func main() {
	input := getInput(false)

	t := time.Now()
	out, _ := ACCBeforeFirstRepeat(input)
	fmt.Printf("%d, %s\n", out, time.Since(t))

	t = time.Now()
	out = ACCAfterEnd(input)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
