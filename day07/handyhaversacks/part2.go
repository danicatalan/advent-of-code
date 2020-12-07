package handyhaversacks

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Part2 accepts a set of color-coded bag rules and returns the amount of bags inside a "shiny gold" bag.
func Part2(input string) int {
	rules := normalize2(input)
	c := recursion2(rules, "shiny gold")
	return c
}

func recursion2(rules map[string][]string, input string) int {
	bags := rules[input]
	if len(bags) == 0 {
		return 0
	}
	total := 0
	for _, b := range bags {
		name := b[2:]
		number, err := strconv.Atoi(b[:1])
		if err != nil {
			log.Fatal(err)
		}
		total += number + number*recursion2(rules, name)
	}
	return total
}

func normalize2(input string) map[string][]string {
	input = strings.ReplaceAll(input, ".", "")
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, "contain ", "")
	input = strings.ReplaceAll(input, "bags", "")
	input = strings.ReplaceAll(input, "bag", "")
	re := regexp.MustCompile("(  )|(no other)")
	rules := make(map[string][]string)
	for _, r := range strings.Split(input, "\n") {
		t := re.Split(strings.TrimSpace(r), -1)
		if t[1] != "" {
			rules[t[0]] = t[1:]
		} else {
			rules[t[0]] = []string{}
		}
	}
	return rules
}
