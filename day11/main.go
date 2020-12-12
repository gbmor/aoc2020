package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func getSeats(test bool) [][]rune {
	if test {
		out := make([][]rune, 10)
		for i := range out {
			out[i] = make([]rune, 10)
		}
		input := []string{
			"L.LL.LL.LL",
			"LLLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLLL",
			"L.LLLLLL.L",
			"L.LLLLL.LL",
		}
		for i, e := range input {
			for k, v := range e {
				out[i][k] = rune(v)
			}
		}
		return out
	}
	b, _ := ioutil.ReadFile("input.txt")
	bs := strings.Split(string(b), "\n")
	out := make([][]rune, len(bs))
	for i := range out {
		out[i] = make([]rune, len(bs[0]))
	}
	for i, e := range bs {
		for k, v := range e {
			out[i][k] = rune(v)
		}
	}
	return out
}

func clone(s [][]rune) [][]rune {
	n := make([][]rune, len(s))
	for i := range n {
		n[i] = make([]rune, len(s[0]))
	}
	for i := range s {
		copy(n[i], s[i])
	}
	return n
}

func exec(s [][]rune, execFunc func([][]rune) ([][]rune, int)) (int, int) {
	var old [][]rune
	newx, diffs := execFunc(s)
	generations := 1
	for diffs > 0 {
		old = clone(newx)
		newx, diffs = execFunc(old)
		generations++
	}
	occupied := 0
	for _, e := range newx {
		for _, v := range e {
			if v == '#' {
				occupied++
			}
		}
	}
	return generations, occupied
}

func exec1(s [][]rune) ([][]rune, int) {
	current := clone(s)
	diffs := 0
	for i, e := range s {
		for j, f := range e {
			if f == '.' {
				continue
			}
			nearby := 0
			if j > 0 && e[j-1] == '#' {
				nearby++
			}
			if j < len(e)-1 && e[j+1] == '#' {
				nearby++
			}
			if i > 0 && s[i-1][j] == '#' {
				nearby++
			}
			if i < len(s)-1 && s[i+1][j] == '#' {
				nearby++
			}
			if i < len(s)-1 && j < len(e)-1 && s[i+1][j+1] == '#' {
				nearby++
			}
			if i < len(s)-1 && j > 0 && s[i+1][j-1] == '#' {
				nearby++
			}
			if i > 0 && j > 0 && s[i-1][j-1] == '#' {
				nearby++
			}
			if i > 0 && j < len(e)-1 && s[i-1][j+1] == '#' {
				nearby++
			}
			if f == 'L' && nearby == 0 {
				current[i][j] = '#'
				diffs++
			}
			if f == '#' && nearby > 3 {
				current[i][j] = 'L'
				diffs++
			}
		}
	}
	return current, diffs
}

func exec2(s [][]rune) ([][]rune, int) {
	current := clone(s)
	diffs := 0
	for i, e := range s {
		for j, f := range e {
			if f == '.' {
				continue
			}
			nearby := lineOfSight(s, j, 1, i, 0)
			nearby += lineOfSight(s, j, -1, i, 0)
			nearby += lineOfSight(s, j, 0, i, 1)
			nearby += lineOfSight(s, j, 0, i, -1)
			nearby += lineOfSight(s, j, 1, i, 1)
			nearby += lineOfSight(s, j, 1, i, -1)
			nearby += lineOfSight(s, j, -1, i, -1)
			nearby += lineOfSight(s, j, -1, i, 1)
			if f == 'L' && nearby == 0 {
				current[i][j] = '#'
				diffs++
			}
			if f == '#' && nearby > 4 {
				current[i][j] = 'L'
				diffs++
			}
		}
	}
	return current, diffs
}

func lineOfSight(s [][]rune, x, dx, y, dy int) int {
	x += dx
	y += dy
	for x < len(s[0]) && x > -1 && y < len(s) && y > -1 {
		if s[y][x] == '#' {
			return 1
		}
		if s[y][x] == 'L' {
			return 0
		}
		x += dx
		y += dy
	}
	return 0
}

func main() {
	input := getSeats(false)
	t := time.Now()
	gens, occ := exec(input, exec1)
	fmt.Printf("Generations: %d, Seats Occupied: %d :: %s\n", gens, occ, time.Since(t))

	input = getSeats(false)
	t = time.Now()
	gens, occ = exec(input, exec2)
	fmt.Printf("Generations: %d, Seats Occupied: %d :: %s\n", gens, occ, time.Since(t))
}
