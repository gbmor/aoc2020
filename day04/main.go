package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var hclRegex = regexp.MustCompile("[#][0-9a-f]+")
var pidRegex = regexp.MustCompile("[0-9]+")

func getInput(test bool) []string {
	var b []byte
	if test {
		b = []byte("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\niyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929\n\nhcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm\n\nhcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in")
	} else {
		b, _ = ioutil.ReadFile("input.txt")
	}
	return strings.Split(string(b), "\n")
}

func calc(s []string, partTwo bool) int {
	count := 0
	i := 0
	for i < len(s) {
		l := s[i]
		fields := make(map[string]int)
		for l != "" {
			current := strings.Split(l, " ")
			for _, c := range current {
				d := strings.Split(c, ":")
				if !partTwo {
					fields[d[0]]++
					continue
				}
				switch d[0] {
				case "byr":
					d1, _ := strconv.Atoi(d[1])
					if d1 > 1919 && d1 < 2003 {
						fields[d[0]]++
					}
				case "iyr":
					d1, _ := strconv.Atoi(d[1])
					if d1 > 2009 && d1 < 2021 {
						fields[d[0]]++
					}
				case "eyr":
					d1, _ := strconv.Atoi(d[1])
					if d1 > 2019 && d1 < 2031 {
						fields[d[0]]++
					}
				case "hgt":
					d1, _ := strconv.Atoi(d[1][:len(d[1])-2])
					if strings.HasSuffix(d[1], "cm") {
						if d1 > 149 && d1 < 194 {
							fields[d[0]]++
						}
					} else {
						if d1 > 58 && d1 < 77 {
							fields[d[0]]++
						}
					}
				case "hcl":
					m := hclRegex.FindString(d[1])
					if m != "" {
						fields[d[0]]++
					}
				case "ecl":
					if d[1] == "amb" || d[1] == "blu" || d[1] == "brn" || d[1] == "gry" || d[1] == "grn" || d[1] == "hzl" || d[1] == "oth" {
						fields[d[0]]++
					}
				case "pid":
					m := pidRegex.FindString(d[1])
					if m != "" && len(d[1]) == 9 {
						fields[d[0]]++
					}
				}
			}
			i++
			if i < len(s) {
				l = s[i]
			} else {
				break
			}
		}
		if fields["byr"] > 0 && fields["iyr"] > 0 && fields["eyr"] > 0 && fields["hgt"] > 0 && fields["hcl"] > 0 && fields["ecl"] > 0 && fields["pid"] > 0 {
			count++
		}
		i++
	}
	return count
}

func main() {
	input := getInput(false)

	t := time.Now()
	out := calc(input, false)
	fmt.Printf("%d, %s\n", out, time.Since(t))

	t = time.Now()
	out = calc(input, true)
	fmt.Printf("%d, %s\n", out, time.Since(t))
}
