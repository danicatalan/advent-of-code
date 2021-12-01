package passportprocessing

import (
	"log"
	"strconv"
	"strings"
)

// Part2 checks if all the fields in the provided passport are valid.
func Part2(pass string) bool {
	n := strings.Replace(pass, "\n", " ", -1)
	fields := strings.Split(n, " ")
	for _, f := range fields {
		k, v := f[0:3], f[4:]
		switch k {
		case "byr":
			y, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			if y < 1920 || y > 2002 {
				return false
			}
		case "iyr":
			y, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			if y < 2010 || y > 2020 {
				return false
			}
		case "eyr":
			y, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			if y < 2020 || y > 2030 {
				return false
			}
		case "hgt":
			u := v[len(v)-2:]
			if u != "cm" && u != "in" {
				return false
			}
			h, err := strconv.Atoi(v[:len(v)-2])
			if err != nil {
				log.Fatal(err)
			}
			if (u == "cm" && (h < 150 || h > 193)) || (u == "in" && (h < 59 || h > 76)) {
				return false
			}
		case "hcl":
			p1, p2 := v[0:1], v[1:]
			if p1 != "#" || len(p2) != 6 {
				return false
			}
		case "ecl":
			if !strings.Contains("amb blu brn gry grn hzl oth", v) {
				return false
			}
		case "pid":
			if len(v) != 9 {
				return false
			}
		}
	}
	return true
}
