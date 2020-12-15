package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func getInput(test int) []string {
	if test == 1 {
		return []string{
			"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			"mem[8] = 11",
			"mem[7] = 101",
			"mem[8] = 0",
		}
	}
	if test == 2 {
		return []string{
			"mask = 000000000000000000000000000000X1001X",
			"mem[42] = 100",
			"mask = 00000000000000000000000000000000X0XX",
			"mem[26] = 1",
		}
	}
	b, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(b), "\n")
}

func floaters(s string, d int64, cache map[string]int64) {
	if strings.Contains(s, "X") {
		n := strings.Replace(s, "X", "0", 1)
		z := strings.Replace(s, "X", "1", 1)
		floaters(n, d, cache)
		floaters(z, d, cache)
	} else {
		cache[s] = d
	}
}

func part1(s []string) int64 {
	memmap := make(map[int]int64)
	mask := ""
	for _, e := range s {
		if strings.HasPrefix(e, "mask") {
			mask = strings.TrimPrefix(e, "mask = ")
			continue
		}
		split := strings.Split(e, " = ")
		addr, _ := strconv.Atoi(split[0][4 : len(split[0])-1])
		val, _ := strconv.Atoi(split[1])
		binary := strconv.FormatInt(int64(val), 2)
		gap := 36 - len(binary)
		bldr := strings.Builder{}
		for j := 0; j < gap; j++ {
			bldr.WriteRune('0')
		}
		bldr.WriteString(binary)
		bin := bldr.String()
		newVal := strings.Builder{}
		for j := 0; j < len(mask); j++ {
			if mask[j] == 'X' {
				newVal.WriteRune(rune(bin[j]))
				continue
			}
			newVal.WriteRune(rune(mask[j]))
		}
		memmap[addr], _ = strconv.ParseInt(newVal.String(), 2, 64)
	}
	val := int64(0)
	for _, v := range memmap {
		val += v
	}
	return val
}

func part2(s []string) int64 {
	mask := ""
	memmap := make(map[string]int64)
	for _, e := range s {
		if strings.HasPrefix(e, "mask") {
			mask = strings.TrimPrefix(e, "mask = ")
			continue
		}
		split := strings.Split(e, " = ")
		memVal, _ := strconv.ParseInt(split[1], 10, 64)
		addr, _ := strconv.Atoi(split[0][4 : len(split[0])-1])
		binary := strconv.FormatInt(int64(addr), 2)
		gap := 36 - len(binary)
		bldr := strings.Builder{}
		for j := 0; j < gap; j++ {
			bldr.WriteRune('0')
		}
		bldr.WriteString(binary)
		bin := bldr.String()
		newVal := strings.Builder{}
		for j := 0; j < len(mask); j++ {
			switch mask[j] {
			case '1':
				newVal.WriteRune('1')
			case 'X':
				newVal.WriteRune('X')
			case '0':
				newVal.WriteRune(rune(bin[j]))
			}
		}
		addrVal := newVal.String()
		floaters(addrVal, memVal, memmap)
	}
	sum := int64(0)
	for _, v := range memmap {
		sum += v
	}
	return sum
}

func main() {
	i := getInput(0)

	t := time.Now()
	out := part1(i)
	fmt.Printf("%d, %s\n", out, time.Since(t))

	t = time.Now()
	out = part2(i)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
