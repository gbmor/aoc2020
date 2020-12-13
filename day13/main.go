package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func getSchedule(test bool) (int, []int) {
	if test {
		sample := []string{
			"939",
			"7,13,x,x,59,x,31,19",
		}
		depart, _ := strconv.Atoi(sample[0])
		bs := make([]int, 0)
		ss := strings.Split(sample[1], ",")
		for _, e := range ss {
			if e == "x" {
				bs = append(bs, 0)
			} else {
				bd, _ := strconv.Atoi(e)
				bs = append(bs, bd)
			}
		}
		return depart, bs
	}
	b, _ := ioutil.ReadFile("input.txt")
	bs := strings.Split(string(b), "\n")
	depart, _ := strconv.Atoi(bs[0])
	bsd := make([]int, 0)
	ss := strings.Split(bs[1], ",")
	for _, e := range ss {
		if e == "x" {
			bsd = append(bsd, 0)
		} else {
			bd, _ := strconv.Atoi(e)
			bsd = append(bsd, bd)
		}
	}
	return depart, bsd
}

func part1(d int, b []int) int {
	for i := d; i > 0; i++ {
		for _, c := range b {
			if c != 0 && i%c == 0 {
				return (i - d) * c
			}
		}
	}
	return -1
}

func crt(a, n []*big.Int) *big.Int {
	one := big.NewInt(1)
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p)
}

func part2(b []int) *big.Int {
	a := make([]*big.Int, 0)
	bb := make([]*big.Int, 0)
	for i, e := range b {
		if e == 0 {
			continue
		}
		bb = append(bb, big.NewInt(int64(e)))
		a = append(a, big.NewInt(-1*int64(i)))
	}
	return crt(a, bb)
}

func main() {
	d, b := getSchedule(false)

	t := time.Now()
	out := part1(d, b)
	fmt.Printf("%d, %s\n", out, time.Since(t))

	t = time.Now()
	out2 := part2(b)
	fmt.Printf("%d, %s\n", out2, time.Since(t))
}
