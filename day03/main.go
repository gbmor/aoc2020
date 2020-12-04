package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func getMap(test bool) []string {
	if test {
		return []string{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		}
	}
	b, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(b), "\n")
}

func calc(hillMap []string, dx, dy int) int {
	var x, y, trees int
	width := len(hillMap[0])
	for {
		if hillMap[y][x] == '#' {
			trees++
		}
		x += dx
		x %= width
		y += dy
		if y >= len(hillMap) {
			break
		}
	}
	return trees
}

func main() {
	input := getMap(false)

	t := time.Now()
	out := calc(input, 3, 1)
	fmt.Printf("%d, %s\n", out, time.Since(t))

	pairs := [][]int{
		{1, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	t = time.Now()
	for _, p := range pairs {
		out *= calc(input, p[0], p[1])
	}
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
