package handyhaversacks

import (
	"regexp"
	"strings"
)

// Part1 accepts a set of color-coded bag rules and returns the amount of rules that could contain a "shiny gold" bag.
func Part1(input string) int {
	rules := normalize(input)
	c := 0
	for k, v := range rules {
		if k != "shiny gold" {
			if recursion(rules, v) {
				c++
			}
		}
	}
	return c
}

func recursion(rules map[string][]string, input []string) bool {
	if len(input) == 0 {
		return false
	}
	new := []string{}
	for _, v := range input {
		if v == "shiny gold" {
			return true
		}
		for _, w := range rules[v] {
			new = append(new, w)
		}
	}
	return recursion(rules, new)
}

func normalize(input string) map[string][]string {
	input = strings.ReplaceAll(input, ".", "")
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, "contain ", "")
	input = strings.ReplaceAll(input, "bags", "")
	input = strings.ReplaceAll(input, "bag", "")
	input = strings.ReplaceAll(input, "  ", " ")
	re := regexp.MustCompile("( [0-9] )|( no other)")
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
