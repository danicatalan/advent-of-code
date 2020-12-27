package tickettranslation

import (
	"log"
	"strconv"
	"strings"
)

//Part2 ...
func Part2(input string) int {
	notes := strings.Split(input, "\n\n")
	fields := strings.Split(notes[0], "\n")
	rules := make(map[string]func(int) bool)
	for _, f := range fields {
		name := strings.Split(f, ": ")[0]
		ranges := strings.Split(strings.Split(f, ": ")[1], " or ")
		range1, range2 := strings.Split(ranges[0], "-"), strings.Split(ranges[1], "-")

		var numbers1, numbers2 []int
		for _, n := range range1 {
			ni, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			numbers1 = append(numbers1, ni)
		}
		for _, n := range range2 {
			ni, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			numbers2 = append(numbers2, ni)
		}

		rules[name] = func(i int) bool {
			if numbers1[0] <= i && i <= numbers1[1] {
				return true
			}
			if numbers2[0] <= i && i <= numbers2[1] {
				return true
			}
			return false
		}
	}

	var validTickets []string
	tickets := strings.Split(notes[2], "\n")[1:]
	for _, t := range tickets {
		validTicket := true
		for _, v := range strings.Split(t, ",") {
			vi, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			var validValue bool
			for _, f := range rules {
				if f(vi) {
					validValue = true
				}
			}
			if !validValue {
				validTicket = false
			}
		}
		if validTicket {
			validTickets = append(validTickets, t)
		}
	}

	columns := make([][]string, len(fields))
	for _, t := range validTickets {
		for i, v := range strings.Split(t, ",") {
			columns[i] = append(columns[i], v)
		}
	}

	possible := make([][]string, len(fields))
	for i, c := range columns {
		for k, f := range rules {
			valid := true
			for _, v := range c {
				vi, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				if !f(vi) {
					valid = false
					break
				}
			}
			if valid {
				possible[i] = append(possible[i], k)
			}
		}
	}

	resolve, c := make([]string, len(fields)), 0
	for c < len(fields) {
		//time.Sleep(time.Second)
		//fmt.Println(possible, resolve, len(resolve), len(fields))
		for i, p := range possible {
			if len(p) == 1 {
				resolve[i] = p[0]
				possible = remove(possible, p[0])
				c++
				break
			}
		}
	}

	mult := 1
	myTicket := strings.Split(strings.Split(notes[1], "\n")[1], ",")
	for i, f := range resolve {
		if strings.HasPrefix(f, "departure") {
			n, err := strconv.Atoi(myTicket[i])
			if err != nil {
				log.Fatal(err)
			}
			mult *= n
		}
	}
	return mult
}

func remove(old [][]string, k string) [][]string {
	var new [][]string
	for _, s := range old {
		temp := []string{}
		for _, f := range s {
			if f != k {
				temp = append(temp, f)
			}
		}
		new = append(new, temp)
	}
	return new
}
