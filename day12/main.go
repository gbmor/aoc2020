package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

type Movement struct {
	direction rune
	distance  int
}

type Plane struct {
	x, y, wx, wy, facing int
}

func getInput(test bool) []Movement {
	if test {
		in := []string{
			"F10",
			"N3",
			"F7",
			"R90",
			"F11",
		}
		out := make([]Movement, 5)
		for i := range out {
			out[i].direction = rune(in[i][0])
			out[i].distance, _ = strconv.Atoi(in[i][1:])
		}
		return out
	}
	b, _ := ioutil.ReadFile("input.txt")
	bs := strings.Split(string(b), "\n")
	out := make([]Movement, len(bs))
	for i := range out {
		out[i].direction = rune(bs[i][0])
		out[i].distance, _ = strconv.Atoi(bs[i][1:])
	}
	return out
}

func moveShip(p *Plane, m Movement) {
	switch m.direction {
	case 'F':
		switch p.facing {
		case 0:
			p.y += m.distance
		case 1:
			p.x += m.distance
		case 2:
			p.y -= m.distance
		case 3:
			p.x -= m.distance
		}
	case 'N':
		p.y += m.distance
	case 'E':
		p.x += m.distance
	case 'S':
		p.y -= m.distance
	case 'W':
		p.x -= m.distance
	case 'R':
		turns := m.distance / 90
		p.facing += turns
	case 'L':
		turns := m.distance / 90
		p.facing -= turns
	}
	p.facing = ((p.facing % 4) + 4) % 4
}

func moveWithWaypoint(p *Plane, m Movement) {
	switch m.direction {
	case 'F':
		p.x += m.distance * p.wx
		p.y += m.distance * p.wy
	case 'N':
		p.wy += m.distance
	case 'E':
		p.wx += m.distance
	case 'S':
		p.wy -= m.distance
	case 'W':
		p.wx -= m.distance
	default:
		turns := m.distance / 90
		for i := 0; i < turns; i++ {
			x := p.wx
			p.wx = p.wy
			p.wy = x
			if m.direction == 'R' {
				p.wy *= -1
			} else {
				p.wx *= -1
			}
		}
	}
}

func exec(input []Movement, moveFunc func(p *Plane, m Movement)) int {
	plane := new(Plane)
	plane.facing = 1
	plane.wx = 10
	plane.wy = 1
	for i := range input {
		moveFunc(plane, input[i])
	}
	return int(math.Abs(float64(plane.x)) + math.Abs(float64(plane.y)))
}

func main() {
	input := getInput(false)

	t := time.Now()
	out := exec(input, moveShip)
	fmt.Printf("%d, %s\n", out, time.Since(t))

	t = time.Now()
	out = exec(input, moveWithWaypoint)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
