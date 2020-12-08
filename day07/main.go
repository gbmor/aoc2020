package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type bagList struct {
	m map[string][]Bag
}

type Bag struct {
	q int
	s string
}

func newBagList() *bagList {
	b := new(bagList)
	b.m = make(map[string][]Bag)
	return b
}

func getInput(test, p2test bool) []string {
	if test {
		return []string{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"bright white bags contain 1 shiny gold bag.",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		}
	}
	if p2test {
		return []string{
			"shiny gold bags contain 2 dark red bags.",
			"dark red bags contain 2 dark orange bags.",
			"dark orange bags contain 2 dark yellow bags.",
			"dark yellow bags contain 2 dark green bags.",
			"dark green bags contain 2 dark blue bags.",
			"dark blue bags contain 2 dark violet bags.",
			"dark violet bags contain no other bags.",
		}
	}
	b, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(b), "\n")
}

func (b *bagList) load(s []string) {
	for _, e := range s {
		split := strings.Split(e, " bags contain ")
		src := split[0]
		dst := split[1]
		if dst == "no other bags" {
			b.m[src] = nil
			continue
		}
		bags := make([]Bag, 0)
		dstS := strings.Split(dst, ", ")
		for _, v := range dstS {
			vs := strings.Fields(v)
			i, _ := strconv.Atoi(vs[0])
			c := fmt.Sprintf("%s %s", vs[1], vs[2])
			bags = append(bags, Bag{q: i, s: c})
		}
		b.m[src] = bags
	}
}

func (b *bagList) hasGoldBagInside() int {
	count := 0
	for _, v := range b.m {
		count += holdsGold(b, v)

	}
	return count
}

func holdsGold(b *bagList, bags []Bag) int {
	for _, bag := range bags {
		if bag.s == "shiny gold" {
			return 1
		}
		if bag.q != 0 {
			out := holdsGold(b, b.m[bag.s])
			if out == 1 {
				return 1
			}
		}
	}
	return 0
}

func (b *bagList) bagsInsideGold(bags []Bag, count int) int {
	if bags == nil {
		return count
	}

	for _, k := range bags {
		if k.q == 0 {
			continue
		}
		count += k.q
		count += k.q * b.bagsInsideGold(b.m[k.s], 0)
	}

	return count
}

func main() {
	input := getInput(false, false)
	bl := newBagList()
	bl.load(input)

	t := time.Now()
	out := bl.hasGoldBagInside()
	fmt.Printf("%d, %s\n", out, time.Since(t))

	t = time.Now()
	out = bl.bagsInsideGold(bl.m["shiny gold"], 0)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
