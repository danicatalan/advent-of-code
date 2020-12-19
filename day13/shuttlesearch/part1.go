package shuttlesearch

import (
	"log"
	"math"
	"strconv"
	"strings"
)

//Part1 ...
func Part1(input string) int {
	data := strings.Split(input, "\n")
	timestamp, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}
	buses := []int{}
	for _, b := range strings.Split(data[1], ",") {
		if b == "x" {
			continue
		}
		bi, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal(err)
		}
		buses = append(buses, bi)
	}
	ID, wait := 0, math.MaxInt64
	for _, b := range buses {
		if b-(timestamp%b) < wait {
			wait = b - (timestamp % b)
			ID = b
		}
	}
	return ID * wait
}
